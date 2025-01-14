// ############################################################
// Copyright (c) 2024 wsadza
// Released under the MIT license
// ------------------------------------------------------------
//
//
// ############################################################

package v1 

import (
  "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ------------------------------------------------------------
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Overlay is a specification for a Overlay resource
type Overlay struct {
	runtime.TypeMeta `json:",inline"`
	v1.ObjectMeta    `json:"metadata,omitempty"`

	Spec   OverlaySpec   `json:"spec"`
	Status OverlayStatus `json:"status"`
}

// OverlaySpec is the spec for a Overlay resource
type OverlaySpec struct {
  Data runtime.RawExtension `json:"data,omitempty"`
}

// OverlayStatus is the status for a Overlay resource
type OverlayStatus struct {
  Data runtime.RawExtension `json:"data,omitempty"`
}

// ------------------------------------------------------------
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OverlayList is a list of Overlay resources
type OverlayList struct {
	v1.TypeMeta `json:",inline"`
	v1.ListMeta `json:"metadata"`

	Items []Overlay `json:"items"`
}
