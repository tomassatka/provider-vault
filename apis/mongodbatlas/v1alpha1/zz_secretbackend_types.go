/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SecretBackendInitParameters struct {

	// Path where the MongoDB Atlas Secrets Engine is mounted.
	// Path where MongoDB Atlas secret backend is mounted
	// +crossplane:generate:reference:type=github.com/upbound/provider-vault/v2/apis/vault/v1alpha1.Mount
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("path",false)
	Mount *string `json:"mount,omitempty" tf:"mount,omitempty"`

	// Reference to a Mount in vault to populate mount.
	// +kubebuilder:validation:Optional
	MountRef *v1.Reference `json:"mountRef,omitempty" tf:"-"`

	// Selector for a Mount in vault to populate mount.
	// +kubebuilder:validation:Optional
	MountSelector *v1.Selector `json:"mountSelector,omitempty" tf:"-"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Specifies the Private API Key used to authenticate with the MongoDB Atlas API.
	// The Private Programmatic API Key used to connect with MongoDB Atlas API
	PrivateKey *string `json:"privateKey,omitempty" tf:"private_key,omitempty"`

	// Specifies the Public API Key used to authenticate with the MongoDB Atlas API.
	// The Public Programmatic API Key used to authenticate with the MongoDB Atlas API
	PublicKey *string `json:"publicKey,omitempty" tf:"public_key,omitempty"`
}

type SecretBackendObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Path where the MongoDB Atlas Secrets Engine is mounted.
	// Path where MongoDB Atlas secret backend is mounted
	Mount *string `json:"mount,omitempty" tf:"mount,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Path where MongoDB Atlas configuration is located
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Specifies the Private API Key used to authenticate with the MongoDB Atlas API.
	// The Private Programmatic API Key used to connect with MongoDB Atlas API
	PrivateKey *string `json:"privateKey,omitempty" tf:"private_key,omitempty"`

	// Specifies the Public API Key used to authenticate with the MongoDB Atlas API.
	// The Public Programmatic API Key used to authenticate with the MongoDB Atlas API
	PublicKey *string `json:"publicKey,omitempty" tf:"public_key,omitempty"`
}

type SecretBackendParameters struct {

	// Path where the MongoDB Atlas Secrets Engine is mounted.
	// Path where MongoDB Atlas secret backend is mounted
	// +crossplane:generate:reference:type=github.com/upbound/provider-vault/v2/apis/vault/v1alpha1.Mount
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("path",false)
	// +kubebuilder:validation:Optional
	Mount *string `json:"mount,omitempty" tf:"mount,omitempty"`

	// Reference to a Mount in vault to populate mount.
	// +kubebuilder:validation:Optional
	MountRef *v1.Reference `json:"mountRef,omitempty" tf:"-"`

	// Selector for a Mount in vault to populate mount.
	// +kubebuilder:validation:Optional
	MountSelector *v1.Selector `json:"mountSelector,omitempty" tf:"-"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Specifies the Private API Key used to authenticate with the MongoDB Atlas API.
	// The Private Programmatic API Key used to connect with MongoDB Atlas API
	// +kubebuilder:validation:Optional
	PrivateKey *string `json:"privateKey,omitempty" tf:"private_key,omitempty"`

	// Specifies the Public API Key used to authenticate with the MongoDB Atlas API.
	// The Public Programmatic API Key used to authenticate with the MongoDB Atlas API
	// +kubebuilder:validation:Optional
	PublicKey *string `json:"publicKey,omitempty" tf:"public_key,omitempty"`
}

// SecretBackendSpec defines the desired state of SecretBackend
type SecretBackendSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecretBackendParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider SecretBackendInitParameters `json:"initProvider,omitempty"`
}

// SecretBackendStatus defines the observed state of SecretBackend.
type SecretBackendStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecretBackendObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SecretBackend is the Schema for the SecretBackends API. Creates a MongoDB Atlas secret backend for Vault.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type SecretBackend struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.privateKey) || (has(self.initProvider) && has(self.initProvider.privateKey))",message="spec.forProvider.privateKey is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.publicKey) || (has(self.initProvider) && has(self.initProvider.publicKey))",message="spec.forProvider.publicKey is a required parameter"
	Spec   SecretBackendSpec   `json:"spec"`
	Status SecretBackendStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecretBackendList contains a list of SecretBackends
type SecretBackendList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecretBackend `json:"items"`
}

// Repository type metadata.
var (
	SecretBackend_Kind             = "SecretBackend"
	SecretBackend_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecretBackend_Kind}.String()
	SecretBackend_KindAPIVersion   = SecretBackend_Kind + "." + CRDGroupVersion.String()
	SecretBackend_GroupVersionKind = CRDGroupVersion.WithKind(SecretBackend_Kind)
)

func init() {
	SchemeBuilder.Register(&SecretBackend{}, &SecretBackendList{})
}
