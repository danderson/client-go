// +build !ignore_autogenerated

/*
Copyright 2018 The Kubernetes Authors.

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

// This file was autogenerated by conversion-gen. Do not edit it manually!

package extensionsv1beta1

import (
	v1beta1 "k8s.io/api/extensions/v1beta1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	scheme "k8s.io/client-go/v6/scale/scheme"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1beta1_Scale_To_scheme_Scale,
		Convert_scheme_Scale_To_v1beta1_Scale,
		Convert_v1beta1_ScaleSpec_To_scheme_ScaleSpec,
		Convert_scheme_ScaleSpec_To_v1beta1_ScaleSpec,
		Convert_v1beta1_ScaleStatus_To_scheme_ScaleStatus,
		Convert_scheme_ScaleStatus_To_v1beta1_ScaleStatus,
	)
}

func autoConvert_v1beta1_Scale_To_scheme_Scale(in *v1beta1.Scale, out *scheme.Scale, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_ScaleSpec_To_scheme_ScaleSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_ScaleStatus_To_scheme_ScaleStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_Scale_To_scheme_Scale is an autogenerated conversion function.
func Convert_v1beta1_Scale_To_scheme_Scale(in *v1beta1.Scale, out *scheme.Scale, s conversion.Scope) error {
	return autoConvert_v1beta1_Scale_To_scheme_Scale(in, out, s)
}

func autoConvert_scheme_Scale_To_v1beta1_Scale(in *scheme.Scale, out *v1beta1.Scale, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_scheme_ScaleSpec_To_v1beta1_ScaleSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_scheme_ScaleStatus_To_v1beta1_ScaleStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_scheme_Scale_To_v1beta1_Scale is an autogenerated conversion function.
func Convert_scheme_Scale_To_v1beta1_Scale(in *scheme.Scale, out *v1beta1.Scale, s conversion.Scope) error {
	return autoConvert_scheme_Scale_To_v1beta1_Scale(in, out, s)
}

func autoConvert_v1beta1_ScaleSpec_To_scheme_ScaleSpec(in *v1beta1.ScaleSpec, out *scheme.ScaleSpec, s conversion.Scope) error {
	out.Replicas = in.Replicas
	return nil
}

// Convert_v1beta1_ScaleSpec_To_scheme_ScaleSpec is an autogenerated conversion function.
func Convert_v1beta1_ScaleSpec_To_scheme_ScaleSpec(in *v1beta1.ScaleSpec, out *scheme.ScaleSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_ScaleSpec_To_scheme_ScaleSpec(in, out, s)
}

func autoConvert_scheme_ScaleSpec_To_v1beta1_ScaleSpec(in *scheme.ScaleSpec, out *v1beta1.ScaleSpec, s conversion.Scope) error {
	out.Replicas = in.Replicas
	return nil
}

// Convert_scheme_ScaleSpec_To_v1beta1_ScaleSpec is an autogenerated conversion function.
func Convert_scheme_ScaleSpec_To_v1beta1_ScaleSpec(in *scheme.ScaleSpec, out *v1beta1.ScaleSpec, s conversion.Scope) error {
	return autoConvert_scheme_ScaleSpec_To_v1beta1_ScaleSpec(in, out, s)
}

func autoConvert_v1beta1_ScaleStatus_To_scheme_ScaleStatus(in *v1beta1.ScaleStatus, out *scheme.ScaleStatus, s conversion.Scope) error {
	out.Replicas = in.Replicas
	// WARNING: in.Selector requires manual conversion: inconvertible types (map[string]string vs *k8s.io/apimachinery/pkg/apis/meta/v1.LabelSelector)
	// WARNING: in.TargetSelector requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_scheme_ScaleStatus_To_v1beta1_ScaleStatus(in *scheme.ScaleStatus, out *v1beta1.ScaleStatus, s conversion.Scope) error {
	out.Replicas = in.Replicas
	// WARNING: in.Selector requires manual conversion: inconvertible types (*k8s.io/apimachinery/pkg/apis/meta/v1.LabelSelector vs map[string]string)
	return nil
}
