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

package apiserver

import (
	"testing"

	"k8s.io/apimachinery/pkg/types"
)

func TestApplyAlsoCreates(t *testing.T) {
	_, client, closeFn := setup(t)
	defer closeFn()

	_, err := client.CoreV1().RESTClient().Patch(types.StrategicMergePatchType).
		Namespace("default").
		Resource("pods").
		Name("test-pod").
		Body([]byte(`{"name": "test-pod"}`)).
		Do().
		Get()
	if err != nil {
		t.Fatalf("Failed to create object using Apply patch: %v", err)
	}

	obj, err := client.CoreV1().RESTClient().Get().Namespace("").Resource("pods").Name("test-pod").Do().Get()
	if err != nil {
		t.Fatalf("Failed to retrieve object: %v", err)
	}
}
