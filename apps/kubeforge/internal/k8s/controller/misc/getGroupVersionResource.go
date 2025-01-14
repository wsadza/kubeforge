// ############################################################
// Copyright (c) 2024 wsadza
// Released under the MIT license
// ------------------------------------------------------------
//
//
// ############################################################

package controller

import (
	"fmt"
	"reflect"
	"strings"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
)

func GetGroupVersionResource(
  resource        interface{},
  discoveryClient discovery.DiscoveryInterface,
) (
  *schema.GroupVersionResource,
  error,
) {

  resourceType := reflect.TypeOf(resource)
  if (resourceType == nil || resourceType.Kind() != reflect.Ptr) && resourceType.Kind() != reflect.String {
		return nil, fmt.Errorf("invalid object type: expected pointer, got %v", resourceType)
  }

  var resourceKind string 
  if resourceType.Kind() == reflect.String {
    resourceKind = resource.(string)
  } else {
    resourceKind = resourceType.Elem().Name()
  }

  return lookupGroupVersionResource(resourceKind, discoveryClient)
}

func lookupGroupVersionResource(
  resourceKind    string,
  discoveryClient discovery.DiscoveryInterface,
) (
  *schema.GroupVersionResource,
  error,
) {

	// List all available API groups and versions in the cluster
	apiGroups, err := discoveryClient.ServerGroups()
	if err != nil {
		return nil, fmt.Errorf("failed to get API groups: %v", err)
	}

	// Iterate over each API group to find the resource's group and version
	for _, group := range apiGroups.Groups {
		for _, version := range group.Versions {
			apiResources, err := discoveryClient.ServerResourcesForGroupVersion(version.GroupVersion)
			if err != nil {
				return nil, fmt.Errorf("failed to get resources for group version %s: %v", version.GroupVersion, err)
			}

			// Search for the resource by kind or plural name
			for _, resource := range apiResources.APIResources {
				if strings.EqualFold(resource.Kind, resourceKind) || strings.EqualFold(resource.Name, resourceKind) {
					return &schema.GroupVersionResource{
						Group:    group.Name,
						Version:  version.Version,
						Resource: resource.Name,
					}, nil
				}
			}
		}
	}

	return nil, fmt.Errorf("no GVR found for resource: %s", resourceKind)
}
