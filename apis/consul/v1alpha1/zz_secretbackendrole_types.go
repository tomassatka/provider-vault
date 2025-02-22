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

type SecretBackendRoleInitParameters struct {

	// The unique name of an existing Consul secrets backend mount. Must not begin or end with a /. One of path or backend is required.
	// The path of the Consul Secret Backend the role belongs to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-vault/v2/apis/consul/v1alpha1.SecretBackend
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("path",false)
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// Reference to a SecretBackend in consul to populate backend.
	// +kubebuilder:validation:Optional
	BackendRef *v1.Reference `json:"backendRef,omitempty" tf:"-"`

	// Selector for a SecretBackend in consul to populate backend.
	// +kubebuilder:validation:Optional
	BackendSelector *v1.Selector `json:"backendSelector,omitempty" tf:"-"`

	// The Consul namespace that the token will be created in.
	// Applicable for Vault 1.10+ and Consul 1.7+".
	// The Consul namespace that the token will be created in. Applicable for Vault 1.10+ and Consul 1.7+
	ConsulNamespace *string `json:"consulNamespace,omitempty" tf:"consul_namespace,omitempty"`

	// SEE NOTE The list of Consul ACL policies to associate with these roles.
	// List of Consul policies to associate with this role
	// +listType=set
	ConsulPolicies []*string `json:"consulPolicies,omitempty" tf:"consul_policies,omitempty"`

	// SEE NOTE Set of Consul roles to attach to the token.
	// Applicable for Vault 1.10+ with Consul 1.5+.
	// Set of Consul roles to attach to the token. Applicable for Vault 1.10+ with Consul 1.5+
	// +listType=set
	ConsulRoles []*string `json:"consulRoles,omitempty" tf:"consul_roles,omitempty"`

	// Indicates that the token should not be replicated globally and instead be local to the current datacenter.
	// Indicates that the token should not be replicated globally and instead be local to the current datacenter.
	Local *bool `json:"local,omitempty" tf:"local,omitempty"`

	// Maximum TTL for leases associated with this role, in seconds.
	// Maximum TTL for leases associated with this role, in seconds.
	MaxTTL *float64 `json:"maxTtl,omitempty" tf:"max_ttl,omitempty"`

	// The name of the Consul secrets engine role to create.
	// The name of an existing role against which to create this Consul credential
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// SEE NOTE Set of Consul node
	// identities to attach to the token. Applicable for Vault 1.11+ with Consul 1.8+.
	// Set of Consul node identities to attach to
	// the token. Applicable for Vault 1.11+ with Consul 1.8+
	// +listType=set
	NodeIdentities []*string `json:"nodeIdentities,omitempty" tf:"node_identities,omitempty"`

	// The admin partition that the token will be created in.
	// Applicable for Vault 1.10+ and Consul 1.11+".
	// The Consul admin partition that the token will be created in. Applicable for Vault 1.10+ and Consul 1.11+
	Partition *string `json:"partition,omitempty" tf:"partition,omitempty"`

	// The list of Consul ACL policies to associate with these roles.
	// NOTE: The new parameter consul_policies should be used in favor of this. This parameter,
	// policies, remains supported for legacy users, but Vault has deprecated this field.
	// List of Consul policies to associate with this role
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`

	// SEE NOTE Set of Consul
	// service identities to attach to the token. Applicable for Vault 1.11+ with Consul 1.5+.
	// Set of Consul service identities to attach to
	// the token. Applicable for Vault 1.11+ with Consul 1.5+
	// +listType=set
	ServiceIdentities []*string `json:"serviceIdentities,omitempty" tf:"service_identities,omitempty"`

	// Specifies the TTL for this role.
	// Specifies the TTL for this role.
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`
}

