// ############################################################
// Copyright (c) 2024 wsadza
// Released under the MIT license
// ------------------------------------------------------------
//
// Purpose:
// This file defines the core logic for building and configuring a custom Kubernetes controller.
// It initializes necessary Kubernetes clients, informers, work queues, and event recorders, 
// setting up the environment to manage resources effectively within the cluster.
//
// Key Components:
// 1. **controllerDirector**: This is the central struct that drives the setup of the controller. 
//    It takes a `controllerBuilder` as input, validates it, and uses it to configure various components of the controller.
//
// 2. **Kubernetes Clients**: Clients are initialized for interacting with Kubernetes API resources (e.g., K8S API server, 
//    dynamic resources, and custom resources) through the `setupKubernetesClients` method.
//
// 3. **Dynamic Informers**: This component is responsible for setting up informers to watch standard Kubernetes resources 
//    (e.g., pods, PVCs, config maps). The `setupDynamicInformer` method configures this part of the controller.
//
// 4. **Custom Resource Informers**: A custom informer for managing the custom resources, specifically for instances of 
//    a custom CRD, is set up by `setupCustomResourceInformer`.
//
// 5. **Work Queue**: The work queue is initialized and managed with a rate-limited queue to handle work items efficiently. 
//    This setup is done via the `setupWorkQueue` method.
//
// 6. **Event Recorder**: This part handles event recording within Kubernetes, ensuring events related to the controller's 
//    actions are logged properly. The `setupEventRecorder` method sets up the event recorder and links it with Kubernetes' event system.
//
// ############################################################

package controller

import (
	"fmt"
	"reflect"
	"time"

	"golang.org/x/time/rate"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/dynamic/dynamicinformer"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
  "k8s.io/klog/v2"

	crdClientSet "kubeforge/pkg/generated/clientset/versioned"
	crdScheme "kubeforge/pkg/generated/clientset/versioned/scheme"
	crdInformeres "kubeforge/pkg/generated/informers/externalversions"

	corev1 "k8s.io/api/core/v1"
)

type controllerDirector struct {
  builder controllerBuilder
}

func NewControllerDirector(builder *controllerBuilder) *controllerDirector{
  return &controllerDirector{builder: *builder}
}

func (director *controllerDirector) ConstructController() (*controller, error) {

	logger := klog.FromContext(director.builder.workingContext)
  logger.Info("Starting building controller")

	// Validate if all `mandatory` controllerBuilder fields was assigned
  logger.Info("Validate controller struct")
  if err := director.validateBuilder(); err != nil {
    return nil, err 
	}

	// Create initial controller overlay 
	controller := &controller{
		controllerName:      director.builder.controllerName,
		workingContext:      director.builder.workingContext,
		workingWorkers:      director.builder.workingWorkers,
    sourceConfiguration: director.builder.sourceConfiguration,
    updateHealthz:       director.builder.updateHealthz,
    updateReadyz:        director.builder.updateReadyz,   
	}

  // Create Connection Configuration
  logger.Info("Create kubernetes connections")
  if err := director.setupKubernetesClients(controller); err != nil {
    return nil, err
  }

  // Sets up dynamic informers for the specified resources
  logger.Info("Create kubernetes dynamic informer factory")
  if err := director.setupDynamicInformer(controller); err != nil {
    return nil, err
  }

  // Sets up custom (CRD) informers for the specified resources
  logger.Info("Create kubernetes custom informer factory")
  if err := director.setupCustomResourceInformer(controller); err != nil {
    return nil, err
  }

  // Sets up the work queue
  logger.Info("Create workqueue")
  if err := director.setupWorkQueue(controller); err != nil {
    return nil, err
  }

  // Sets up the event recorder
  logger.Info("Create event recorder")
  if err := director.setupEventRecorder(controller); err != nil {
    return nil, err
  }

  logger.Info("Controller initialized sucesffully")
  return controller, nil
}

// ############################################################

