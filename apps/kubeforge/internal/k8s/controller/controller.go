// ############################################################
// Copyright (c) 2024 wsadza
// Released under the MIT license
// ------------------------------------------------------------
//
// This Kubernetes controller is designed to manage and synchronize
// custom resources (CRDs) with dynamic Kubernetes resources. It
// listens to CRD changes, processes the associated YAML configurations,
// and ensures that dynamic resources are created or updated based on
// the desired state defined in the YAML files. The controller handles
// resource creation, update, and deletion efficiently by leveraging
// Kubernetes informers, dynamic clients, and workqueues.
//
// Key Features:
// - Synchronization of CRDs with dynamic Kubernetes resources
// - YAML-based configuration handling and merging
// - Automatic resource creation and updates based on CRD state
// - Error handling with retries and backoffs
//
// The controller manages multiple worker goroutines to process
// resource synchronization tasks and ensures smooth operation by
// handling resource state transitions appropriately.
//
// ############################################################

package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"time"

	"dario.cat/mergo"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog/v2"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	pkgRuntime "k8s.io/apimachinery/pkg/runtime"

	crdv1 "kubeforge/internal/k8s/api/v1"
	yaml "kubeforge/internal/ops/yaml"
	yamlMisc "kubeforge/internal/ops/yaml/misc"

	crdClientSet "kubeforge/pkg/generated/clientset/versioned"
	crdListers "kubeforge/pkg/generated/listers/api/v1"

	controllerMisc "kubeforge/internal/k8s/controller/misc"
)

type controller struct {
  workingContext            context.Context
	workingWorkers	          int
	workqueue                 workqueue.TypedRateLimitingInterface[cache.ObjectName]  
  recorder                  record.EventRecorder
  controllerName            string
  dynInformers              []cache.SharedInformer
  k8sClient                 kubernetes.Interface
  crdClient                 crdClientSet.Interface
  dynClient                 dynamic.Interface 
  crdLister                 crdListers.OverlayLister
  crdsSynced                cache.InformerSynced
  sourceConfiguration       string
	updateReadyz              func(bool)
	updateHealthz             func(bool)
}

// Run will set up the event handlers for types we are interested in, as well
// as syncing informer caches and starting workers. It will block until stopCh
// is closed, at which point it will shutdown the workqueue and wait for
// workers to finish processing their current work items.
func (controller *controller) Run() error {

  defer controller.updateReadyz(false)
  defer controller.updateHealthz(false)

	defer runtime.HandleCrash()
	defer controller.workqueue.ShutDown()
	logger := klog.FromContext(controller.workingContext)

	// Start the informer factories to begin populating the informer caches
	logger.Info("Starting Controller")

	// Wait for the caches to be synced before starting workers
	logger.Info("Waiting for informer caches to sync")

	// Create a slice to hold HasSynced functions for dynamic resources
	var dynamicHasSynced []cache.InformerSynced

	// Assuming you have dynamic informers set up for Pods, PVCs, ConfigMaps
	for _, informer := range controller.dynInformers {
		dynamicHasSynced = append(dynamicHasSynced, informer.HasSynced)
	}
  dynamicHasSynced = append(dynamicHasSynced, controller.crdsSynced)

  // Wait for syncs
  ok := cache.WaitForCacheSync(
    controller.workingContext.Done(), 
    dynamicHasSynced...,
  );
	if !ok {
		return fmt.Errorf("failed to wait for caches to sync")
	}

	logger.Info("Starting workers", "count", controller.workingWorkers)

  // Set ready just before launching 
  controller.updateReadyz(true)

	// Launch two workers to process resources
	for i := 0; i < controller.workingWorkers; i++ {
		go wait.UntilWithContext(
      controller.workingContext, 
      controller.runWorker, 
      time.Second,
    )
	}

	logger.Info("Started workers")
	<-controller.workingContext.Done()
	logger.Info("Shutting down workers")

	return nil
}

// runWorker is a long-running function that will continually call the
// processNextWorkItem function in order to read and process a message on the
// workqueue.
func (controller *controller) runWorker(ctx context.Context) {
	for controller.processNextWorkItem(ctx) {}
}

// processNextWorkItem will read a single work item off the workqueue and
// attempt to process it, by calling the syncHandler.
func (controller *controller) processNextWorkItem(ctx context.Context) bool {
	objRef, shutdown := controller.workqueue.Get()
	logger := klog.FromContext(ctx)

	if shutdown {
		return false
	}

	defer controller.workqueue.Done(objRef)

	err := controller.syncHandler(ctx, objRef)
	if err == nil {
    // Procesing correctly
    controller.updateHealthz(true)
		controller.workqueue.Forget(objRef)
		logger.Info("Successfully synced", "objectName", objRef)
		return true
	} 

  // Procesing uncorrectly
  controller.updateHealthz(false)

	runtime.HandleErrorWithContext(
    ctx, 
    err, 
    "Error syncing; requeuing for later retry", 
    "objectReference", 
    objRef,
  );
	controller.workqueue.AddRateLimited(objRef)
	return true
}

