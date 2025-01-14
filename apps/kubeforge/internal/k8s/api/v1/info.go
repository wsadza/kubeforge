// ############################################################
// Copyright (c) 2024 wsadza
// Released under the MIT license
// ------------------------------------------------------------
//
//
// ############################################################

// ------------------------------------------------------------
// +kubebuilder:object:generate=true
// +groupName=metal

package v1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	// SchemeGroupVersion is group version used to register these objects
	SchemeGroupVersion = schema.GroupVersion{
    Group:    "kubeforge.sh", 
    Version:  "v1",
  }
)