// validateBuilder checks if all mandatory fields in the builder struct are set.
func (director *controllerDirector) validateBuilder() error {
	var missingFields []string

  val := reflect.ValueOf(director.builder)
	if val.Kind() == reflect.Ptr {
		val = val.Elem() // dereference pointer if needed
	}

	// Iterate over all struct fields
	for i := 0; i < val.NumField(); i++ {

    field      := val.Type().Field(i)
		fieldValue := val.Field(i)
    fieldName  := field.Name
    fieldTag   := field.Tag

    mandatoryTag := fieldTag.Get("mandatory");
    if mandatoryTag == "true" {
      if fieldValue.IsZero() {
        missingFields = append(missingFields, fieldName)
      }
    }
	}

	// Return whether the struct is fully filled and the list of missing fields
	if len(missingFields) > 0 {
    return fmt.Errorf(
      "controllerBuilder struct is missing values in the following fields: %v", 
      missingFields,
    )
	}
	return nil
}

// setupKubernetesClients initializes Kubernetes clients for interacting with K8S API,
// dynamic resources, and CRD resources.
func (director *controllerDirector) setupKubernetesClients(controller *controller) error {

  // Build a configuration object for connecting to the Kubernetes API server
  // using the specified address and configuration file (if any).
  connectionConfig, err := clientcmd.BuildConfigFromFlags(
      director.builder.kubernetesAddress, 
      director.builder.kubernetesConfig,
  )
  if err != nil {
    return fmt.Errorf("failed to setup building Kubernetes connection object: %w", err)
  }

  // Instantiate a new client for interacting with K8S resources via API calls.
  controller.k8sClient, err = kubernetes.NewForConfig(connectionConfig)
  if err != nil {
    return fmt.Errorf("failed to create Kubernetes client: %w", err)
  }

  // Instantiate a new client for interacting with dynamic resources via API calls.
  controller.dynClient, err = dynamic.NewForConfig(connectionConfig)
  if err != nil {
    return fmt.Errorf("failed to create dynamic client: %w", err)
  }

  // Instantiate a new client for interacting with CRD resources via API calls.
  controller.crdClient, err = crdClientSet.NewForConfig(connectionConfig)
  if err != nil {
    return fmt.Errorf("failed to create CRD client: %w", err)
  }

  return nil
}

// setupDynamicInformer sets up dynamic informers for the specified resources,
// watches resources like pods, persistent volume claims, and config maps, and sets up event handlers.
func (director *controllerDirector) setupDynamicInformer(controller *controller) error {
    // Define GVRs (GroupVersionResources) for resources to watch
    resourcesToWatch := []schema.GroupVersionResource{
        {Group: "", Version: "v1", Resource: "pods"},
        {Group: "", Version: "v1", Resource: "persistentvolumeclaims"},
        {Group: "", Version: "v1", Resource: "configmaps"},
    }

    // Setup namespaceFilter
    namespaceFilter := "" 
    if director.builder.namespaceFilter != "" {
      namespaceFilter = director.builder.namespaceFilter
    }

    // Create a dynamic informer factory
    informerFactory := dynamicinformer.NewFilteredDynamicSharedInformerFactory(
        controller.dynClient,
        time.Minute,     // Resync period of 1 minute
        namespaceFilter, // No namespace filter, watch across all namespaces
        nil,             // No label selector filter
    )

    // Slice to store the dynamic informers
    dynInformers := []cache.SharedInformer{}

    // Create informers for each resource and set up event handlers
    for _, gvr := range resourcesToWatch {
        // Get the informer for the current resource type
        informer := informerFactory.ForResource(gvr).Informer()

        // Append the informer to the list of dynamic informers
        dynInformers = append(dynInformers, informer)

        // Set event handlers for add, update, and delete events
        informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
            // handleObject handles add events
            AddFunc: controller.handleObject,

            // UpdateFunc compares old and new objects to avoid processing if unchanged
            UpdateFunc: func(oldObj, newObj interface{}) {
                oldObject := oldObj.(*unstructured.Unstructured)
                newObject := newObj.(*unstructured.Unstructured)

                // If the resource version is the same, there's no meaningful update
                if oldObject.GetResourceVersion() == newObject.GetResourceVersion() {
                    return // No meaningful change; skip processing
                }

                // Handle the update event (e.g., call a method on the controller)
                controller.handleObject(newObject)
            },

            // handleObject handles delete events
            DeleteFunc: controller.handleObject,
        })

        // Start the informer asynchronously
        go informer.Run(controller.workingContext.Done())
    }

    // Store the informers in the controller
    controller.dynInformers = dynInformers

    return nil
}

