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

type fieldMap map[string]fieldMap

func newFields() fieldMap {
	return fieldMap{}
}

func fieldsSet(f fieldMap, path fieldpath.Path, set *fieldpath.Set) error {
	if len(f) == 0 {
		set.Insert(path)
	}
	for k := range f {
		if k == "." {
			set.Insert(path)
			continue
		}
		pe, err := NewPathElement(k)
		if err != nil {
			return err
		}
		path = append(path, pe)
		err = fieldsSet(f[k], path, set)
		if err != nil {
			return err
		}
		path = path[:len(path)-1]
	}
	return nil
}

// fieldMapToSet creates a set paths from an input trie of fields
func fieldMapToSet(f fieldMap) (fieldpath.Set, error) {
	set := fieldpath.Set{}
	return set, fieldsSet(f, fieldpath.Path{}, &set)
}

func removeUselessDots(f fieldMap) fieldMap {
	if _, ok := f["."]; ok && len(f) == 1 {
		delete(f, ".")
		return f
	}
	for k, tf := range f {
		f[k] = removeUselessDots(tf)
	}
	return f
}

// setToFieldMap creates a trie of fields from an input set of paths
func setToFieldMap(s fieldpath.Set) (fieldMap, error) {
	var err error
	f := newFields()
	s.Iterate(func(path fieldpath.Path) {
		if err != nil {
			return
		}
		tf := f
		for _, pe := range path {
			var str string
			str, err = PathElementString(pe)
			if err != nil {
				break
			}
			if _, ok := tf[str]; ok {
				tf = tf[str]
			} else {
				tf[str] = newFields()
				tf = tf[str]
			}
		}
		tf["."] = newFields()
	})
	f = removeUselessDots(f)
	return f, err
}

var EmptyFields metav1.Fields = func() metav1.Fields {
	fm, err := setToFieldMap(*fieldpath.NewSet())
	if err != nil {
		panic("should never happen")
	}
	f, err := json.Marshal(fm)
	if err != nil {
		panic("should never happen")
	}
	return metav1.Fields(string(f))
}()

// FieldsToSet creates a set paths from an input trie of fields
func FieldsToSet(f metav1.Fields) (fieldpath.Set, error) {
	fm := fieldMap{}
	err := json.Unmarshal([]byte(string(f)), &fm)
	if err != nil {
		return fieldpath.Set{}, err
	}
	return fieldMapToSet(fm)
}

// SetToFields creates a trie of fields from an input set of paths
func SetToFields(s fieldpath.Set) (metav1.Fields, error) {
	fm, err := setToFieldMap(s)
	if err != nil {
		return EmptyFields, err
	}
	f, err := json.Marshal(fm)
	if err != nil {
		return EmptyFields, err
	}
	return metav1.Fields(string(f)), nil
}
