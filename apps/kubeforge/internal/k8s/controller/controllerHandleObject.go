// ############################################################
// Copyright (c) 2024 wsadza
// Released under the MIT license
// ------------------------------------------------------------
//
// The `handleObject` function processes incoming resources, checking
// for appropriate OwnerReferences to determine if they are part of
// the desired CRD. If so, it enqueues the associated Overlay for
// further handling. The `enqueue` function converts the CRD resource
// into a namespace/name string and adds it to the work queue for processing.
//
// ############################################################

package controller

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog/v2"
)

// handleObject will take any resource implementing metav1.Object and attempt
// to find the CRD resource that 'owns' it. It does this by looking at the
// objects metadata.ownerReferences field for an appropriate OwnerReference.
// It then enqueues that Foo resource to be processed. If the object does not
// have an appropriate OwnerReference, it will simply be skipped.
func (controller *controller) handleObject(obj interface{}) {
	var object metav1.Object
	var ok bool
	logger := klog.FromContext(context.Background())

	if object, ok = obj.(metav1.Object); !ok {
		tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
    logger.Info("Processing tombstone", "object", tombstone.Obj)
		if !ok {
			// If the object value is not too big and does not contain sensitive information then
			// it may be useful to include it.
			runtime.HandleErrorWithContext(
        context.Background(), 
        nil, 
        "Error decoding object, invalid type", 
        "type", 
        fmt.Sprintf("%T", obj),
      )
			return
		}
		object, ok = tombstone.Obj.(metav1.Object)
		if !ok {
			// If the object value is not too big and does not contain sensitive information then
			// it may be useful to include it.
			runtime.HandleErrorWithContext(
        context.Background(), 
        nil, 
        "Error decoding object tombstone, invalid type", 
        "type", 
        fmt.Sprintf("%T", tombstone.Obj),
      )
			return
		}
		logger.Info("Recovered deleted object", "resourceName", object.GetName())
	}

	logger.Info("Processing object", "object", klog.KObj(object))
	if ownerRef := metav1.GetControllerOf(object); ownerRef != nil {
		// If this object is not owned by a Overlay, we should not do anything more with it.
		if ownerRef.Kind != "Overlay" {
			return
		}

		instance, err := 
      controller.
        crdLister.
        Overlays(object.GetNamespace()).
        Get(ownerRef.Name)

		if err != nil {
			logger.V(4).Info(
        "Ignore orphaned object", 
        "object", 
        klog.KObj(object), 
        "instance", 
        ownerRef.Name,
      )
			return
		}

		controller.enqueue(instance)
		return
	}
}

// enqueue takes a CRD resource and converts it into a namespace/name
// string which is then put onto the work queue. This method should *not* be
// passed resources of any type other than CRD.
func (controller *controller) enqueue (obj interface{}) {
  objectRef, err := cache.ObjectToName(obj); 
  if err != nil {
    runtime.HandleError(err)
    return
  } else { 
    controller.workqueue.Add(objectRef)
  }
}
