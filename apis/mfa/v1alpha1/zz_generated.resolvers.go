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
	v1alpha1 "github.com/upbound/provider-vault/v2/apis/auth/v1alpha1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Duo.
func (mg *Duo) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.MountAccessor),
		Extract:      resource.ExtractParamPath("accessor", true),
		Reference:    mg.Spec.ForProvider.MountAccessorRef,
		Selector:     mg.Spec.ForProvider.MountAccessorSelector,
		To: reference.To{
			List:    &v1alpha1.BackendList{},
			Managed: &v1alpha1.Backend{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.MountAccessor")
	}
	mg.Spec.ForProvider.MountAccessor = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.MountAccessorRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.MountAccessor),
		Extract:      resource.ExtractParamPath("accessor", true),
		Reference:    mg.Spec.InitProvider.MountAccessorRef,
		Selector:     mg.Spec.InitProvider.MountAccessorSelector,
		To: reference.To{
			List:    &v1alpha1.BackendList{},
			Managed: &v1alpha1.Backend{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.MountAccessor")
	}
	mg.Spec.InitProvider.MountAccessor = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.MountAccessorRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Okta.
func (mg *Okta) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.MountAccessor),
		Extract:      resource.ExtractParamPath("accessor", true),
		Reference:    mg.Spec.ForProvider.MountAccessorRef,
		Selector:     mg.Spec.ForProvider.MountAccessorSelector,
		To: reference.To{
			List:    &v1alpha1.BackendList{},
			Managed: &v1alpha1.Backend{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.MountAccessor")
	}
	mg.Spec.ForProvider.MountAccessor = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.MountAccessorRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.MountAccessor),
		Extract:      resource.ExtractParamPath("accessor", true),
		Reference:    mg.Spec.InitProvider.MountAccessorRef,
		Selector:     mg.Spec.InitProvider.MountAccessorSelector,
		To: reference.To{
			List:    &v1alpha1.BackendList{},
			Managed: &v1alpha1.Backend{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.MountAccessor")
	}
	mg.Spec.InitProvider.MountAccessor = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.MountAccessorRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Pingid.
func (mg *Pingid) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.MountAccessor),
		Extract:      resource.ExtractParamPath("accessor", true),
		Reference:    mg.Spec.ForProvider.MountAccessorRef,
		Selector:     mg.Spec.ForProvider.MountAccessorSelector,
		To: reference.To{
			List:    &v1alpha1.BackendList{},
			Managed: &v1alpha1.Backend{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.MountAccessor")
	}
	mg.Spec.ForProvider.MountAccessor = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.MountAccessorRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.MountAccessor),
		Extract:      resource.ExtractParamPath("accessor", true),
		Reference:    mg.Spec.InitProvider.MountAccessorRef,
		Selector:     mg.Spec.InitProvider.MountAccessorSelector,
		To: reference.To{
			List:    &v1alpha1.BackendList{},
			Managed: &v1alpha1.Backend{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.MountAccessor")
	}
	mg.Spec.InitProvider.MountAccessor = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.MountAccessorRef = rsp.ResolvedReference

	return nil
}
