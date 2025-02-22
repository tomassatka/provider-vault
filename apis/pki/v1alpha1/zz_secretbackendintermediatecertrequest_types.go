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

type SecretBackendIntermediateCertRequestInitParameters struct {

	// Adds a Basic Constraints extension with 'CA: true'.
	// Only needed as a workaround in some compatibility scenarios with Active Directory
	// Certificate Services
	// Set 'CA: true' in a Basic Constraints extension. Only needed as
	// a workaround in some compatibility scenarios with Active Directory Certificate Services.
	AddBasicConstraints *bool `json:"addBasicConstraints,omitempty" tf:"add_basic_constraints,omitempty"`

	// List of alternative names
	// List of alternative names.
	AltNames []*string `json:"altNames,omitempty" tf:"alt_names,omitempty"`

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

	// CN of intermediate to create
	// CN of intermediate to create.
	CommonName *string `json:"commonName,omitempty" tf:"common_name,omitempty"`

	// The country
	// The country.
	Country *string `json:"country,omitempty" tf:"country,omitempty"`

	// Flag to exclude CN from SANs
	// Flag to exclude CN from SANs.
	ExcludeCnFromSans *bool `json:"excludeCnFromSans,omitempty" tf:"exclude_cn_from_sans,omitempty"`

	// The format of data
	// The format of data.
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	// List of alternative IPs
	// List of alternative IPs.
	IPSans []*string `json:"ipSans,omitempty" tf:"ip_sans,omitempty"`

	// The number of bits to use
	// The number of bits to use.
	KeyBits *float64 `json:"keyBits,omitempty" tf:"key_bits,omitempty"`

	// When a new key is created with this request, optionally specifies
	// the name for this. The global ref default may not be used as a name.
	// When a new key is created with this request, optionally specifies the name for this.
	KeyName *string `json:"keyName,omitempty" tf:"key_name,omitempty"`

	// Specifies the key (either default, by name, or by identifier) to use
	// for generating this request. Only suitable for type=existing requests.
	// Specifies the key to use for generating this request.
	KeyRef *string `json:"keyRef,omitempty" tf:"key_ref,omitempty"`

	// The desired key type
	// The desired key type.
	KeyType *string `json:"keyType,omitempty" tf:"key_type,omitempty"`

	// The locality
	// The locality.
	Locality *string `json:"locality,omitempty" tf:"locality,omitempty"`

	// The ID of the previously configured managed key. This field is
	// required if type is kms and it conflicts with managed_key_name
	// The ID of the previously configured managed key.
	ManagedKeyID *string `json:"managedKeyId,omitempty" tf:"managed_key_id,omitempty"`

	// The name of the previously configured managed key. This field is
	// required if type is kms  and it conflicts with managed_key_id
	// The name of the previously configured managed key.
	ManagedKeyName *string `json:"managedKeyName,omitempty" tf:"managed_key_name,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// The organization
	// The organization.
	Organization *string `json:"organization,omitempty" tf:"organization,omitempty"`

	// List of other SANs
	// List of other SANs.
	OtherSans []*string `json:"otherSans,omitempty" tf:"other_sans,omitempty"`

	// The organization unit
	// The organization unit.
	Ou *string `json:"ou,omitempty" tf:"ou,omitempty"`

	// The postal code
	// The postal code.
	PostalCode *string `json:"postalCode,omitempty" tf:"postal_code,omitempty"`

	// The private key format
	// The private key format.
	PrivateKeyFormat *string `json:"privateKeyFormat,omitempty" tf:"private_key_format,omitempty"`

	// The province
	// The province.
	Province *string `json:"province,omitempty" tf:"province,omitempty"`

	// The street address
	// The street address.
	StreetAddress *string `json:"streetAddress,omitempty" tf:"street_address,omitempty"`

	// Type of intermediate to create. Must be either "exported" or "internal"
	// or "kms"
	// Type of intermediate to create. Must be either "existing", "exported", "internal" or "kms"
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// List of alternative URIs
	// List of alternative URIs.
	URISans []*string `json:"uriSans,omitempty" tf:"uri_sans,omitempty"`
}

type SecretBackendIntermediateCertRequestObservation struct {

	// Adds a Basic Constraints extension with 'CA: true'.
	// Only needed as a workaround in some compatibility scenarios with Active Directory
	// Certificate Services
	// Set 'CA: true' in a Basic Constraints extension. Only needed as
	// a workaround in some compatibility scenarios with Active Directory Certificate Services.
	AddBasicConstraints *bool `json:"addBasicConstraints,omitempty" tf:"add_basic_constraints,omitempty"`

	// List of alternative names
	// List of alternative names.
	AltNames []*string `json:"altNames,omitempty" tf:"alt_names,omitempty"`

	// The PKI secret backend the resource belongs to.
	// The PKI secret backend the resource belongs to.
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// CN of intermediate to create
	// CN of intermediate to create.
	CommonName *string `json:"commonName,omitempty" tf:"common_name,omitempty"`

	// The country
	// The country.
	Country *string `json:"country,omitempty" tf:"country,omitempty"`

	// The CSR
	// The CSR.
	Csr *string `json:"csr,omitempty" tf:"csr,omitempty"`

	// Flag to exclude CN from SANs
	// Flag to exclude CN from SANs.
	ExcludeCnFromSans *bool `json:"excludeCnFromSans,omitempty" tf:"exclude_cn_from_sans,omitempty"`

	// The format of data
	// The format of data.
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// List of alternative IPs
	// List of alternative IPs.
	IPSans []*string `json:"ipSans,omitempty" tf:"ip_sans,omitempty"`

	// The number of bits to use
	// The number of bits to use.
	KeyBits *float64 `json:"keyBits,omitempty" tf:"key_bits,omitempty"`

	// The ID of the generated key.
	// The ID of the generated key.
	KeyID *string `json:"keyId,omitempty" tf:"key_id,omitempty"`

	// When a new key is created with this request, optionally specifies
	// the name for this. The global ref default may not be used as a name.
	// When a new key is created with this request, optionally specifies the name for this.
	KeyName *string `json:"keyName,omitempty" tf:"key_name,omitempty"`

	// Specifies the key (either default, by name, or by identifier) to use
	// for generating this request. Only suitable for type=existing requests.
	// Specifies the key to use for generating this request.
	KeyRef *string `json:"keyRef,omitempty" tf:"key_ref,omitempty"`

	// The desired key type
	// The desired key type.
	KeyType *string `json:"keyType,omitempty" tf:"key_type,omitempty"`

	// The locality
	// The locality.
	Locality *string `json:"locality,omitempty" tf:"locality,omitempty"`

	// The ID of the previously configured managed key. This field is
	// required if type is kms and it conflicts with managed_key_name
	// The ID of the previously configured managed key.
	ManagedKeyID *string `json:"managedKeyId,omitempty" tf:"managed_key_id,omitempty"`

	// The name of the previously configured managed key. This field is
	// required if type is kms  and it conflicts with managed_key_id
	// The name of the previously configured managed key.
	ManagedKeyName *string `json:"managedKeyName,omitempty" tf:"managed_key_name,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// The organization
	// The organization.
	Organization *string `json:"organization,omitempty" tf:"organization,omitempty"`

	// List of other SANs
	// List of other SANs.
	OtherSans []*string `json:"otherSans,omitempty" tf:"other_sans,omitempty"`

	// The organization unit
	// The organization unit.
	Ou *string `json:"ou,omitempty" tf:"ou,omitempty"`

	// The postal code
	// The postal code.
	PostalCode *string `json:"postalCode,omitempty" tf:"postal_code,omitempty"`

	// The private key format
	// The private key format.
	PrivateKeyFormat *string `json:"privateKeyFormat,omitempty" tf:"private_key_format,omitempty"`

	// The private key type
	// The private key type.
	PrivateKeyType *string `json:"privateKeyType,omitempty" tf:"private_key_type,omitempty"`

	// The province
	// The province.
	Province *string `json:"province,omitempty" tf:"province,omitempty"`

	// The street address
	// The street address.
	StreetAddress *string `json:"streetAddress,omitempty" tf:"street_address,omitempty"`

	// Type of intermediate to create. Must be either "exported" or "internal"
	// or "kms"
	// Type of intermediate to create. Must be either "existing", "exported", "internal" or "kms"
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// List of alternative URIs
	// List of alternative URIs.
	URISans []*string `json:"uriSans,omitempty" tf:"uri_sans,omitempty"`
}