// syncHandler compares the actual state with the desired, and attempts to converge the two. 
func (controller *controller) syncHandler (ctx context.Context, obj cache.ObjectName) error {
	  logger := klog.FromContext(ctx)

    // Get the overllay ~ CRD
    crdOverlay, err := controller.getCRDOverlay(obj, logger)
    if err != nil {
        return err
    }    

    // Unmarshal custom YAML data
    customData, err := controller.unmarshalCustomYAML(crdOverlay, logger)
    if err != nil {
        return err
    }

    // Unmarshal default YAML configuration
    sourceData, err := controller.unmarshalSourceYAML(logger)
    if err != nil {
        return err
    }        

    // Merge YAML data
    dataMergedMap, err := controller.mergeYAML(sourceData, customData)
    if err != nil {
        return err
    }

    // Set up default metadata
    objectMetadata, err := controller.getMetadata(crdOverlay, logger)
    if err != nil {
        return err
    }       
    
    // Iterate over resource types
    discoveryClient := controller.k8sClient.Discovery()
    for resourceType, resourceList := range dataMergedMap {

      schema, err := controller.getResourceSchema(resourceType, discoveryClient, logger)
      if err != nil {
          return err
      }

      for _, resourceDefinition := range resourceList.([]interface{}) {
        if err := controller.processResource(resourceDefinition, schema, objectMetadata, logger); err != nil {
            continue
        }        
      }
    }

  controller.recorder.Event(crdOverlay, corev1.EventTypeNormal, "Success", "Success")
  return nil
}

// ------------------------------------------------------------

// getCRDOverlay retrieves the CRD overlay.
func (controller *controller) getCRDOverlay(obj cache.ObjectName, logger klog.Logger) (*crdv1.Overlay, error) {
    crdOverlay, err := controller.crdLister.Overlays(obj.Namespace).Get(obj.Name)
    if err != nil {
        logger.Error(err, "Failed to get CRD overlay")
        return nil, err
    }
    return crdOverlay, nil
}

// unmarshalCustomYAML unmarshals the custom YAML data from the CRD overlay.
func (controller *controller) unmarshalCustomYAML(crdOverlay *crdv1.Overlay, logger klog.Logger) (map[string]interface{}, error) {
    var dataCustom map[string]interface{}
    err := yaml.Unmarshal(string(crdOverlay.Spec.Data.Raw), &dataCustom, true)
    if err != nil {
        logger.Error(err, "Failed to unmarshal custom YAML")
        return nil, err
    }
    return dataCustom, nil
}

// unmarshalDefaultYAML unmarshals the default YAML configuration.
func (controller *controller) unmarshalSourceYAML(logger klog.Logger) (map[string]interface{}, error) {
    var defaultRaw map[string]interface{}
    err := yaml.Unmarshal(controller.sourceConfiguration, &defaultRaw, true)
    if err != nil {
        logger.Error(err, "Error unmarshaling default YAML")
        return nil, err
    }
    return defaultRaw, nil
}

// mergeYAML merges the custom YAML with the default YAML configuration.
func (controller *controller) mergeYAML(defaultRaw, dataCustom map[string]interface{}) (map[string]interface{}, error) {
    dataMerged := yamlMisc.StructuresMergeByName(defaultRaw, dataCustom)
    dataMergedMap := dataMerged.(map[string]interface{})
    
    err := mergo.Merge(&dataCustom, dataMergedMap, mergo.WithOverride)
    if err != nil {
        return nil, fmt.Errorf("error merging YAML: %v", err)
    }

    return dataMergedMap, nil
}

// getMetadata sets up the default Kubernetes object metadata.
func (controller *controller) getMetadata(
  crdOverlay *crdv1.Overlay, 
  logger klog.Logger,
) (
  metav1.ObjectMeta, 
  error,
) {

    crdOverlayKind := crdOverlay.Kind
    if crdOverlayKind == "" {
	    re := regexp.MustCompile(`"kind":"([^"]+)"`)
	    match := re.FindStringSubmatch(fmt.Sprint(crdOverlay))

	    // Check if a match was found and print the result
	    if len(match) > 1 {
        crdOverlayKind = match[1]
      } else {
        logger.Error(nil, "Failed to get CRD overlay Kind")
        return metav1.ObjectMeta{}, fmt.Errorf("Failed to get CRD overlay Kind")
      }
    }

    objectMetadataReferences := []metav1.OwnerReference{
        *metav1.NewControllerRef(crdOverlay, crdv1.SchemeGroupVersion.WithKind(crdOverlayKind)),
    }
    objectMetadata := metav1.ObjectMeta{
        Namespace:       crdOverlay.Namespace,
        OwnerReferences: objectMetadataReferences,
    }

    return objectMetadata, nil
}

