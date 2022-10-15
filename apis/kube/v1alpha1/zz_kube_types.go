/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AdmissionpluginsObservation struct {
}

type AdmissionpluginsParameters struct {

	// +kubebuilder:validation:Optional
	Disabled []*string `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// +kubebuilder:validation:Optional
	Enabled []*string `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type ApiserverObservation struct {
}

type ApiserverParameters struct {

	// +kubebuilder:validation:Optional
	Admissionplugins []AdmissionpluginsParameters `json:"admissionplugins,omitempty" tf:"admissionplugins,omitempty"`
}

type CustomizationObservation struct {
}

type CustomizationParameters struct {

	// +kubebuilder:validation:Optional
	Apiserver []ApiserverParameters `json:"apiserver,omitempty" tf:"apiserver,omitempty"`
}

type KubeObservation struct {
	ControlPlaneIsUpToDate *bool `json:"controlPlaneIsUpToDate,omitempty" tf:"control_plane_is_up_to_date,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	IsUpToDate *bool `json:"isUpToDate,omitempty" tf:"is_up_to_date,omitempty"`

	NextUpgradeVersions []*string `json:"nextUpgradeVersions,omitempty" tf:"next_upgrade_versions,omitempty"`

	NodesURL *string `json:"nodesUrl,omitempty" tf:"nodes_url,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type KubeParameters struct {

	// +kubebuilder:validation:Optional
	Customization []CustomizationParameters `json:"customization,omitempty" tf:"customization,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	PrivateNetworkConfiguration []PrivateNetworkConfigurationParameters `json:"privateNetworkConfiguration,omitempty" tf:"private_network_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	PrivateNetworkID *string `json:"privateNetworkId,omitempty" tf:"private_network_id,omitempty"`

	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`

	// +kubebuilder:validation:Required
	ServiceName *string `json:"serviceName" tf:"service_name,omitempty"`

	// +kubebuilder:validation:Optional
	UpdatePolicy *string `json:"updatePolicy,omitempty" tf:"update_policy,omitempty"`

	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type PrivateNetworkConfigurationObservation struct {
}

type PrivateNetworkConfigurationParameters struct {

	// If defined, all egress traffic will be routed towards this IP address, which should belong to the private network. Empty string means disabled.
	// +kubebuilder:validation:Required
	DefaultVrackGateway *string `json:"defaultVrackGateway" tf:"default_vrack_gateway,omitempty"`

	// Defines whether routing should default to using the nodes' private interface, instead of their public interface. Default is false.
	// +kubebuilder:validation:Required
	PrivateNetworkRoutingAsDefault *bool `json:"privateNetworkRoutingAsDefault" tf:"private_network_routing_as_default,omitempty"`
}

// KubeSpec defines the desired state of Kube
type KubeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     KubeParameters `json:"forProvider"`
}

// KubeStatus defines the observed state of Kube.
type KubeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        KubeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Kube is the Schema for the Kubes API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ovhjet}
type Kube struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KubeSpec   `json:"spec"`
	Status            KubeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KubeList contains a list of Kubes
type KubeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Kube `json:"items"`
}

// Repository type metadata.
var (
	Kube_Kind             = "Kube"
	Kube_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Kube_Kind}.String()
	Kube_KindAPIVersion   = Kube_Kind + "." + CRDGroupVersion.String()
	Kube_GroupVersionKind = CRDGroupVersion.WithKind(Kube_Kind)
)

func init() {
	SchemeBuilder.Register(&Kube{}, &KubeList{})
}