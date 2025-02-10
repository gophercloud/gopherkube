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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	apiv1alpha1 "github.com/k-orc/openstack-resource-controller/api/v1alpha1"
)

// ImageContentApplyConfiguration represents a declarative configuration of the ImageContent type for use
// with apply.
type ImageContentApplyConfiguration struct {
	ContainerFormat *apiv1alpha1.ImageContainerFormat             `json:"containerFormat,omitempty"`
	DiskFormat      *apiv1alpha1.ImageDiskFormat                  `json:"diskFormat,omitempty"`
	Download        *ImageContentSourceDownloadApplyConfiguration `json:"download,omitempty"`
}

// ImageContentApplyConfiguration constructs a declarative configuration of the ImageContent type for use with
// apply.
func ImageContent() *ImageContentApplyConfiguration {
	return &ImageContentApplyConfiguration{}
}

// WithContainerFormat sets the ContainerFormat field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ContainerFormat field is set to the value of the last call.
func (b *ImageContentApplyConfiguration) WithContainerFormat(value apiv1alpha1.ImageContainerFormat) *ImageContentApplyConfiguration {
	b.ContainerFormat = &value
	return b
}

// WithDiskFormat sets the DiskFormat field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DiskFormat field is set to the value of the last call.
func (b *ImageContentApplyConfiguration) WithDiskFormat(value apiv1alpha1.ImageDiskFormat) *ImageContentApplyConfiguration {
	b.DiskFormat = &value
	return b
}

// WithDownload sets the Download field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Download field is set to the value of the last call.
func (b *ImageContentApplyConfiguration) WithDownload(value *ImageContentSourceDownloadApplyConfiguration) *ImageContentApplyConfiguration {
	b.Download = value
	return b
}
