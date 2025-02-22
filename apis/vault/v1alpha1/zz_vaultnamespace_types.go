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

type VaultNamespaceInitParameters struct {

	// Custom metadata describing this namespace. Value type
	// is map[string]string. Requires Vault version 1.12+.
	// Custom metadata describing this namespace. Value type is map[string]string.
	// +mapType=granular
	CustomMetadata map[string]*string `json:"customMetadata,omitempty" tf:"custom_metadata,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	// +crossplane:generate:reference:type=github.com/upbound/provider-vault/v2/apis/vault/v1alpha1.VaultNamespace
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("path",false)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Reference to a VaultNamespace in vault to populate namespace.
	// +kubebuilder:validation:Optional
	NamespaceRef *v1.Reference `json:"namespaceRef,omitempty" tf:"-"`

	// Selector for a VaultNamespace in vault to populate namespace.
	// +kubebuilder:validation:Optional
	NamespaceSelector *v1.Selector `json:"namespaceSelector,omitempty" tf:"-"`

	// The path of the namespace. Must not have a trailing /.
	// Namespace path.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// The fully qualified path to the namespace. Useful when provisioning resources in a child namespace.
	// The path is relative to the provider's namespace argument.
	// The fully qualified namespace path.
	PathFq *string `json:"pathFq,omitempty" tf:"path_fq,omitempty"`
}

type VaultNamespaceObservation struct {

	// Custom metadata describing this namespace. Value type
	// is map[string]string. Requires Vault version 1.12+.
	// Custom metadata describing this namespace. Value type is map[string]string.
	// +mapType=granular
	CustomMetadata map[string]*string `json:"customMetadata,omitempty" tf:"custom_metadata,omitempty"`

	// The fully qualified path to the namespace, including the provider namespace and a trailing slash.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Vault server's internal ID of the namespace.
	// Namespace ID.
	NamespaceID *string `json:"namespaceId,omitempty" tf:"namespace_id,omitempty"`

	// The path of the namespace. Must not have a trailing /.
	// Namespace path.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// The fully qualified path to the namespace. Useful when provisioning resources in a child namespace.
	// The path is relative to the provider's namespace argument.
	// The fully qualified namespace path.
	PathFq *string `json:"pathFq,omitempty" tf:"path_fq,omitempty"`
}

type VaultNamespaceParameters struct {

	// Custom metadata describing this namespace. Value type
	// is map[string]string. Requires Vault version 1.12+.
	// Custom metadata describing this namespace. Value type is map[string]string.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	CustomMetadata map[string]*string `json:"customMetadata,omitempty" tf:"custom_metadata,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	// +crossplane:generate:reference:type=github.com/upbound/provider-vault/v2/apis/vault/v1alpha1.VaultNamespace
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("path",false)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Reference to a VaultNamespace in vault to populate namespace.
	// +kubebuilder:validation:Optional
	NamespaceRef *v1.Reference `json:"namespaceRef,omitempty" tf:"-"`

	// Selector for a VaultNamespace in vault to populate namespace.
	// +kubebuilder:validation:Optional
	NamespaceSelector *v1.Selector `json:"namespaceSelector,omitempty" tf:"-"`

	// The path of the namespace. Must not have a trailing /.
	// Namespace path.
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// The fully qualified path to the namespace. Useful when provisioning resources in a child namespace.
	// The path is relative to the provider's namespace argument.
	// The fully qualified namespace path.
	// +kubebuilder:validation:Optional
	PathFq *string `json:"pathFq,omitempty" tf:"path_fq,omitempty"`
}

// VaultNamespaceSpec defines the desired state of VaultNamespace
type VaultNamespaceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VaultNamespaceParameters `json:"forProvider"`
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
	InitProvider VaultNamespaceInitParameters `json:"initProvider,omitempty"`
}

// VaultNamespaceStatus defines the observed state of VaultNamespace.
type VaultNamespaceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VaultNamespaceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// VaultNamespace is the Schema for the VaultNamespaces API. Writes namespaces for Vault
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type VaultNamespace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.path) || (has(self.initProvider) && has(self.initProvider.path))",message="spec.forProvider.path is a required parameter"
	Spec   VaultNamespaceSpec   `json:"spec"`
	Status VaultNamespaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VaultNamespaceList contains a list of VaultNamespaces
type VaultNamespaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VaultNamespace `json:"items"`
}

// Repository type metadata.
var (
	VaultNamespace_Kind             = "VaultNamespace"
	VaultNamespace_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VaultNamespace_Kind}.String()
	VaultNamespace_KindAPIVersion   = VaultNamespace_Kind + "." + CRDGroupVersion.String()
	VaultNamespace_GroupVersionKind = CRDGroupVersion.WithKind(VaultNamespace_Kind)
)

func init() {
	SchemeBuilder.Register(&VaultNamespace{}, &VaultNamespaceList{})
}
