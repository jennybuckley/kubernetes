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

// Code generated by deepcopy-gen. DO NOT EDIT.

package apiextensions

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomResourceDefinition) DeepCopyInto(out *CustomResourceDefinition) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomResourceDefinition.
func (in *CustomResourceDefinition) DeepCopy() *CustomResourceDefinition {
	if in == nil {
		return nil
	}
	out := new(CustomResourceDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CustomResourceDefinition) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomResourceDefinitionCondition) DeepCopyInto(out *CustomResourceDefinitionCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomResourceDefinitionCondition.
func (in *CustomResourceDefinitionCondition) DeepCopy() *CustomResourceDefinitionCondition {
	if in == nil {
		return nil
	}
	out := new(CustomResourceDefinitionCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomResourceDefinitionList) DeepCopyInto(out *CustomResourceDefinitionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CustomResourceDefinition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomResourceDefinitionList.
func (in *CustomResourceDefinitionList) DeepCopy() *CustomResourceDefinitionList {
	if in == nil {
		return nil
	}
	out := new(CustomResourceDefinitionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CustomResourceDefinitionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomResourceDefinitionNames) DeepCopyInto(out *CustomResourceDefinitionNames) {
	*out = *in
	if in.ShortNames != nil {
		in, out := &in.ShortNames, &out.ShortNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomResourceDefinitionNames.
func (in *CustomResourceDefinitionNames) DeepCopy() *CustomResourceDefinitionNames {
	if in == nil {
		return nil
	}
	out := new(CustomResourceDefinitionNames)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomResourceDefinitionSpec) DeepCopyInto(out *CustomResourceDefinitionSpec) {
	*out = *in
	in.Names.DeepCopyInto(&out.Names)
	if in.Validation != nil {
		in, out := &in.Validation, &out.Validation
		if *in == nil {
			*out = nil
		} else {
			*out = new(CustomResourceValidation)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Subresources != nil {
		in, out := &in.Subresources, &out.Subresources
		if *in == nil {
			*out = nil
		} else {
			*out = new(CustomResourceSubresources)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomResourceDefinitionSpec.
func (in *CustomResourceDefinitionSpec) DeepCopy() *CustomResourceDefinitionSpec {
	if in == nil {
		return nil
	}
	out := new(CustomResourceDefinitionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomResourceDefinitionStatus) DeepCopyInto(out *CustomResourceDefinitionStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]CustomResourceDefinitionCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.AcceptedNames.DeepCopyInto(&out.AcceptedNames)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomResourceDefinitionStatus.
func (in *CustomResourceDefinitionStatus) DeepCopy() *CustomResourceDefinitionStatus {
	if in == nil {
		return nil
	}
	out := new(CustomResourceDefinitionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomResourceSubresourceScale) DeepCopyInto(out *CustomResourceSubresourceScale) {
	*out = *in
	if in.LabelSelectorPath != nil {
		in, out := &in.LabelSelectorPath, &out.LabelSelectorPath
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomResourceSubresourceScale.
func (in *CustomResourceSubresourceScale) DeepCopy() *CustomResourceSubresourceScale {
	if in == nil {
		return nil
	}
	out := new(CustomResourceSubresourceScale)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomResourceSubresourceStatus) DeepCopyInto(out *CustomResourceSubresourceStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomResourceSubresourceStatus.
func (in *CustomResourceSubresourceStatus) DeepCopy() *CustomResourceSubresourceStatus {
	if in == nil {
		return nil
	}
	out := new(CustomResourceSubresourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomResourceSubresources) DeepCopyInto(out *CustomResourceSubresources) {
	*out = *in
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		if *in == nil {
			*out = nil
		} else {
			*out = new(CustomResourceSubresourceStatus)
			**out = **in
		}
	}
	if in.Scale != nil {
		in, out := &in.Scale, &out.Scale
		if *in == nil {
			*out = nil
		} else {
			*out = new(CustomResourceSubresourceScale)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomResourceSubresources.
func (in *CustomResourceSubresources) DeepCopy() *CustomResourceSubresources {
	if in == nil {
		return nil
	}
	out := new(CustomResourceSubresources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomResourceValidation) DeepCopyInto(out *CustomResourceValidation) {
	*out = *in
	if in.OpenAPIV3Schema != nil {
		in, out := &in.OpenAPIV3Schema, &out.OpenAPIV3Schema
		if *in == nil {
			*out = nil
		} else {
			*out = new(JSONSchemaProps)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomResourceValidation.
func (in *CustomResourceValidation) DeepCopy() *CustomResourceValidation {
	if in == nil {
		return nil
	}
	out := new(CustomResourceValidation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalDocumentation) DeepCopyInto(out *ExternalDocumentation) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalDocumentation.
func (in *ExternalDocumentation) DeepCopy() *ExternalDocumentation {
	if in == nil {
		return nil
	}
	out := new(ExternalDocumentation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JSONSchemaProps) DeepCopyInto(out *JSONSchemaProps) {
	clone := in.DeepCopy()
	*out = *clone
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JSONSchemaPropsOrArray) DeepCopyInto(out *JSONSchemaPropsOrArray) {
	*out = *in
	if in.Schema != nil {
		in, out := &in.Schema, &out.Schema
		if *in == nil {
			*out = nil
		} else {
			*out = new(JSONSchemaProps)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.JSONSchemas != nil {
		in, out := &in.JSONSchemas, &out.JSONSchemas
		*out = make([]JSONSchemaProps, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JSONSchemaPropsOrArray.
func (in *JSONSchemaPropsOrArray) DeepCopy() *JSONSchemaPropsOrArray {
	if in == nil {
		return nil
	}
	out := new(JSONSchemaPropsOrArray)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JSONSchemaPropsOrBool) DeepCopyInto(out *JSONSchemaPropsOrBool) {
	*out = *in
	if in.Schema != nil {
		in, out := &in.Schema, &out.Schema
		if *in == nil {
			*out = nil
		} else {
			*out = new(JSONSchemaProps)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JSONSchemaPropsOrBool.
func (in *JSONSchemaPropsOrBool) DeepCopy() *JSONSchemaPropsOrBool {
	if in == nil {
		return nil
	}
	out := new(JSONSchemaPropsOrBool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JSONSchemaPropsOrStringArray) DeepCopyInto(out *JSONSchemaPropsOrStringArray) {
	*out = *in
	if in.Schema != nil {
		in, out := &in.Schema, &out.Schema
		if *in == nil {
			*out = nil
		} else {
			*out = new(JSONSchemaProps)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Property != nil {
		in, out := &in.Property, &out.Property
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JSONSchemaPropsOrStringArray.
func (in *JSONSchemaPropsOrStringArray) DeepCopy() *JSONSchemaPropsOrStringArray {
	if in == nil {
		return nil
	}
	out := new(JSONSchemaPropsOrStringArray)
	in.DeepCopyInto(out)
	return out
}