// setupCustomResourceInformer sets up the informer for custom resources
// (overlays),
// adds event handlers for resource events (Add, Update, Delete), and starts the informer factory.
func (director *controllerDirector) setupCustomResourceInformer(controller *controller) error {
	// Create a shared informer factory with a 30-second resync period
	crdInformerFactory := crdInformeres.NewSharedInformerFactory(
		controller.crdClient,
		time.Second*5,
	)

	// Get the informer for the "Overlays" custom resource
	crdInformer := crdInformerFactory.Kubeforge().V1().Overlays()

	// Add event handler for custom resources (Overlays)
	crdInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: controller.enqueue,
		UpdateFunc: func(old, new interface{}) {
			controller.enqueue(new)
		},
		DeleteFunc: controller.enqueue,
	})

	// Pass the lister and sync checker to the controller
	controller.crdLister = crdInformer.Lister()
	controller.crdsSynced = crdInformer.Informer().HasSynced

	// Start the informer factory in the background and wait for it to sync
	go crdInformerFactory.Start(director.builder.workingContext.Done())

	return nil
}

// setupWorkQueue sets up the work queue
func (director *controllerDirector) setupWorkQueue(controller *controller) error {
	// Create a rate limiter for the work queue
	ratelimiter := workqueue.NewTypedMaxOfRateLimiter(
		workqueue.NewTypedItemExponentialFailureRateLimiter[cache.ObjectName](
			5*time.Millisecond,
			1000*time.Second,
		),
		&workqueue.TypedBucketRateLimiter[cache.ObjectName]{
			Limiter: rate.NewLimiter(
				rate.Limit(50),
				300,
			),
		},
	)
	if ratelimiter == nil {
		return fmt.Errorf("failed to create rate limiter")
	}

	// Create the rate-limiting queue
	workqueue := workqueue.NewTypedRateLimitingQueue(ratelimiter)
	if workqueue == nil {
		return fmt.Errorf("failed to initialize rate-limiting queue")
	}

	// Assign the workqueue to the controller
	controller.workqueue = workqueue
	return nil
}

// setupEventRecorder sets up the event recorder
func (director *controllerDirector) setupEventRecorder(controller *controller) error {

  // Add CRD scheme to the global scheme
  if err := crdScheme.AddToScheme(scheme.Scheme); err != nil {
      return fmt.Errorf("failed to add CRD scheme to scheme: %w", err)
  }

  // Create an event broadcaster
	eventBroadcaster := record.NewBroadcaster(
    record.WithContext(director.builder.workingContext),
  )

  // Start structured logging
	eventBroadcaster.StartStructuredLogging(0)

  // Start recording events to the Kubernetes sink
  eventBroadcaster.StartRecordingToSink(
    &v1.EventSinkImpl{
      Interface: controller.k8sClient.CoreV1().Events(""),
    },
  )

  // Create the event recorder
	recorder := 
    eventBroadcaster.NewRecorder(
      crdScheme.Scheme, 
      corev1.EventSource {
        Component: director.builder.controllerName,
      },
    )
  if recorder == nil {
    return fmt.Errorf("failed to create eventRecorder")
  }

  controller.recorder = recorder

  return nil
}
