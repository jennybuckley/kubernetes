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

package internal

import (
	"reflect"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TestFieldsRoundTrip tests that a fields trie can be round tripped as a path set
func TestFieldsRoundTrip(t *testing.T) {
	tests := []metav1.Fields{
		EmptyFields,
	}

	for _, test := range tests {
		set, err := FieldsToSet(test)
		if err != nil {
			t.Fatalf("Failed to create path set: %v", err)
		}
		output, err := SetToFields(set)
		if err != nil {
			t.Fatalf("Failed to create fields trie from path set: %v", err)
		}
		if !reflect.DeepEqual(test, output) {
			t.Fatalf("Expected round-trip:\ninput: %v\noutput: %v", test, output)
		}
	}
}
