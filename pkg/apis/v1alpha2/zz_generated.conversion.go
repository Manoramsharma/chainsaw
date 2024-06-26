//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha2

import (
	unsafe "unsafe"

	v1alpha1 "github.com/kyverno/chainsaw/pkg/apis/v1alpha1"
	policyv1alpha1 "github.com/kyverno/kyverno-json/pkg/apis/policy/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*ActionCheck)(nil), (*v1alpha1.ActionCheck)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_ActionCheck_To_v1alpha1_ActionCheck(a.(*ActionCheck), b.(*v1alpha1.ActionCheck), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.ActionCheck)(nil), (*ActionCheck)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ActionCheck_To_v1alpha2_ActionCheck(a.(*v1alpha1.ActionCheck), b.(*ActionCheck), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ActionCheckRef)(nil), (*v1alpha1.ActionCheckRef)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_ActionCheckRef_To_v1alpha1_ActionCheckRef(a.(*ActionCheckRef), b.(*v1alpha1.ActionCheckRef), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.ActionCheckRef)(nil), (*ActionCheckRef)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ActionCheckRef_To_v1alpha2_ActionCheckRef(a.(*v1alpha1.ActionCheckRef), b.(*ActionCheckRef), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ActionDryRun)(nil), (*v1alpha1.ActionDryRun)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_ActionDryRun_To_v1alpha1_ActionDryRun(a.(*ActionDryRun), b.(*v1alpha1.ActionDryRun), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.ActionDryRun)(nil), (*ActionDryRun)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ActionDryRun_To_v1alpha2_ActionDryRun(a.(*v1alpha1.ActionDryRun), b.(*ActionDryRun), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ActionEnv)(nil), (*v1alpha1.ActionEnv)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_ActionEnv_To_v1alpha1_ActionEnv(a.(*ActionEnv), b.(*v1alpha1.ActionEnv), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.ActionEnv)(nil), (*ActionEnv)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ActionEnv_To_v1alpha2_ActionEnv(a.(*v1alpha1.ActionEnv), b.(*ActionEnv), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ActionExpectations)(nil), (*v1alpha1.ActionExpectations)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_ActionExpectations_To_v1alpha1_ActionExpectations(a.(*ActionExpectations), b.(*v1alpha1.ActionExpectations), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.ActionExpectations)(nil), (*ActionExpectations)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ActionExpectations_To_v1alpha2_ActionExpectations(a.(*v1alpha1.ActionExpectations), b.(*ActionExpectations), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ActionFormat)(nil), (*v1alpha1.ActionFormat)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_ActionFormat_To_v1alpha1_ActionFormat(a.(*ActionFormat), b.(*v1alpha1.ActionFormat), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.ActionFormat)(nil), (*ActionFormat)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ActionFormat_To_v1alpha2_ActionFormat(a.(*v1alpha1.ActionFormat), b.(*ActionFormat), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ActionInlineResource)(nil), (*v1alpha1.ActionInlineResource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_ActionInlineResource_To_v1alpha1_ActionInlineResource(a.(*ActionInlineResource), b.(*v1alpha1.ActionInlineResource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.ActionInlineResource)(nil), (*ActionInlineResource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ActionInlineResource_To_v1alpha2_ActionInlineResource(a.(*v1alpha1.ActionInlineResource), b.(*ActionInlineResource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ActionObject)(nil), (*v1alpha1.ActionObject)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_ActionObject_To_v1alpha1_ActionObject(a.(*ActionObject), b.(*v1alpha1.ActionObject), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.ActionObject)(nil), (*ActionObject)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ActionObject_To_v1alpha2_ActionObject(a.(*v1alpha1.ActionObject), b.(*ActionObject), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ActionObjectSelector)(nil), (*v1alpha1.ActionObjectSelector)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_ActionObjectSelector_To_v1alpha1_ActionObjectSelector(a.(*ActionObjectSelector), b.(*v1alpha1.ActionObjectSelector), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.ActionObjectSelector)(nil), (*ActionObjectSelector)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ActionObjectSelector_To_v1alpha2_ActionObjectSelector(a.(*v1alpha1.ActionObjectSelector), b.(*ActionObjectSelector), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ActionResourceRef)(nil), (*v1alpha1.ActionResourceRef)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_ActionResourceRef_To_v1alpha1_ActionResourceRef(a.(*ActionResourceRef), b.(*v1alpha1.ActionResourceRef), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.ActionResourceRef)(nil), (*ActionResourceRef)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ActionResourceRef_To_v1alpha2_ActionResourceRef(a.(*v1alpha1.ActionResourceRef), b.(*ActionResourceRef), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ActionTimeout)(nil), (*v1alpha1.ActionTimeout)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_ActionTimeout_To_v1alpha1_ActionTimeout(a.(*ActionTimeout), b.(*v1alpha1.ActionTimeout), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.ActionTimeout)(nil), (*ActionTimeout)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ActionTimeout_To_v1alpha2_ActionTimeout(a.(*v1alpha1.ActionTimeout), b.(*ActionTimeout), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Configuration)(nil), (*v1alpha1.Configuration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_Configuration_To_v1alpha1_Configuration(a.(*Configuration), b.(*v1alpha1.Configuration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.Configuration)(nil), (*Configuration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Configuration_To_v1alpha2_Configuration(a.(*v1alpha1.Configuration), b.(*Configuration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*FileRef)(nil), (*v1alpha1.FileRef)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_FileRef_To_v1alpha1_FileRef(a.(*FileRef), b.(*v1alpha1.FileRef), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.FileRef)(nil), (*FileRef)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_FileRef_To_v1alpha2_FileRef(a.(*v1alpha1.FileRef), b.(*FileRef), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Test)(nil), (*v1alpha1.Test)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_Test_To_v1alpha1_Test(a.(*Test), b.(*v1alpha1.Test), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.Test)(nil), (*Test)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Test_To_v1alpha2_Test(a.(*v1alpha1.Test), b.(*Test), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1alpha1.ConfigurationSpec)(nil), (*ConfigurationSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ConfigurationSpec_To_v1alpha2_ConfigurationSpec(a.(*v1alpha1.ConfigurationSpec), b.(*ConfigurationSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1alpha1.TestSpec)(nil), (*TestSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_TestSpec_To_v1alpha2_TestSpec(a.(*v1alpha1.TestSpec), b.(*TestSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*ConfigurationSpec)(nil), (*v1alpha1.ConfigurationSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_ConfigurationSpec_To_v1alpha1_ConfigurationSpec(a.(*ConfigurationSpec), b.(*v1alpha1.ConfigurationSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*TestSpec)(nil), (*v1alpha1.TestSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_TestSpec_To_v1alpha1_TestSpec(a.(*TestSpec), b.(*v1alpha1.TestSpec), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha2_ActionCheck_To_v1alpha1_ActionCheck(in *ActionCheck, out *v1alpha1.ActionCheck, s conversion.Scope) error {
	out.Check = (*policyv1alpha1.Any)(unsafe.Pointer(in.Check))
	return nil
}

