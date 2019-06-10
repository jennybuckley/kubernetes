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
	"encoding/json"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"sigs.k8s.io/structured-merge-diff/fieldpath"
)

var EmptyFields metav1.Fields = func() metav1.Fields {
	f, err := json.Marshal(fieldpath.NewSet())
	if err != nil {
		panic("should never happen")
	}
	return metav1.Fields(string(f))
}()

// FieldsToSet creates a set paths from an input trie of fields
func FieldsToSet(f metav1.Fields) (fieldpath.Set, error) {
	set := fieldpath.Set{}
	err := json.Unmarshal([]byte(string(f)), &set)
	return set, err
}

// SetToFields creates a trie of fields from an input set of paths
func SetToFields(s fieldpath.Set) (metav1.Fields, error) {
	f, err := json.Marshal(s)
	if err != nil {
		return EmptyFields, err
	}
	return metav1.Fields(string(f)), nil
}
