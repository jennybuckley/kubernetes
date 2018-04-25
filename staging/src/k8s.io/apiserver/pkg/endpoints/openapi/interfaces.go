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

package openapi

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/kube-openapi/pkg/util/proto"
)

// Getter is an interface for fetching openapi specs and parsing them into an Resources struct
type Getter interface {
	// OpenAPIData returns the parsed OpenAPIData
	Get() (Resources, error)
}

// Resources interface describe a resources provider, that can give you
// resource based on group-version-kind.
type Resources interface {
	LookupResource(gvk schema.GroupVersionKind) proto.Schema
}