// Convert_v1alpha2_ActionCheck_To_v1alpha1_ActionCheck is an autogenerated conversion function.
func Convert_v1alpha2_ActionCheck_To_v1alpha1_ActionCheck(in *ActionCheck, out *v1alpha1.ActionCheck, s conversion.Scope) error {
	return autoConvert_v1alpha2_ActionCheck_To_v1alpha1_ActionCheck(in, out, s)
}

func autoConvert_v1alpha1_ActionCheck_To_v1alpha2_ActionCheck(in *v1alpha1.ActionCheck, out *ActionCheck, s conversion.Scope) error {
	out.Check = (*policyv1alpha1.Any)(unsafe.Pointer(in.Check))
	return nil
}

// Convert_v1alpha1_ActionCheck_To_v1alpha2_ActionCheck is an autogenerated conversion function.
func Convert_v1alpha1_ActionCheck_To_v1alpha2_ActionCheck(in *v1alpha1.ActionCheck, out *ActionCheck, s conversion.Scope) error {
	return autoConvert_v1alpha1_ActionCheck_To_v1alpha2_ActionCheck(in, out, s)
}

func autoConvert_v1alpha2_ActionCheckRef_To_v1alpha1_ActionCheckRef(in *ActionCheckRef, out *v1alpha1.ActionCheckRef, s conversion.Scope) error {
	if err := Convert_v1alpha2_FileRef_To_v1alpha1_FileRef(&in.FileRef, &out.FileRef, s); err != nil {
		return err
	}
	out.Check = (*policyv1alpha1.Any)(unsafe.Pointer(in.Check))
	out.Template = (*bool)(unsafe.Pointer(in.Template))
	return nil
}

