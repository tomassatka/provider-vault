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

type GroupInitParameters struct {

	// false by default. If set to true, this resource will ignore any Entity IDs
	// returned from Vault or specified in the resource. You can use
	// vault_identity_group_member_entity_ids to manage Entity IDs for this group in a
	// decoupled manner.
	// Manage member entities externally through `vault_identity_group_member_entity_ids`
	ExternalMemberEntityIds *bool `json:"externalMemberEntityIds,omitempty" tf:"external_member_entity_ids,omitempty"`

	// false by default. If set to true, this resource will ignore any Group IDs
	// returned from Vault or specified in the resource. You can use
	// vault_identity_group_member_group_ids to manage Group IDs for this group in a
	// decoupled manner.
	// Manage member groups externally through `vault_identity_group_member_group_ids`
	ExternalMemberGroupIds *bool `json:"externalMemberGroupIds,omitempty" tf:"external_member_group_ids,omitempty"`

	// false by default. If set to true, this resource will ignore any policies returned from
	// Vault or specified in the resource. You can use vault_identity_group_policies to manage
	// policies for this group in a decoupled manner.
	// Manage policies externally through `vault_identity_group_policies`, allows using group ID in assigned policies.
	ExternalPolicies *bool `json:"externalPolicies,omitempty" tf:"external_policies,omitempty"`

	// A list of Entity IDs to be assigned as group members. Not allowed on external groups.
	// Entity IDs to be assigned as group members.
	// +listType=set
	MemberEntityIds []*string `json:"memberEntityIds,omitempty" tf:"member_entity_ids,omitempty"`

	// A list of Group IDs to be assigned as group members. Not allowed on external groups.
	// Group IDs to be assigned as group members.
	// +listType=set
	MemberGroupIds []*string `json:"memberGroupIds,omitempty" tf:"member_group_ids,omitempty"`

	// A Map of additional metadata to associate with the group.
	// Metadata to be associated with the group.
	// +mapType=granular
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// Name of the identity group to create.
	// Name of the group.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// A list of policies to apply to the group.
	// Policies to be tied to the group.
	// +listType=set
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`

	// Type of the group, internal or external. Defaults to internal.
	// Type of the group, internal or external. Defaults to internal.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type GroupObservation struct {

	// false by default. If set to true, this resource will ignore any Entity IDs
	// returned from Vault or specified in the resource. You can use
	// vault_identity_group_member_entity_ids to manage Entity IDs for this group in a
	// decoupled manner.
	// Manage member entities externally through `vault_identity_group_member_entity_ids`
	ExternalMemberEntityIds *bool `json:"externalMemberEntityIds,omitempty" tf:"external_member_entity_ids,omitempty"`

	// false by default. If set to true, this resource will ignore any Group IDs
	// returned from Vault or specified in the resource. You can use
	// vault_identity_group_member_group_ids to manage Group IDs for this group in a
	// decoupled manner.
	// Manage member groups externally through `vault_identity_group_member_group_ids`
	ExternalMemberGroupIds *bool `json:"externalMemberGroupIds,omitempty" tf:"external_member_group_ids,omitempty"`

	// false by default. If set to true, this resource will ignore any policies returned from
	// Vault or specified in the resource. You can use vault_identity_group_policies to manage
	// policies for this group in a decoupled manner.
	// Manage policies externally through `vault_identity_group_policies`, allows using group ID in assigned policies.
	ExternalPolicies *bool `json:"externalPolicies,omitempty" tf:"external_policies,omitempty"`

	// The id of the created group.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A list of Entity IDs to be assigned as group members. Not allowed on external groups.
	// Entity IDs to be assigned as group members.
	// +listType=set
	MemberEntityIds []*string `json:"memberEntityIds,omitempty" tf:"member_entity_ids,omitempty"`

	// A list of Group IDs to be assigned as group members. Not allowed on external groups.
	// Group IDs to be assigned as group members.
	// +listType=set
	MemberGroupIds []*string `json:"memberGroupIds,omitempty" tf:"member_group_ids,omitempty"`

	// A Map of additional metadata to associate with the group.
	// Metadata to be associated with the group.
	// +mapType=granular
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// Name of the identity group to create.
	// Name of the group.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// A list of policies to apply to the group.
	// Policies to be tied to the group.
	// +listType=set
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`

	// Type of the group, internal or external. Defaults to internal.
	// Type of the group, internal or external. Defaults to internal.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type GroupParameters struct {

	// false by default. If set to true, this resource will ignore any Entity IDs
	// returned from Vault or specified in the resource. You can use
	// vault_identity_group_member_entity_ids to manage Entity IDs for this group in a
	// decoupled manner.
	// Manage member entities externally through `vault_identity_group_member_entity_ids`
	// +kubebuilder:validation:Optional
	ExternalMemberEntityIds *bool `json:"externalMemberEntityIds,omitempty" tf:"external_member_entity_ids,omitempty"`

	// false by default. If set to true, this resource will ignore any Group IDs
	// returned from Vault or specified in the resource. You can use
	// vault_identity_group_member_group_ids to manage Group IDs for this group in a
	// decoupled manner.
	// Manage member groups externally through `vault_identity_group_member_group_ids`
	// +kubebuilder:validation:Optional
	ExternalMemberGroupIds *bool `json:"externalMemberGroupIds,omitempty" tf:"external_member_group_ids,omitempty"`

	// false by default. If set to true, this resource will ignore any policies returned from
	// Vault or specified in the resource. You can use vault_identity_group_policies to manage
	// policies for this group in a decoupled manner.
	// Manage policies externally through `vault_identity_group_policies`, allows using group ID in assigned policies.
	// +kubebuilder:validation:Optional
	ExternalPolicies *bool `json:"externalPolicies,omitempty" tf:"external_policies,omitempty"`

	// A list of Entity IDs to be assigned as group members. Not allowed on external groups.
	// Entity IDs to be assigned as group members.
	// +kubebuilder:validation:Optional
	// +listType=set
	MemberEntityIds []*string `json:"memberEntityIds,omitempty" tf:"member_entity_ids,omitempty"`

	// A list of Group IDs to be assigned as group members. Not allowed on external groups.
	// Group IDs to be assigned as group members.
	// +kubebuilder:validation:Optional
	// +listType=set
	MemberGroupIds []*string `json:"memberGroupIds,omitempty" tf:"member_group_ids,omitempty"`

	// A Map of additional metadata to associate with the group.
	// Metadata to be associated with the group.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// Name of the identity group to create.
	// Name of the group.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// A list of policies to apply to the group.
	// Policies to be tied to the group.
	// +kubebuilder:validation:Optional
	// +listType=set
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`

	// Type of the group, internal or external. Defaults to internal.
	// Type of the group, internal or external. Defaults to internal.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// GroupSpec defines the desired state of Group
type GroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GroupParameters `json:"forProvider"`
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
	InitProvider GroupInitParameters `json:"initProvider,omitempty"`
}

// GroupStatus defines the observed state of Group.
type GroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Group is the Schema for the Groups API. Creates an Identity Group for Vault.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type Group struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GroupSpec   `json:"spec"`
	Status            GroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GroupList contains a list of Groups
type GroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Group `json:"items"`
}

// Repository type metadata.
var (
	Group_Kind             = "Group"
	Group_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Group_Kind}.String()
	Group_KindAPIVersion   = Group_Kind + "." + CRDGroupVersion.String()
	Group_GroupVersionKind = CRDGroupVersion.WithKind(Group_Kind)
)

func init() {
	SchemeBuilder.Register(&Group{}, &GroupList{})
}
