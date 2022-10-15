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

type DatabaseObservation struct {
	BackupTime *string `json:"backupTime,omitempty" tf:"backup_time,omitempty"`

	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	Endpoints []EndpointsObservation `json:"endpoints,omitempty" tf:"endpoints,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	MaintenanceTime *string `json:"maintenanceTime,omitempty" tf:"maintenance_time,omitempty"`

	NetworkType *string `json:"networkType,omitempty" tf:"network_type,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type DatabaseParameters struct {

	// Description of the cluster
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Name of the engine of the service
	// +kubebuilder:validation:Required
	Engine *string `json:"engine" tf:"engine,omitempty"`

	// The node flavor used for this cluster
	// +kubebuilder:validation:Required
	Flavor *string `json:"flavor" tf:"flavor,omitempty"`

	// Defines whether the REST API is enabled on a Kafka cluster
	// +kubebuilder:validation:Optional
	KafkaRestAPI *bool `json:"kafkaRestApi,omitempty" tf:"kafka_rest_api,omitempty"`

	// List of nodes composing the service
	// +kubebuilder:validation:Required
	Nodes []NodesParameters `json:"nodes" tf:"nodes,omitempty"`

	// Defines whether the ACLs are enabled on an Opensearch cluster
	// +kubebuilder:validation:Optional
	OpensearchAclsEnabled *bool `json:"opensearchAclsEnabled,omitempty" tf:"opensearch_acls_enabled,omitempty"`

	// Plan of the cluster
	// +kubebuilder:validation:Required
	Plan *string `json:"plan" tf:"plan,omitempty"`

	// +kubebuilder:validation:Required
	ServiceName *string `json:"serviceName" tf:"service_name,omitempty"`

	// Version of the engine deployed on the cluster
	// +kubebuilder:validation:Required
	Version *string `json:"version" tf:"version,omitempty"`
}

type EndpointsObservation struct {
	Component *string `json:"component,omitempty" tf:"component,omitempty"`

	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	SSL *bool `json:"ssl,omitempty" tf:"ssl,omitempty"`

	SSLMode *string `json:"sslMode,omitempty" tf:"ssl_mode,omitempty"`

	Scheme *string `json:"scheme,omitempty" tf:"scheme,omitempty"`

	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`
}

type EndpointsParameters struct {
}

type NodesObservation struct {
}

type NodesParameters struct {

	// Private network ID in which the node is. It's the regional openstackId of the private network.
	// +kubebuilder:validation:Optional
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Region of the node
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`

	// Private subnet ID in which the node is
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

// DatabaseSpec defines the desired state of Database
type DatabaseSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DatabaseParameters `json:"forProvider"`
}

// DatabaseStatus defines the observed state of Database.
type DatabaseStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DatabaseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Database is the Schema for the Databases API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ovhjet}
type Database struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseSpec   `json:"spec"`
	Status            DatabaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatabaseList contains a list of Databases
type DatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Database `json:"items"`
}

// Repository type metadata.
var (
	Database_Kind             = "Database"
	Database_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Database_Kind}.String()
	Database_KindAPIVersion   = Database_Kind + "." + CRDGroupVersion.String()
	Database_GroupVersionKind = CRDGroupVersion.WithKind(Database_Kind)
)

func init() {
	SchemeBuilder.Register(&Database{}, &DatabaseList{})
}