// Convert_v1alpha2_ActionCheckRef_To_v1alpha1_ActionCheckRef is an autogenerated conversion function.
func Convert_v1alpha2_ActionCheckRef_To_v1alpha1_ActionCheckRef(in *ActionCheckRef, out *v1alpha1.ActionCheckRef, s conversion.Scope) error {
	return autoConvert_v1alpha2_ActionCheckRef_To_v1alpha1_ActionCheckRef(in, out, s)
}

func autoConvert_v1alpha1_ActionCheckRef_To_v1alpha2_ActionCheckRef(in *v1alpha1.ActionCheckRef, out *ActionCheckRef, s conversion.Scope) error {
	if err := Convert_v1alpha1_FileRef_To_v1alpha2_FileRef(&in.FileRef, &out.FileRef, s); err != nil {
		return err
	}
	out.Check = (*policyv1alpha1.Any)(unsafe.Pointer(in.Check))
	out.Template = (*bool)(unsafe.Pointer(in.Template))
	return nil
}

// Convert_v1alpha1_ActionCheckRef_To_v1alpha2_ActionCheckRef is an autogenerated conversion function.
func Convert_v1alpha1_ActionCheckRef_To_v1alpha2_ActionCheckRef(in *v1alpha1.ActionCheckRef, out *ActionCheckRef, s conversion.Scope) error {
	return autoConvert_v1alpha1_ActionCheckRef_To_v1alpha2_ActionCheckRef(in, out, s)
}

func autoConvert_v1alpha2_ActionDryRun_To_v1alpha1_ActionDryRun(in *ActionDryRun, out *v1alpha1.ActionDryRun, s conversion.Scope) error {
	out.DryRun = (*bool)(unsafe.Pointer(in.DryRun))
	return nil
}

// Convert_v1alpha2_ActionDryRun_To_v1alpha1_ActionDryRun is an autogenerated conversion function.
func Convert_v1alpha2_ActionDryRun_To_v1alpha1_ActionDryRun(in *ActionDryRun, out *v1alpha1.ActionDryRun, s conversion.Scope) error {
	return autoConvert_v1alpha2_ActionDryRun_To_v1alpha1_ActionDryRun(in, out, s)
}

func autoConvert_v1alpha1_ActionDryRun_To_v1alpha2_ActionDryRun(in *v1alpha1.ActionDryRun, out *ActionDryRun, s conversion.Scope) error {
	out.DryRun = (*bool)(unsafe.Pointer(in.DryRun))
	return nil
}

// Convert_v1alpha1_ActionDryRun_To_v1alpha2_ActionDryRun is an autogenerated conversion function.
func Convert_v1alpha1_ActionDryRun_To_v1alpha2_ActionDryRun(in *v1alpha1.ActionDryRun, out *ActionDryRun, s conversion.Scope) error {
	return autoConvert_v1alpha1_ActionDryRun_To_v1alpha2_ActionDryRun(in, out, s)
}