// getResourceSchema retrieves the schema of a given resource type.
func (controller *controller) getResourceSchema(
  resourceType    string, 
  discoveryClient discovery.DiscoveryInterface, 
  logger          klog.Logger,
) (
  *schema.GroupVersionResource,
  error,
) {
    schema, err := controllerMisc.GetGroupVersionResource(resourceType, discoveryClient)
    if err != nil {
        logger.Error(err, fmt.Sprintf("Error retrieving schema for resource type '%v'", resourceType))
        return nil, err
    }
    return schema, nil
}

// processResource processes each resource (converts, compares, and applies changes).
func (controller *controller) processResource(
  resourceDefinition interface{}, 
  schema         *schema.GroupVersionResource, 
  objectMetadata metav1.ObjectMeta, 
  logger         klog.Logger,
) error {

    objMeta, err := pkgRuntime.DefaultUnstructuredConverter.ToUnstructured(&resourceDefinition)
    if err != nil {
        return fmt.Errorf("failed to convert resource to unstructured format: %v", err)
    }

    marshaledData, _ := json.Marshal(objMeta)
    appliedConfiguration := strings.ReplaceAll(string(marshaledData), "\n", " ")
    metadataAnnotations := map[string]string{
        "kubeforge.sh/last-applied-configuration": appliedConfiguration,
    }

    createdResource := &unstructured.Unstructured{Object: objMeta}
    createdResource.SetNamespace(objectMetadata.Namespace)

    createdAnnotations := createdResource.GetAnnotations()
    if createdAnnotations != nil {
      for key, value := range metadataAnnotations {
        createdAnnotations[key] = value
      }
      createdResource.SetAnnotations(createdAnnotations)
    } else {
      createdResource.SetAnnotations(metadataAnnotations)
    }

    createdResource.SetOwnerReferences(objectMetadata.OwnerReferences)

    overrideNameAnnotation := createdResource.GetAnnotations()["kubeforge.sh/override-name"]
    if overrideNameAnnotation != "" {
        createdResource.SetName(overrideNameAnnotation)
    }

    resourceClient := controller.dynClient.Resource(*schema).Namespace(createdResource.GetNamespace())
    resourceName := createdResource.GetName()

    return controller.createOrUpdateResource(resourceClient, createdResource, resourceName, logger)
}

// createOrUpdateResource checks if the resource exists and either updates or creates it.
func (controller *controller) createOrUpdateResource(
  resourceClient  dynamic.ResourceInterface, 
  createdResource *unstructured.Unstructured, 
  resourceName    string, 
  logger          klog.Logger,
) error {

    existingResource, err := resourceClient.Get(context.Background(), resourceName, metav1.GetOptions{})
    if err != nil && !errors.IsNotFound(err) {
        return err
    }

    if existingResource != nil {
        createdAnnotation := createdResource.GetAnnotations()["kubeforge.sh/last-applied-configuration"]
        existingAnnotation := existingResource.GetAnnotations()["kubeforge.sh/last-applied-configuration"]

        if createdAnnotation == existingAnnotation {
            logger.Info("Resource already exists and is up-to-date")
            return nil
        }

        _, err = resourceClient.Create(
          context.Background(), 
          createdResource, 
          metav1.CreateOptions{FieldManager: controller.controllerName, DryRun: []string{"All"}},
        )
        if err != nil {
          logger.Error(err, "Validation failed")
          return err
        }

        err := resourceClient.Delete(context.Background(), resourceName, metav1.DeleteOptions{})
        if err != nil {
            logger.Error(err, "Failed to delete existing resource")
            return err
        }
    } else {
        _, err := resourceClient.Create(
          context.Background(), 
          createdResource, 
          metav1.CreateOptions{FieldManager: "controller"},
        )
        if err != nil {
            logger.Error(err, "Failed to create resource")
            return err
        }
    }

    return nil
}

// logSuccessEvent logs a success event after completing the operation.
func (controller *controller) logSuccessEvent(crdOverlay *crdv1.Overlay) {
    controller.recorder.Event(crdOverlay, corev1.EventTypeNormal, "Success", "Success")
}