type SecretBackendIntermediateCertRequestParameters struct {

	// Adds a Basic Constraints extension with 'CA: true'.
	// Only needed as a workaround in some compatibility scenarios with Active Directory
	// Certificate Services
	// Set 'CA: true' in a Basic Constraints extension. Only needed as
	// a workaround in some compatibility scenarios with Active Directory Certificate Services.
	// +kubebuilder:validation:Optional
	AddBasicConstraints *bool `json:"addBasicConstraints,omitempty" tf:"add_basic_constraints,omitempty"`

	// List of alternative names
	// List of alternative names.
	// +kubebuilder:validation:Optional
	AltNames []*string `json:"altNames,omitempty" tf:"alt_names,omitempty"`

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

	// CN of intermediate to create
	// CN of intermediate to create.
	// +kubebuilder:validation:Optional
	CommonName *string `json:"commonName,omitempty" tf:"common_name,omitempty"`

	// The country
	// The country.
	// +kubebuilder:validation:Optional
	Country *string `json:"country,omitempty" tf:"country,omitempty"`

	// Flag to exclude CN from SANs
	// Flag to exclude CN from SANs.
	// +kubebuilder:validation:Optional
	ExcludeCnFromSans *bool `json:"excludeCnFromSans,omitempty" tf:"exclude_cn_from_sans,omitempty"`

	// The format of data
	// The format of data.
	// +kubebuilder:validation:Optional
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	// List of alternative IPs
	// List of alternative IPs.
	// +kubebuilder:validation:Optional
	IPSans []*string `json:"ipSans,omitempty" tf:"ip_sans,omitempty"`

	// The number of bits to use
	// The number of bits to use.
	// +kubebuilder:validation:Optional
	KeyBits *float64 `json:"keyBits,omitempty" tf:"key_bits,omitempty"`

	// When a new key is created with this request, optionally specifies
	// the name for this. The global ref default may not be used as a name.
	// When a new key is created with this request, optionally specifies the name for this.
	// +kubebuilder:validation:Optional
	KeyName *string `json:"keyName,omitempty" tf:"key_name,omitempty"`

	// Specifies the key (either default, by name, or by identifier) to use
	// for generating this request. Only suitable for type=existing requests.
	// Specifies the key to use for generating this request.
	// +kubebuilder:validation:Optional
	KeyRef *string `json:"keyRef,omitempty" tf:"key_ref,omitempty"`

	// The desired key type
	// The desired key type.
	// +kubebuilder:validation:Optional
	KeyType *string `json:"keyType,omitempty" tf:"key_type,omitempty"`

	// The locality
	// The locality.
	// +kubebuilder:validation:Optional
	Locality *string `json:"locality,omitempty" tf:"locality,omitempty"`

	// The ID of the previously configured managed key. This field is
	// required if type is kms and it conflicts with managed_key_name
	// The ID of the previously configured managed key.
	// +kubebuilder:validation:Optional
	ManagedKeyID *string `json:"managedKeyId,omitempty" tf:"managed_key_id,omitempty"`

	// The name of the previously configured managed key. This field is
	// required if type is kms  and it conflicts with managed_key_id
	// The name of the previously configured managed key.
	// +kubebuilder:validation:Optional
	ManagedKeyName *string `json:"managedKeyName,omitempty" tf:"managed_key_name,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// The organization
	// The organization.
	// +kubebuilder:validation:Optional
	Organization *string `json:"organization,omitempty" tf:"organization,omitempty"`

	// List of other SANs
	// List of other SANs.
	// +kubebuilder:validation:Optional
	OtherSans []*string `json:"otherSans,omitempty" tf:"other_sans,omitempty"`

	// The organization unit
	// The organization unit.
	// +kubebuilder:validation:Optional
	Ou *string `json:"ou,omitempty" tf:"ou,omitempty"`

	// The postal code
	// The postal code.
	// +kubebuilder:validation:Optional
	PostalCode *string `json:"postalCode,omitempty" tf:"postal_code,omitempty"`

	// The private key format
	// The private key format.
	// +kubebuilder:validation:Optional
	PrivateKeyFormat *string `json:"privateKeyFormat,omitempty" tf:"private_key_format,omitempty"`

	// The province
	// The province.
	// +kubebuilder:validation:Optional
	Province *string `json:"province,omitempty" tf:"province,omitempty"`

	// The street address
	// The street address.
	// +kubebuilder:validation:Optional
	StreetAddress *string `json:"streetAddress,omitempty" tf:"street_address,omitempty"`

	// Type of intermediate to create. Must be either "exported" or "internal"
	// or "kms"
	// Type of intermediate to create. Must be either "existing", "exported", "internal" or "kms"
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// List of alternative URIs
	// List of alternative URIs.
	// +kubebuilder:validation:Optional
	URISans []*string `json:"uriSans,omitempty" tf:"uri_sans,omitempty"`
}

// SecretBackendIntermediateCertRequestSpec defines the desired state of SecretBackendIntermediateCertRequest
type SecretBackendIntermediateCertRequestSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecretBackendIntermediateCertRequestParameters `json:"forProvider"`
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
	InitProvider SecretBackendIntermediateCertRequestInitParameters `json:"initProvider,omitempty"`
}

// SecretBackendIntermediateCertRequestStatus defines the observed state of SecretBackendIntermediateCertRequest.
type SecretBackendIntermediateCertRequestStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecretBackendIntermediateCertRequestObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SecretBackendIntermediateCertRequest is the Schema for the SecretBackendIntermediateCertRequests API. Generate a new private key and a CSR for signing the PKI.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type SecretBackendIntermediateCertRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.commonName) || (has(self.initProvider) && has(self.initProvider.commonName))",message="spec.forProvider.commonName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   SecretBackendIntermediateCertRequestSpec   `json:"spec"`
	Status SecretBackendIntermediateCertRequestStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecretBackendIntermediateCertRequestList contains a list of SecretBackendIntermediateCertRequests
type SecretBackendIntermediateCertRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecretBackendIntermediateCertRequest `json:"items"`
}

// Repository type metadata.
var (
	SecretBackendIntermediateCertRequest_Kind             = "SecretBackendIntermediateCertRequest"
	SecretBackendIntermediateCertRequest_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecretBackendIntermediateCertRequest_Kind}.String()
	SecretBackendIntermediateCertRequest_KindAPIVersion   = SecretBackendIntermediateCertRequest_Kind + "." + CRDGroupVersion.String()
	SecretBackendIntermediateCertRequest_GroupVersionKind = CRDGroupVersion.WithKind(SecretBackendIntermediateCertRequest_Kind)
)

func init() {
	SchemeBuilder.Register(&SecretBackendIntermediateCertRequest{}, &SecretBackendIntermediateCertRequestList{})
}