func autoConvert_v1alpha2_ActionEnv_To_v1alpha1_ActionEnv(in *ActionEnv, out *v1alpha1.ActionEnv, s conversion.Scope) error {
	out.Env = *(*[]v1alpha1.Binding)(unsafe.Pointer(&in.Env))
	out.SkipLogOutput = in.SkipLogOutput
	return nil
}

// Convert_v1alpha2_ActionEnv_To_v1alpha1_ActionEnv is an autogenerated conversion function.
func Convert_v1alpha2_ActionEnv_To_v1alpha1_ActionEnv(in *ActionEnv, out *v1alpha1.ActionEnv, s conversion.Scope) error {
	return autoConvert_v1alpha2_ActionEnv_To_v1alpha1_ActionEnv(in, out, s)
}

func autoConvert_v1alpha1_ActionEnv_To_v1alpha2_ActionEnv(in *v1alpha1.ActionEnv, out *ActionEnv, s conversion.Scope) error {
	out.Env = *(*[]v1alpha1.Binding)(unsafe.Pointer(&in.Env))
	out.SkipLogOutput = in.SkipLogOutput
	return nil
}

// Convert_v1alpha1_ActionEnv_To_v1alpha2_ActionEnv is an autogenerated conversion function.
func Convert_v1alpha1_ActionEnv_To_v1alpha2_ActionEnv(in *v1alpha1.ActionEnv, out *ActionEnv, s conversion.Scope) error {
	return autoConvert_v1alpha1_ActionEnv_To_v1alpha2_ActionEnv(in, out, s)
}

func autoConvert_v1alpha2_ActionExpectations_To_v1alpha1_ActionExpectations(in *ActionExpectations, out *v1alpha1.ActionExpectations, s conversion.Scope) error {
	out.Expect = *(*[]v1alpha1.Expectation)(unsafe.Pointer(&in.Expect))
	return nil
}

// Convert_v1alpha2_ActionExpectations_To_v1alpha1_ActionExpectations is an autogenerated conversion function.
func Convert_v1alpha2_ActionExpectations_To_v1alpha1_ActionExpectations(in *ActionExpectations, out *v1alpha1.ActionExpectations, s conversion.Scope) error {
	return autoConvert_v1alpha2_ActionExpectations_To_v1alpha1_ActionExpectations(in, out, s)
}

func autoConvert_v1alpha1_ActionExpectations_To_v1alpha2_ActionExpectations(in *v1alpha1.ActionExpectations, out *ActionExpectations, s conversion.Scope) error {
	out.Expect = *(*[]v1alpha1.Expectation)(unsafe.Pointer(&in.Expect))
	return nil
}

// Convert_v1alpha1_ActionExpectations_To_v1alpha2_ActionExpectations is an autogenerated conversion function.
func Convert_v1alpha1_ActionExpectations_To_v1alpha2_ActionExpectations(in *v1alpha1.ActionExpectations, out *ActionExpectations, s conversion.Scope) error {
	return autoConvert_v1alpha1_ActionExpectations_To_v1alpha2_ActionExpectations(in, out, s)
}

func autoConvert_v1alpha2_ActionFormat_To_v1alpha1_ActionFormat(in *ActionFormat, out *v1alpha1.ActionFormat, s conversion.Scope) error {
	out.Format = v1alpha1.Format(in.Format)
	return nil
}

// Convert_v1alpha2_ActionFormat_To_v1alpha1_ActionFormat is an autogenerated conversion function.
func Convert_v1alpha2_ActionFormat_To_v1alpha1_ActionFormat(in *ActionFormat, out *v1alpha1.ActionFormat, s conversion.Scope) error {
	return autoConvert_v1alpha2_ActionFormat_To_v1alpha1_ActionFormat(in, out, s)
}

func autoConvert_v1alpha1_ActionFormat_To_v1alpha2_ActionFormat(in *v1alpha1.ActionFormat, out *ActionFormat, s conversion.Scope) error {
	out.Format = v1alpha1.Format(in.Format)
	return nil
}

