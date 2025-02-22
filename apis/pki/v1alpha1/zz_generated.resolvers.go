/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"
	v1alpha1 "github.com/upbound/provider-vault/v2/apis/vault/v1alpha1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this SecretBackendCert.
func (mg *SecretBackendCert) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Backend),
		Extract:      resource.ExtractParamPath("path", false),
		Reference:    mg.Spec.ForProvider.BackendRef,
		Selector:     mg.Spec.ForProvider.BackendSelector,
		To: reference.To{
			List:    &v1alpha1.MountList{},
			Managed: &v1alpha1.Mount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Backend")
	}
	mg.Spec.ForProvider.Backend = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BackendRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Name),
		Extract:      resource.ExtractParamPath("name", false),
		Reference:    mg.Spec.ForProvider.NameRef,
		Selector:     mg.Spec.ForProvider.NameSelector,
		To: reference.To{
			List:    &SecretBackendRoleList{},
			Managed: &SecretBackendRole{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Name")
	}
	mg.Spec.ForProvider.Name = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Backend),
		Extract:      resource.ExtractParamPath("path", false),
		Reference:    mg.Spec.InitProvider.BackendRef,
		Selector:     mg.Spec.InitProvider.BackendSelector,
		To: reference.To{
			List:    &v1alpha1.MountList{},
			Managed: &v1alpha1.Mount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Backend")
	}
	mg.Spec.InitProvider.Backend = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.BackendRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Name),
		Extract:      resource.ExtractParamPath("name", false),
		Reference:    mg.Spec.InitProvider.NameRef,
		Selector:     mg.Spec.InitProvider.NameSelector,
		To: reference.To{
			List:    &SecretBackendRoleList{},
			Managed: &SecretBackendRole{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Name")
	}
	mg.Spec.InitProvider.Name = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.NameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SecretBackendConfigCA.
func (mg *SecretBackendConfigCA) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Backend),
		Extract:      resource.ExtractParamPath("path", false),
		Reference:    mg.Spec.ForProvider.BackendRef,
		Selector:     mg.Spec.ForProvider.BackendSelector,
		To: reference.To{
			List:    &v1alpha1.MountList{},
			Managed: &v1alpha1.Mount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Backend")
	}
	mg.Spec.ForProvider.Backend = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BackendRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Backend),
		Extract:      resource.ExtractParamPath("path", false),
		Reference:    mg.Spec.InitProvider.BackendRef,
		Selector:     mg.Spec.InitProvider.BackendSelector,
		To: reference.To{
			List:    &v1alpha1.MountList{},
			Managed: &v1alpha1.Mount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Backend")
	}
	mg.Spec.InitProvider.Backend = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.BackendRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SecretBackendConfigUrls.
func (mg *SecretBackendConfigUrls) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Backend),
		Extract:      resource.ExtractParamPath("path", false),
		Reference:    mg.Spec.ForProvider.BackendRef,
		Selector:     mg.Spec.ForProvider.BackendSelector,
		To: reference.To{
			List:    &v1alpha1.MountList{},
			Managed: &v1alpha1.Mount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Backend")
	}
	mg.Spec.ForProvider.Backend = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BackendRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Backend),
		Extract:      resource.ExtractParamPath("path", false),
		Reference:    mg.Spec.InitProvider.BackendRef,
		Selector:     mg.Spec.InitProvider.BackendSelector,
		To: reference.To{
			List:    &v1alpha1.MountList{},
			Managed: &v1alpha1.Mount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Backend")
	}
	mg.Spec.InitProvider.Backend = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.BackendRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SecretBackendCrlConfig.
func (mg *SecretBackendCrlConfig) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Backend),
		Extract:      resource.ExtractParamPath("path", false),
		Reference:    mg.Spec.ForProvider.BackendRef,
		Selector:     mg.Spec.ForProvider.BackendSelector,
		To: reference.To{
			List:    &v1alpha1.MountList{},
			Managed: &v1alpha1.Mount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Backend")
	}
	mg.Spec.ForProvider.Backend = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BackendRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Backend),
		Extract:      resource.ExtractParamPath("path", false),
		Reference:    mg.Spec.InitProvider.BackendRef,
		Selector:     mg.Spec.InitProvider.BackendSelector,
		To: reference.To{
			List:    &v1alpha1.MountList{},
			Managed: &v1alpha1.Mount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Backend")
	}
	mg.Spec.InitProvider.Backend = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.BackendRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SecretBackendIntermediateCertRequest.
func (mg *SecretBackendIntermediateCertRequest) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Backend),
		Extract:      resource.ExtractParamPath("path", false),
		Reference:    mg.Spec.ForProvider.BackendRef,
		Selector:     mg.Spec.ForProvider.BackendSelector,
		To: reference.To{
			List:    &v1alpha1.MountList{},
			Managed: &v1alpha1.Mount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Backend")
	}
	mg.Spec.ForProvider.Backend = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BackendRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Backend),
		Extract:      resource.ExtractParamPath("path", false),
		Reference:    mg.Spec.InitProvider.BackendRef,
		Selector:     mg.Spec.InitProvider.BackendSelector,
		To: reference.To{
			List:    &v1alpha1.MountList{},
			Managed: &v1alpha1.Mount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Backend")
	}
	mg.Spec.InitProvider.Backend = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.BackendRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SecretBackendIntermediateSetSigned.