type SecretBackendRoleObservation struct {

	// The unique name of an existing Consul secrets backend mount. Must not begin or end with a /. One of path or backend is required.
	// The path of the Consul Secret Backend the role belongs to.
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// The Consul namespace that the token will be created in.
	// Applicable for Vault 1.10+ and Consul 1.7+".
	// The Consul namespace that the token will be created in. Applicable for Vault 1.10+ and Consul 1.7+
	ConsulNamespace *string `json:"consulNamespace,omitempty" tf:"consul_namespace,omitempty"`

	// SEE NOTE The list of Consul ACL policies to associate with these roles.
	// List of Consul policies to associate with this role
	// +listType=set
	ConsulPolicies []*string `json:"consulPolicies,omitempty" tf:"consul_policies,omitempty"`

	// SEE NOTE Set of Consul roles to attach to the token.
	// Applicable for Vault 1.10+ with Consul 1.5+.
	// Set of Consul roles to attach to the token. Applicable for Vault 1.10+ with Consul 1.5+
	// +listType=set
	ConsulRoles []*string `json:"consulRoles,omitempty" tf:"consul_roles,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Indicates that the token should not be replicated globally and instead be local to the current datacenter.
	// Indicates that the token should not be replicated globally and instead be local to the current datacenter.
	Local *bool `json:"local,omitempty" tf:"local,omitempty"`

	// Maximum TTL for leases associated with this role, in seconds.
	// Maximum TTL for leases associated with this role, in seconds.
	MaxTTL *float64 `json:"maxTtl,omitempty" tf:"max_ttl,omitempty"`

	// The name of the Consul secrets engine role to create.
	// The name of an existing role against which to create this Consul credential
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// SEE NOTE Set of Consul node
	// identities to attach to the token. Applicable for Vault 1.11+ with Consul 1.8+.
	// Set of Consul node identities to attach to
	// the token. Applicable for Vault 1.11+ with Consul 1.8+
	// +listType=set
	NodeIdentities []*string `json:"nodeIdentities,omitempty" tf:"node_identities,omitempty"`

	// The admin partition that the token will be created in.
	// Applicable for Vault 1.10+ and Consul 1.11+".
	// The Consul admin partition that the token will be created in. Applicable for Vault 1.10+ and Consul 1.11+
	Partition *string `json:"partition,omitempty" tf:"partition,omitempty"`

	// The list of Consul ACL policies to associate with these roles.
	// NOTE: The new parameter consul_policies should be used in favor of this. This parameter,
	// policies, remains supported for legacy users, but Vault has deprecated this field.
	// List of Consul policies to associate with this role
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`

	// SEE NOTE Set of Consul
	// service identities to attach to the token. Applicable for Vault 1.11+ with Consul 1.5+.
	// Set of Consul service identities to attach to
	// the token. Applicable for Vault 1.11+ with Consul 1.5+
	// +listType=set
	ServiceIdentities []*string `json:"serviceIdentities,omitempty" tf:"service_identities,omitempty"`

	// Specifies the TTL for this role.
	// Specifies the TTL for this role.
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`
}

type SecretBackendRoleParameters struct {

	// The unique name of an existing Consul secrets backend mount. Must not begin or end with a /. One of path or backend is required.
	// The path of the Consul Secret Backend the role belongs to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-vault/v2/apis/consul/v1alpha1.SecretBackend
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("path",false)
	// +kubebuilder:validation:Optional
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// Reference to a SecretBackend in consul to populate backend.
	// +kubebuilder:validation:Optional
	BackendRef *v1.Reference `json:"backendRef,omitempty" tf:"-"`

	// Selector for a SecretBackend in consul to populate backend.
	// +kubebuilder:validation:Optional
	BackendSelector *v1.Selector `json:"backendSelector,omitempty" tf:"-"`

	// The Consul namespace that the token will be created in.
	// Applicable for Vault 1.10+ and Consul 1.7+".
	// The Consul namespace that the token will be created in. Applicable for Vault 1.10+ and Consul 1.7+
	// +kubebuilder:validation:Optional
	ConsulNamespace *string `json:"consulNamespace,omitempty" tf:"consul_namespace,omitempty"`

	// SEE NOTE The list of Consul ACL policies to associate with these roles.
	// List of Consul policies to associate with this role
	// +kubebuilder:validation:Optional
	// +listType=set
	ConsulPolicies []*string `json:"consulPolicies,omitempty" tf:"consul_policies,omitempty"`

	// SEE NOTE Set of Consul roles to attach to the token.
	// Applicable for Vault 1.10+ with Consul 1.5+.
	// Set of Consul roles to attach to the token. Applicable for Vault 1.10+ with Consul 1.5+
	// +kubebuilder:validation:Optional
	// +listType=set
	ConsulRoles []*string `json:"consulRoles,omitempty" tf:"consul_roles,omitempty"`

	// Indicates that the token should not be replicated globally and instead be local to the current datacenter.
	// Indicates that the token should not be replicated globally and instead be local to the current datacenter.
	// +kubebuilder:validation:Optional
	Local *bool `json:"local,omitempty" tf:"local,omitempty"`

	// Maximum TTL for leases associated with this role, in seconds.
	// Maximum TTL for leases associated with this role, in seconds.
	// +kubebuilder:validation:Optional
	MaxTTL *float64 `json:"maxTtl,omitempty" tf:"max_ttl,omitempty"`

	// The name of the Consul secrets engine role to create.
	// The name of an existing role against which to create this Consul credential
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// SEE NOTE Set of Consul node
	// identities to attach to the token. Applicable for Vault 1.11+ with Consul 1.8+.
	// Set of Consul node identities to attach to
	// the token. Applicable for Vault 1.11+ with Consul 1.8+
	// +kubebuilder:validation:Optional
	// +listType=set
	NodeIdentities []*string `json:"nodeIdentities,omitempty" tf:"node_identities,omitempty"`

	// The admin partition that the token will be created in.
	// Applicable for Vault 1.10+ and Consul 1.11+".
	// The Consul admin partition that the token will be created in. Applicable for Vault 1.10+ and Consul 1.11+
	// +kubebuilder:validation:Optional
	Partition *string `json:"partition,omitempty" tf:"partition,omitempty"`

	// The list of Consul ACL policies to associate with these roles.
	// NOTE: The new parameter consul_policies should be used in favor of this. This parameter,
	// policies, remains supported for legacy users, but Vault has deprecated this field.
	// List of Consul policies to associate with this role
	// +kubebuilder:validation:Optional
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`

	// SEE NOTE Set of Consul
	// service identities to attach to the token. Applicable for Vault 1.11+ with Consul 1.5+.
	// Set of Consul service identities to attach to
	// the token. Applicable for Vault 1.11+ with Consul 1.5+
	// +kubebuilder:validation:Optional
	// +listType=set
	ServiceIdentities []*string `json:"serviceIdentities,omitempty" tf:"service_identities,omitempty"`

	// Specifies the TTL for this role.
	// Specifies the TTL for this role.
	// +kubebuilder:validation:Optional
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`
}

// SecretBackendRoleSpec defines the desired state of SecretBackendRole
type SecretBackendRoleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecretBackendRoleParameters `json:"forProvider"`
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
	InitProvider SecretBackendRoleInitParameters `json:"initProvider,omitempty"`
}

// SecretBackendRoleStatus defines the observed state of SecretBackendRole.
type SecretBackendRoleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecretBackendRoleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SecretBackendRole is the Schema for the SecretBackendRoles API. Manages a Consul secrets role for a Consul secrets engine in Vault.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type SecretBackendRole struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   SecretBackendRoleSpec   `json:"spec"`
	Status SecretBackendRoleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecretBackendRoleList contains a list of SecretBackendRoles
type SecretBackendRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecretBackendRole `json:"items"`
}

// Repository type metadata.
var (
	SecretBackendRole_Kind             = "SecretBackendRole"
	SecretBackendRole_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecretBackendRole_Kind}.String()
	SecretBackendRole_KindAPIVersion   = SecretBackendRole_Kind + "." + CRDGroupVersion.String()
	SecretBackendRole_GroupVersionKind = CRDGroupVersion.WithKind(SecretBackendRole_Kind)
)

func init() {
	SchemeBuilder.Register(&SecretBackendRole{}, &SecretBackendRoleList{})
}