// Convert_v1alpha1_ActionFormat_To_v1alpha2_ActionFormat is an autogenerated conversion function.
func Convert_v1alpha1_ActionFormat_To_v1alpha2_ActionFormat(in *v1alpha1.ActionFormat, out *ActionFormat, s conversion.Scope) error {
	return autoConvert_v1alpha1_ActionFormat_To_v1alpha2_ActionFormat(in, out, s)
}

func autoConvert_v1alpha2_ActionInlineResource_To_v1alpha1_ActionInlineResource(in *ActionInlineResource, out *v1alpha1.ActionInlineResource, s conversion.Scope) error {
	out.Resource = (*unstructured.Unstructured)(unsafe.Pointer(in.Resource))
	return nil
}

// Convert_v1alpha2_ActionInlineResource_To_v1alpha1_ActionInlineResource is an autogenerated conversion function.
func Convert_v1alpha2_ActionInlineResource_To_v1alpha1_ActionInlineResource(in *ActionInlineResource, out *v1alpha1.ActionInlineResource, s conversion.Scope) error {
	return autoConvert_v1alpha2_ActionInlineResource_To_v1alpha1_ActionInlineResource(in, out, s)
}

func autoConvert_v1alpha1_ActionInlineResource_To_v1alpha2_ActionInlineResource(in *v1alpha1.ActionInlineResource, out *ActionInlineResource, s conversion.Scope) error {
	out.Resource = (*unstructured.Unstructured)(unsafe.Pointer(in.Resource))
	return nil
}

// Convert_v1alpha1_ActionInlineResource_To_v1alpha2_ActionInlineResource is an autogenerated conversion function.
func Convert_v1alpha1_ActionInlineResource_To_v1alpha2_ActionInlineResource(in *v1alpha1.ActionInlineResource, out *ActionInlineResource, s conversion.Scope) error {
	return autoConvert_v1alpha1_ActionInlineResource_To_v1alpha2_ActionInlineResource(in, out, s)
}

