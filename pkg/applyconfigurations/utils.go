/*
Copyright The cert-manager Authors.

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
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package applyconfigurations

import (
	v1alpha1 "github.com/cert-manager/trust-manager/pkg/apis/trust/v1alpha1"
	v1alpha2 "github.com/cert-manager/trust-manager/pkg/apis/trustmanager/v1alpha2"
	internal "github.com/cert-manager/trust-manager/pkg/applyconfigurations/internal"
	trustv1alpha1 "github.com/cert-manager/trust-manager/pkg/applyconfigurations/trust/v1alpha1"
	trustmanagerv1alpha2 "github.com/cert-manager/trust-manager/pkg/applyconfigurations/trustmanager/v1alpha2"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	testing "k8s.io/client-go/testing"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) interface{} {
	switch kind {
	// Group=trust-manager.io, Version=v1alpha2
	case v1alpha2.SchemeGroupVersion.WithKind("AdditionalFormats"):
		return &trustmanagerv1alpha2.AdditionalFormatsApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("BundleSource"):
		return &trustmanagerv1alpha2.BundleSourceApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("BundleSpec"):
		return &trustmanagerv1alpha2.BundleSpecApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("BundleStatus"):
		return &trustmanagerv1alpha2.BundleStatusApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("BundleTarget"):
		return &trustmanagerv1alpha2.BundleTargetApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("ClusterBundle"):
		return &trustmanagerv1alpha2.ClusterBundleApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("JKS"):
		return &trustmanagerv1alpha2.JKSApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("KeySelector"):
		return &trustmanagerv1alpha2.KeySelectorApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("PKCS12"):
		return &trustmanagerv1alpha2.PKCS12ApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("SourceObjectKeySelector"):
		return &trustmanagerv1alpha2.SourceObjectKeySelectorApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("TargetMetadata"):
		return &trustmanagerv1alpha2.TargetMetadataApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("TargetTemplate"):
		return &trustmanagerv1alpha2.TargetTemplateApplyConfiguration{}

		// Group=trust.cert-manager.io, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithKind("AdditionalFormats"):
		return &trustv1alpha1.AdditionalFormatsApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("Bundle"):
		return &trustv1alpha1.BundleApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("BundleSource"):
		return &trustv1alpha1.BundleSourceApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("BundleSpec"):
		return &trustv1alpha1.BundleSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("BundleStatus"):
		return &trustv1alpha1.BundleStatusApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("BundleTarget"):
		return &trustv1alpha1.BundleTargetApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("JKS"):
		return &trustv1alpha1.JKSApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("KeySelector"):
		return &trustv1alpha1.KeySelectorApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("PKCS12"):
		return &trustv1alpha1.PKCS12ApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("SourceObjectKeySelector"):
		return &trustv1alpha1.SourceObjectKeySelectorApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("TargetMetadata"):
		return &trustv1alpha1.TargetMetadataApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("TargetTemplate"):
		return &trustv1alpha1.TargetTemplateApplyConfiguration{}

	}
	return nil
}

func NewTypeConverter(scheme *runtime.Scheme) *testing.TypeConverter {
	return &testing.TypeConverter{Scheme: scheme, TypeResolver: internal.Parser()}
}
