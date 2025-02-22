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

type SecretBackendConfigCAInitParameters struct {

	// The PKI secret backend the resource belongs to.
	// The PKI secret backend the resource belongs to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-vault/v2/apis/vault/v1alpha1.Mount
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("path",false)
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// Reference to a Mount in vault to populate backend.
	// +kubebuilder:validation:Optional
	BackendRef *v1.Reference `json:"backendRef,omitempty" tf:"-"`

	// Selector for a Mount in vault to populate backend.
	// +kubebuilder:validation:Optional
	BackendSelector *v1.Selector `json:"backendSelector,omitempty" tf:"-"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// The key and certificate PEM bundle
	// The key and certificate PEM bundle.
	PemBundleSecretRef v1.SecretKeySelector `json:"pemBundleSecretRef" tf:"-"`
}

type SecretBackendConfigCAObservation struct {

	// The PKI secret backend the resource belongs to.
	// The PKI secret backend the resource belongs to.
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`
}

type SecretBackendConfigCAParameters struct {

	// The PKI secret backend the resource belongs to.
	// The PKI secret backend the resource belongs to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-vault/v2/apis/vault/v1alpha1.Mount
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("path",false)
	// +kubebuilder:validation:Optional
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// Reference to a Mount in vault to populate backend.
	// +kubebuilder:validation:Optional
	BackendRef *v1.Reference `json:"backendRef,omitempty" tf:"-"`

	// Selector for a Mount in vault to populate backend.
	// +kubebuilder:validation:Optional
	BackendSelector *v1.Selector `json:"backendSelector,omitempty" tf:"-"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// The key and certificate PEM bundle
	// The key and certificate PEM bundle.
	// +kubebuilder:validation:Optional
	PemBundleSecretRef v1.SecretKeySelector `json:"pemBundleSecretRef" tf:"-"`
}

// SecretBackendConfigCASpec defines the desired state of SecretBackendConfigCA
type SecretBackendConfigCASpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecretBackendConfigCAParameters `json:"forProvider"`
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
	InitProvider SecretBackendConfigCAInitParameters `json:"initProvider,omitempty"`
}

// SecretBackendConfigCAStatus defines the observed state of SecretBackendConfigCA.
type SecretBackendConfigCAStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecretBackendConfigCAObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SecretBackendConfigCA is the Schema for the SecretBackendConfigCAs API. Submit the CA information to PKI.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type SecretBackendConfigCA struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.pemBundleSecretRef)",message="spec.forProvider.pemBundleSecretRef is a required parameter"
	Spec   SecretBackendConfigCASpec   `json:"spec"`
	Status SecretBackendConfigCAStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecretBackendConfigCAList contains a list of SecretBackendConfigCAs
type SecretBackendConfigCAList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecretBackendConfigCA `json:"items"`
}

// Repository type metadata.
var (
	SecretBackendConfigCA_Kind             = "SecretBackendConfigCA"
	SecretBackendConfigCA_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecretBackendConfigCA_Kind}.String()
	SecretBackendConfigCA_KindAPIVersion   = SecretBackendConfigCA_Kind + "." + CRDGroupVersion.String()
	SecretBackendConfigCA_GroupVersionKind = CRDGroupVersion.WithKind(SecretBackendConfigCA_Kind)
)

func init() {
	SchemeBuilder.Register(&SecretBackendConfigCA{}, &SecretBackendConfigCAList{})
}