func autoConvert_v1alpha2_ActionObject_To_v1alpha1_ActionObject(in *ActionObject, out *v1alpha1.ActionObject, s conversion.Scope) error {
	out.ObjectType = in.ObjectType
	if err := Convert_v1alpha2_ActionObjectSelector_To_v1alpha1_ActionObjectSelector(&in.ActionObjectSelector, &out.ActionObjectSelector, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha2_ActionObject_To_v1alpha1_ActionObject is an autogenerated conversion function.
func Convert_v1alpha2_ActionObject_To_v1alpha1_ActionObject(in *ActionObject, out *v1alpha1.ActionObject, s conversion.Scope) error {
	return autoConvert_v1alpha2_ActionObject_To_v1alpha1_ActionObject(in, out, s)
}

func autoConvert_v1alpha1_ActionObject_To_v1alpha2_ActionObject(in *v1alpha1.ActionObject, out *ActionObject, s conversion.Scope) error {
	out.ObjectType = in.ObjectType
	if err := Convert_v1alpha1_ActionObjectSelector_To_v1alpha2_ActionObjectSelector(&in.ActionObjectSelector, &out.ActionObjectSelector, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_ActionObject_To_v1alpha2_ActionObject is an autogenerated conversion function.
func Convert_v1alpha1_ActionObject_To_v1alpha2_ActionObject(in *v1alpha1.ActionObject, out *ActionObject, s conversion.Scope) error {
	return autoConvert_v1alpha1_ActionObject_To_v1alpha2_ActionObject(in, out, s)
}

func autoConvert_v1alpha2_ActionObjectSelector_To_v1alpha1_ActionObjectSelector(in *ActionObjectSelector, out *v1alpha1.ActionObjectSelector, s conversion.Scope) error {
	out.ObjectName = in.ObjectName
	out.Selector = in.Selector
	return nil
}

// Convert_v1alpha2_ActionObjectSelector_To_v1alpha1_ActionObjectSelector is an autogenerated conversion function.
func Convert_v1alpha2_ActionObjectSelector_To_v1alpha1_ActionObjectSelector(in *ActionObjectSelector, out *v1alpha1.ActionObjectSelector, s conversion.Scope) error {
	return autoConvert_v1alpha2_ActionObjectSelector_To_v1alpha1_ActionObjectSelector(in, out, s)
}

func autoConvert_v1alpha1_ActionObjectSelector_To_v1alpha2_ActionObjectSelector(in *v1alpha1.ActionObjectSelector, out *ActionObjectSelector, s conversion.Scope) error {
	out.ObjectName = in.ObjectName
	out.Selector = in.Selector
	return nil
}

// Convert_v1alpha1_ActionObjectSelector_To_v1alpha2_ActionObjectSelector is an autogenerated conversion function.
func Convert_v1alpha1_ActionObjectSelector_To_v1alpha2_ActionObjectSelector(in *v1alpha1.ActionObjectSelector, out *ActionObjectSelector, s conversion.Scope) error {
	return autoConvert_v1alpha1_ActionObjectSelector_To_v1alpha2_ActionObjectSelector(in, out, s)
}

func autoConvert_v1alpha2_ActionResourceRef_To_v1alpha1_ActionResourceRef(in *ActionResourceRef, out *v1alpha1.ActionResourceRef, s conversion.Scope) error {
	if err := Convert_v1alpha2_FileRef_To_v1alpha1_FileRef(&in.FileRef, &out.FileRef, s); err != nil {
		return err
	}
	out.Resource = (*unstructured.Unstructured)(unsafe.Pointer(in.Resource))
	out.Template = (*bool)(unsafe.Pointer(in.Template))
	return nil
}

// Convert_v1alpha2_ActionResourceRef_To_v1alpha1_ActionResourceRef is an autogenerated conversion function.
func Convert_v1alpha2_ActionResourceRef_To_v1alpha1_ActionResourceRef(in *ActionResourceRef, out *v1alpha1.ActionResourceRef, s conversion.Scope) error {
	return autoConvert_v1alpha2_ActionResourceRef_To_v1alpha1_ActionResourceRef(in, out, s)
}

func autoConvert_v1alpha1_ActionResourceRef_To_v1alpha2_ActionResourceRef(in *v1alpha1.ActionResourceRef, out *ActionResourceRef, s conversion.Scope) error {
	if err := Convert_v1alpha1_FileRef_To_v1alpha2_FileRef(&in.FileRef, &out.FileRef, s); err != nil {
		return err
	}
	out.Resource = (*unstructured.Unstructured)(unsafe.Pointer(in.Resource))
	out.Template = (*bool)(unsafe.Pointer(in.Template))
	return nil
}

// Convert_v1alpha1_ActionResourceRef_To_v1alpha2_ActionResourceRef is an autogenerated conversion function.
func Convert_v1alpha1_ActionResourceRef_To_v1alpha2_ActionResourceRef(in *v1alpha1.ActionResourceRef, out *ActionResourceRef, s conversion.Scope) error {
	return autoConvert_v1alpha1_ActionResourceRef_To_v1alpha2_ActionResourceRef(in, out, s)
}

func autoConvert_v1alpha2_ActionTimeout_To_v1alpha1_ActionTimeout(in *ActionTimeout, out *v1alpha1.ActionTimeout, s conversion.Scope) error {
	out.Timeout = (*v1.Duration)(unsafe.Pointer(in.Timeout))
	return nil
}

// Convert_v1alpha2_ActionTimeout_To_v1alpha1_ActionTimeout is an autogenerated conversion function.
func Convert_v1alpha2_ActionTimeout_To_v1alpha1_ActionTimeout(in *ActionTimeout, out *v1alpha1.ActionTimeout, s conversion.Scope) error {
	return autoConvert_v1alpha2_ActionTimeout_To_v1alpha1_ActionTimeout(in, out, s)
}

func autoConvert_v1alpha1_ActionTimeout_To_v1alpha2_ActionTimeout(in *v1alpha1.ActionTimeout, out *ActionTimeout, s conversion.Scope) error {
	out.Timeout = (*v1.Duration)(unsafe.Pointer(in.Timeout))
	return nil
}

// Convert_v1alpha1_ActionTimeout_To_v1alpha2_ActionTimeout is an autogenerated conversion function.
func Convert_v1alpha1_ActionTimeout_To_v1alpha2_ActionTimeout(in *v1alpha1.ActionTimeout, out *ActionTimeout, s conversion.Scope) error {
	return autoConvert_v1alpha1_ActionTimeout_To_v1alpha2_ActionTimeout(in, out, s)
}

func autoConvert_v1alpha2_Configuration_To_v1alpha1_Configuration(in *Configuration, out *v1alpha1.Configuration, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha2_ConfigurationSpec_To_v1alpha1_ConfigurationSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha2_Configuration_To_v1alpha1_Configuration is an autogenerated conversion function.
func Convert_v1alpha2_Configuration_To_v1alpha1_Configuration(in *Configuration, out *v1alpha1.Configuration, s conversion.Scope) error {
	return autoConvert_v1alpha2_Configuration_To_v1alpha1_Configuration(in, out, s)
}

func autoConvert_v1alpha1_Configuration_To_v1alpha2_Configuration(in *v1alpha1.Configuration, out *Configuration, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_ConfigurationSpec_To_v1alpha2_ConfigurationSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_Configuration_To_v1alpha2_Configuration is an autogenerated conversion function.
func Convert_v1alpha1_Configuration_To_v1alpha2_Configuration(in *v1alpha1.Configuration, out *Configuration, s conversion.Scope) error {
	return autoConvert_v1alpha1_Configuration_To_v1alpha2_Configuration(in, out, s)
}

func autoConvert_v1alpha2_FileRef_To_v1alpha1_FileRef(in *FileRef, out *v1alpha1.FileRef, s conversion.Scope) error {
	out.File = in.File
	return nil
}

// Convert_v1alpha2_FileRef_To_v1alpha1_FileRef is an autogenerated conversion function.
func Convert_v1alpha2_FileRef_To_v1alpha1_FileRef(in *FileRef, out *v1alpha1.FileRef, s conversion.Scope) error {
	return autoConvert_v1alpha2_FileRef_To_v1alpha1_FileRef(in, out, s)
}

func autoConvert_v1alpha1_FileRef_To_v1alpha2_FileRef(in *v1alpha1.FileRef, out *FileRef, s conversion.Scope) error {
	out.File = in.File
	return nil
}

// Convert_v1alpha1_FileRef_To_v1alpha2_FileRef is an autogenerated conversion function.
func Convert_v1alpha1_FileRef_To_v1alpha2_FileRef(in *v1alpha1.FileRef, out *FileRef, s conversion.Scope) error {
	return autoConvert_v1alpha1_FileRef_To_v1alpha2_FileRef(in, out, s)
}

func autoConvert_v1alpha2_Test_To_v1alpha1_Test(in *Test, out *v1alpha1.Test, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha2_TestSpec_To_v1alpha1_TestSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha2_Test_To_v1alpha1_Test is an autogenerated conversion function.
func Convert_v1alpha2_Test_To_v1alpha1_Test(in *Test, out *v1alpha1.Test, s conversion.Scope) error {
	return autoConvert_v1alpha2_Test_To_v1alpha1_Test(in, out, s)
}

func autoConvert_v1alpha1_Test_To_v1alpha2_Test(in *v1alpha1.Test, out *Test, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_TestSpec_To_v1alpha2_TestSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_Test_To_v1alpha2_Test is an autogenerated conversion function.
func Convert_v1alpha1_Test_To_v1alpha2_Test(in *v1alpha1.Test, out *Test, s conversion.Scope) error {
	return autoConvert_v1alpha1_Test_To_v1alpha2_Test(in, out, s)
}