func (mg *SecretBackendIntermediateSetSigned) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Backend),
		Extract:      resource.ExtractParamPath("path", false),
		Reference:    mg.Spec.ForProvider.BackendRef,
		Selector:     mg.Spec.ForProvider.BackendSelector,
		To: reference.To{
			List:    &v1alpha1.MountList{},
			Managed: &v1alpha1.Mount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Backend")
	}
	mg.Spec.ForProvider.Backend = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BackendRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Certificate),
		Extract:      resource.ExtractParamPath("certificate", true),
		Reference:    mg.Spec.ForProvider.CertificateRef,
		Selector:     mg.Spec.ForProvider.CertificateSelector,
		To: reference.To{
			List:    &SecretBackendRootSignIntermediateList{},
			Managed: &SecretBackendRootSignIntermediate{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Certificate")
	}
	mg.Spec.ForProvider.Certificate = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CertificateRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Backend),
		Extract:      resource.ExtractParamPath("path", false),
		Reference:    mg.Spec.InitProvider.BackendRef,
		Selector:     mg.Spec.InitProvider.BackendSelector,
		To: reference.To{
			List:    &v1alpha1.MountList{},
			Managed: &v1alpha1.Mount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Backend")
	}
	mg.Spec.InitProvider.Backend = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.BackendRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Certificate),
		Extract:      resource.ExtractParamPath("certificate", true),
		Reference:    mg.Spec.InitProvider.CertificateRef,
		Selector:     mg.Spec.InitProvider.CertificateSelector,
		To: reference.To{
			List:    &SecretBackendRootSignIntermediateList{},
			Managed: &SecretBackendRootSignIntermediate{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Certificate")
	}
	mg.Spec.InitProvider.Certificate = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.CertificateRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SecretBackendRole.
func (mg *SecretBackendRole) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Backend),
		Extract:      resource.ExtractParamPath("path", false),
		Reference:    mg.Spec.ForProvider.BackendRef,
		Selector:     mg.Spec.ForProvider.BackendSelector,
		To: reference.To{
			List:    &v1alpha1.MountList{},
			Managed: &v1alpha1.Mount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Backend")
	}
	mg.Spec.ForProvider.Backend = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BackendRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Backend),
		Extract:      resource.ExtractParamPath("path", false),
		Reference:    mg.Spec.InitProvider.BackendRef,
		Selector:     mg.Spec.InitProvider.BackendSelector,
		To: reference.To{
			List:    &v1alpha1.MountList{},
			Managed: &v1alpha1.Mount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Backend")
	}
	mg.Spec.InitProvider.Backend = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.BackendRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SecretBackendRootCert.
func (mg *SecretBackendRootCert) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Backend),
		Extract:      resource.ExtractParamPath("path", false),
		Reference:    mg.Spec.ForProvider.BackendRef,
		Selector:     mg.Spec.ForProvider.BackendSelector,
		To: reference.To{
			List:    &v1alpha1.MountList{},
			Managed: &v1alpha1.Mount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Backend")
	}
	mg.Spec.ForProvider.Backend = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BackendRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Backend),
		Extract:      resource.ExtractParamPath("path", false),
		Reference:    mg.Spec.InitProvider.BackendRef,
		Selector:     mg.Spec.InitProvider.BackendSelector,
		To: reference.To{
			List:    &v1alpha1.MountList{},
			Managed: &v1alpha1.Mount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Backend")
	}
	mg.Spec.InitProvider.Backend = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.BackendRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SecretBackendSign.
func (mg *SecretBackendSign) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Backend),
		Extract:      resource.ExtractParamPath("path", false),
		Reference:    mg.Spec.ForProvider.BackendRef,
		Selector:     mg.Spec.ForProvider.BackendSelector,
		To: reference.To{
			List:    &v1alpha1.MountList{},
			Managed: &v1alpha1.Mount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Backend")
	}
	mg.Spec.ForProvider.Backend = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BackendRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Name),
		Extract:      resource.ExtractParamPath("name", false),
		Reference:    mg.Spec.ForProvider.NameRef,
		Selector:     mg.Spec.ForProvider.NameSelector,
		To: reference.To{
			List:    &SecretBackendRoleList{},
			Managed: &SecretBackendRole{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Name")
	}
	mg.Spec.ForProvider.Name = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Backend),
		Extract:      resource.ExtractParamPath("path", false),
		Reference:    mg.Spec.InitProvider.BackendRef,
		Selector:     mg.Spec.InitProvider.BackendSelector,
		To: reference.To{
			List:    &v1alpha1.MountList{},
			Managed: &v1alpha1.Mount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Backend")
	}
	mg.Spec.InitProvider.Backend = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.BackendRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Name),
		Extract:      resource.ExtractParamPath("name", false),
		Reference:    mg.Spec.InitProvider.NameRef,
		Selector:     mg.Spec.InitProvider.NameSelector,
		To: reference.To{
			List:    &SecretBackendRoleList{},
			Managed: &SecretBackendRole{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Name")
	}
	mg.Spec.InitProvider.Name = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.NameRef = rsp.ResolvedReference

	return nil
}
