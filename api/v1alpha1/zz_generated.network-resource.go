// Code generated by resource-generator. DO NOT EDIT.
/*
Copyright 2024 The ORC Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NetworkImport specifies an existing resource which will be imported instead of
// creating a new one
// +kubebuilder:validation:MinProperties:=1
// +kubebuilder:validation:MaxProperties:=1
type NetworkImport struct {
	// id contains the unique identifier of an existing OpenStack resource. Note
	// that when specifying an import by ID, the resource MUST already exist.
	// The ORC object will enter an error state if the resource does not exist.
	// +optional
	// +kubebuilder:validation:Format:=uuid
	ID *string `json:"id,omitempty"`

	// filter contains a resource query which is expected to return a single
	// result. The controller will continue to retry if filter returns no
	// results. If filter returns multiple results the controller will set an
	// error state and will not continue to retry.
	// +optional
	Filter *NetworkFilter `json:"filter,omitempty"`
}

// NetworkSpec defines the desired state of an ORC object.
// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'managed' ? has(self.resource) : true",message="resource must be specified when policy is managed"
// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'managed' ? !has(self.__import__) : true",message="import may not be specified when policy is managed"
// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'unmanaged' ? !has(self.resource) : true",message="resource may not be specified when policy is unmanaged"
// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'unmanaged' ? has(self.__import__) : true",message="import must be specified when policy is unmanaged"
// +kubebuilder:validation:XValidation:rule="has(self.managedOptions) ? self.managementPolicy == 'managed' : true",message="managedOptions may only be provided when policy is managed"
type NetworkSpec struct {
	// import refers to an existing OpenStack resource which will be imported instead of
	// creating a new one.
	// +optional
	Import *NetworkImport `json:"import,omitempty"`

	// resource specifies the desired state of the resource.
	//
	// resource may not be specified if the management policy is `unmanaged`.
	//
	// resource must be specified if the management policy is `managed`.
	// +optional
	Resource *NetworkResourceSpec `json:"resource,omitempty"`

	// managementPolicy defines how ORC will treat the object. Valid values are
	// `managed`: ORC will create, update, and delete the resource; `unmanaged`:
	// ORC will import an existing resource, and will not apply updates to it or
	// delete it.
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="managementPolicy is immutable"
	// +kubebuilder:default:=managed
	// +optional
	ManagementPolicy ManagementPolicy `json:"managementPolicy,omitempty"`

	// managedOptions specifies options which may be applied to managed objects.
	// +optional
	ManagedOptions *ManagedOptions `json:"managedOptions,omitempty"`

	// cloudCredentialsRef points to a secret containing OpenStack credentials
	// +required
	CloudCredentialsRef CloudCredentialsReference `json:"cloudCredentialsRef"`
}

// NetworkStatus defines the observed state of an ORC resource.
type NetworkStatus struct {
	// conditions represents the observed status of the object.
	// Known .status.conditions.type are: "Available", "Progressing"
	//
	// Available represents the availability of the OpenStack resource. If it is
	// true then the resource is ready for use.
	//
	// Progressing indicates whether the controller is still attempting to
	// reconcile the current state of the OpenStack resource to the desired
	// state. Progressing will be False either because the desired state has
	// been achieved, or because some terminal error prevents it from ever being
	// achieved and the controller is no longer attempting to reconcile. If
	// Progressing is True, an observer waiting on the resource should continue
	// to wait.
	//
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=type
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`

	// id is the unique identifier of the OpenStack resource.
	// +optional
	ID *string `json:"id,omitempty"`

	// resource contains the observed state of the OpenStack resource.
	// +optional
	Resource *NetworkResourceStatus `json:"resource,omitempty"`
}

var _ ObjectWithConditions = &Network{}

func (i *Network) GetConditions() []metav1.Condition {
	return i.Status.Conditions
}

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:resource:categories=openstack
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="ID",type="string",JSONPath=".status.id",description="Resource ID"
// +kubebuilder:printcolumn:name="Available",type="string",JSONPath=".status.conditions[?(@.type=='Available')].status",description="Availability status of resource"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Available')].message",description="Message describing current availability status"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp",description="Time duration since creation"

// Network is the Schema for an ORC resource.
type Network struct {
	metav1.TypeMeta   `json:",inline"`

	// metadata contains the object metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// spec specifies the desired state of the resource.
	// +optional
	Spec   NetworkSpec   `json:"spec,omitempty"`

	// status defines the observed state of the resource.
	// +optional
	Status NetworkStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkList contains a list of Network.
type NetworkList struct {
	metav1.TypeMeta `json:",inline"`

	// metadata contains the list metadata
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	// items contains a list of Network.
	// +required
	Items           []Network `json:"items"`
}

func (l *NetworkList) GetItems() []Network {
	return l.Items
}

func init() {
	SchemeBuilder.Register(&Network{}, &NetworkList{})
}

func (i *Network) GetCloudCredentialsRef() (*string, *CloudCredentialsReference) {
	if i == nil {
		return nil, nil
	}

	return &i.Namespace, &i.Spec.CloudCredentialsRef
}

var _ CloudCredentialsRefProvider = &Network{}
