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
	"fmt"
	"reflect"
	"testing"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"sigs.k8s.io/yaml"
)

// TestRoundTripManagedFields will roundtrip ManagedFields from the wire format
// (api format) to the format used by sigs.k8s.io/structured-merge-diff and back
func TestRoundTripManagedFields(t *testing.T) {
	tests := []string{
		`- apiVersion: v1
  fields: '{"Members":{},"Children":{"Members":{"[=3.1415]":{"PathElement":{"Value":{"FloatValue":3.1415}},"Set":{"Members":{"Members":{".pi":{"FieldName":"pi"}}},"Children":{}}},"[=3]":{"PathElement":{"Value":{"FloatValue":3}},"Set":{"Members":{"Members":{".alsoPi":{"FieldName":"alsoPi"}}},"Children":{}}},"[=false]":{"PathElement":{"Value":{"BooleanValue":false}},"Set":{"Members":{"Members":{".notTrue":{"FieldName":"notTrue"}}},"Children":{}}}}}}'
- apiVersion: v1beta1
  fields: '{"Members":{},"Children":{"Members":{"[5]":{"PathElement":{"Index":5},"Set":{"Members":{"Members":{".i":{"FieldName":"i"}}},"Children":{}}}}}}'
  manager: foo
  operation: Update
  time: "2011-12-13T14:15:16Z"
`,
		`- apiVersion: v1
  fields: '{"Members":{},"Children":{"Members":{".spec":{"PathElement":{"FieldName":"spec"},"Set":{"Members":{},"Children":{"Members":{".containers":{"PathElement":{"FieldName":"containers"},"Set":{"Members":{},"Children":{"Members":{"[name=\"c\"]":{"PathElement":{"Key":[{"Name":"name","Value":{"StringValue":"c"}}]},"Set":{"Members":{"Members":{".image":{"FieldName":"image"},".name":{"FieldName":"name"}}},"Children":{}}}}}}}}}}}}}}'
  manager: foo
  operation: Apply
`,
		`- apiVersion: v1
  fields: '{"Members":{"Members":{".apiVersion":{"FieldName":"apiVersion"},".kind":{"FieldName":"kind"}}},"Children":{"Members":{".metadata":{"PathElement":{"FieldName":"metadata"},"Set":{"Members":{"Members":{".name":{"FieldName":"name"}}},"Children":{"Members":{".labels":{"PathElement":{"FieldName":"labels"},"Set":{"Members":{"Members":{".app":{"FieldName":"app"}}},"Children":{}}}}}}},".spec":{"PathElement":{"FieldName":"spec"},"Set":{"Members":{"Members":{".replicas":{"FieldName":"replicas"}}},"Children":{"Members":{".selector":{"PathElement":{"FieldName":"selector"},"Set":{"Members":{},"Children":{"Members":{".matchLabels":{"PathElement":{"FieldName":"matchLabels"},"Set":{"Members":{"Members":{".app":{"FieldName":"app"}}},"Children":{}}}}}}},".template":{"PathElement":{"FieldName":"template"},"Set":{"Members":{},"Children":{"Members":{".medatada":{"PathElement":{"FieldName":"medatada"},"Set":{"Members":{},"Children":{"Members":{".labels":{"PathElement":{"FieldName":"labels"},"Set":{"Members":{"Members":{".app":{"FieldName":"app"}}},"Children":{}}}}}}},".spec":{"PathElement":{"FieldName":"spec"},"Set":{"Members":{},"Children":{"Members":{".containers":{"PathElement":{"FieldName":"containers"},"Set":{"Members":{"Members":{"[name=\"nginx\"]":{"Key":[{"Name":"name","Value":{"StringValue":"nginx"}}]}}},"Children":{"Members":{"[name=\"nginx\"]":{"PathElement":{"Key":[{"Name":"name","Value":{"StringValue":"nginx"}}]},"Set":{"Members":{"Members":{".image":{"FieldName":"image"},".name":{"FieldName":"name"}}},"Children":{"Members":{".ports":{"PathElement":{"FieldName":"ports"},"Set":{"Members":{},"Children":{"Members":{"[0]":{"PathElement":{"Index":0},"Set":{"Members":{"Members":{".containerPort":{"FieldName":"containerPort"}}},"Children":{}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}'
  manager: foo
  operation: Update
`,
		`- apiVersion: v1
  fields: '{"Members":{"Members":{".allowVolumeExpansion":{"FieldName":"allowVolumeExpansion"},".apiVersion":{"FieldName":"apiVersion"},".kind":{"FieldName":"kind"},".provisioner":{"FieldName":"provisioner"}}},"Children":{"Members":{".metadata":{"PathElement":{"FieldName":"metadata"},"Set":{"Members":{"Members":{".name":{"FieldName":"name"}}},"Children":{"Members":{".parameters":{"PathElement":{"FieldName":"parameters"},"Set":{"Members":{"Members":{".resturl":{"FieldName":"resturl"},".restuser":{"FieldName":"restuser"},".secretName":{"FieldName":"secretName"},".secretNamespace":{"FieldName":"secretNamespace"}}},"Children":{}}}}}}}}}}'
  manager: foo
  operation: Apply
`,
		`- apiVersion: v1
  fields: '{"Members":{"Members":{".apiVersion":{"FieldName":"apiVersion"},".kind":{"FieldName":"kind"}}},"Children":{"Members":{".metadata":{"PathElement":{"FieldName":"metadata"},"Set":{"Members":{"Members":{".name":{"FieldName":"name"}}},"Children":{}}},".spec":{"PathElement":{"FieldName":"spec"},"Set":{"Members":{"Members":{".group":{"FieldName":"group"},".scope":{"FieldName":"scope"}}},"Children":{"Members":{".names":{"PathElement":{"FieldName":"names"},"Set":{"Members":{"Members":{".kind":{"FieldName":"kind"},".plural":{"FieldName":"plural"},".singular":{"FieldName":"singular"}}},"Children":{"Members":{".shortNames":{"PathElement":{"FieldName":"shortNames"},"Set":{"Members":{"Members":{"[0]":{"Index":0}}},"Children":{}}}}}}},".versions":{"PathElement":{"FieldName":"versions"},"Set":{"Members":{},"Children":{"Members":{"[name=\"v1\"]":{"PathElement":{"Key":[{"Name":"name","Value":{"StringValue":"v1"}}]},"Set":{"Members":{"Members":{".name":{"FieldName":"name"},".served":{"FieldName":"served"},".storage":{"FieldName":"storage"}}},"Children":{}}}}}}}}}}}}}}'
  manager: foo
  operation: Update
`,
	}

	for _, test := range tests {
		t.Run(test, func(t *testing.T) {
			var unmarshaled []metav1.ManagedFieldsEntry
			if err := yaml.Unmarshal([]byte(test), &unmarshaled); err != nil {
				t.Fatalf("did not expect yaml unmarshalling error but got: %v", err)
			}
			decoded, err := decodeManagedFields(unmarshaled)
			if err != nil {
				t.Fatalf("did not expect decoding error but got: %v", err)
			}
			encoded, err := encodeManagedFields(decoded)
			if err != nil {
				t.Fatalf("did not expect encoding error but got: %v", err)
			}
			marshaled, err := yaml.Marshal(&encoded)
			if err != nil {
				t.Fatalf("did not expect yaml marshalling error but got: %v", err)
			}
			if !reflect.DeepEqual(string(marshaled), test) {
				t.Fatalf("expected:\n%v\nbut got:\n%v", test, string(marshaled))
			}
		})
	}
}

func TestBuildManagerIdentifier(t *testing.T) {
	tests := []struct {
		managedFieldsEntry string
		expected           string
	}{
		{
			managedFieldsEntry: `
apiVersion: v1
manager: foo
operation: Update
time: "2001-02-03T04:05:06Z"
`,
			expected: "{\"manager\":\"foo\",\"operation\":\"Update\",\"apiVersion\":\"v1\",\"time\":\"2001-02-03T04:05:06Z\"}",
		},
		{
			managedFieldsEntry: `
apiVersion: v1
manager: foo
operation: Apply
time: "2001-02-03T04:05:06Z"
`,
			expected: "{\"manager\":\"foo\",\"operation\":\"Apply\"}",
		},
	}

	for _, test := range tests {
		t.Run(test.managedFieldsEntry, func(t *testing.T) {
			var unmarshaled metav1.ManagedFieldsEntry
			if err := yaml.Unmarshal([]byte(test.managedFieldsEntry), &unmarshaled); err != nil {
				t.Fatalf("did not expect yaml unmarshalling error but got: %v", err)
			}
			decoded, err := BuildManagerIdentifier(&unmarshaled)
			if err != nil {
				t.Fatalf("did not expect decoding error but got: %v", err)
			}
			if !reflect.DeepEqual(decoded, test.expected) {
				t.Fatalf("expected:\n%v\nbut got:\n%v", test.expected, decoded)
			}
		})
	}
}

func TestSortEncodedManagedFields(t *testing.T) {
	tests := []struct {
		name          string
		managedFields []metav1.ManagedFieldsEntry
		expected      []metav1.ManagedFieldsEntry
	}{
		{
			name:          "empty",
			managedFields: []metav1.ManagedFieldsEntry{},
			expected:      []metav1.ManagedFieldsEntry{},
		},
		{
			name:          "nil",
			managedFields: nil,
			expected:      nil,
		},
		{
			name: "remains untouched",
			managedFields: []metav1.ManagedFieldsEntry{
				{Manager: "a", Operation: metav1.ManagedFieldsOperationApply, Time: nil},
				{Manager: "a", Operation: metav1.ManagedFieldsOperationUpdate, Time: parseTimeOrPanic("2001-01-01T01:00:00Z")},
			},
			expected: []metav1.ManagedFieldsEntry{
				{Manager: "a", Operation: metav1.ManagedFieldsOperationApply, Time: nil},
				{Manager: "a", Operation: metav1.ManagedFieldsOperationUpdate, Time: parseTimeOrPanic("2001-01-01T01:00:00Z")},
			},
		},
		{
			name: "manager without time first",
			managedFields: []metav1.ManagedFieldsEntry{
				{Manager: "a", Operation: metav1.ManagedFieldsOperationApply, Time: nil},
				{Manager: "a", Operation: metav1.ManagedFieldsOperationApply, Time: parseTimeOrPanic("2001-01-01T01:00:00Z")},
			},
			expected: []metav1.ManagedFieldsEntry{
				{Manager: "a", Operation: metav1.ManagedFieldsOperationApply, Time: nil},
				{Manager: "a", Operation: metav1.ManagedFieldsOperationApply, Time: parseTimeOrPanic("2001-01-01T01:00:00Z")},
			},
		},
		{
			name: "manager without time first name last",
			managedFields: []metav1.ManagedFieldsEntry{
				{Manager: "a", Operation: metav1.ManagedFieldsOperationApply, Time: parseTimeOrPanic("2001-01-01T01:00:00Z")},
				{Manager: "b", Operation: metav1.ManagedFieldsOperationApply, Time: nil},
				{Manager: "a", Operation: metav1.ManagedFieldsOperationApply, Time: nil},
			},
			expected: []metav1.ManagedFieldsEntry{
				{Manager: "a", Operation: metav1.ManagedFieldsOperationApply, Time: nil},
				{Manager: "b", Operation: metav1.ManagedFieldsOperationApply, Time: nil},
				{Manager: "a", Operation: metav1.ManagedFieldsOperationApply, Time: parseTimeOrPanic("2001-01-01T01:00:00Z")},
			},
		},
		{
			name: "apply first",
			managedFields: []metav1.ManagedFieldsEntry{
				{Manager: "a", Operation: metav1.ManagedFieldsOperationUpdate, Time: parseTimeOrPanic("2001-01-01T01:00:00Z")},
				{Manager: "a", Operation: metav1.ManagedFieldsOperationApply, Time: nil},
			},
			expected: []metav1.ManagedFieldsEntry{
				{Manager: "a", Operation: metav1.ManagedFieldsOperationApply, Time: nil},
				{Manager: "a", Operation: metav1.ManagedFieldsOperationUpdate, Time: parseTimeOrPanic("2001-01-01T01:00:00Z")},
			},
		},
		{
			name: "newest last",
			managedFields: []metav1.ManagedFieldsEntry{
				{Manager: "c", Operation: metav1.ManagedFieldsOperationUpdate, Time: parseTimeOrPanic("2001-01-01T01:00:00Z")},
				{Manager: "a", Operation: metav1.ManagedFieldsOperationApply, Time: nil},
				{Manager: "c", Operation: metav1.ManagedFieldsOperationUpdate, Time: parseTimeOrPanic("2002-01-01T01:00:00Z")},
				{Manager: "b", Operation: metav1.ManagedFieldsOperationApply, Time: nil},
			},
			expected: []metav1.ManagedFieldsEntry{
				{Manager: "a", Operation: metav1.ManagedFieldsOperationApply, Time: nil},
				{Manager: "b", Operation: metav1.ManagedFieldsOperationApply, Time: nil},
				{Manager: "c", Operation: metav1.ManagedFieldsOperationUpdate, Time: parseTimeOrPanic("2001-01-01T01:00:00Z")},
				{Manager: "c", Operation: metav1.ManagedFieldsOperationUpdate, Time: parseTimeOrPanic("2002-01-01T01:00:00Z")},
			},
		},
		{
			name: "manager last",
			managedFields: []metav1.ManagedFieldsEntry{
				{Manager: "c", Operation: metav1.ManagedFieldsOperationUpdate, Time: parseTimeOrPanic("2001-01-01T01:00:00Z")},
				{Manager: "a", Operation: metav1.ManagedFieldsOperationApply, Time: nil},
				{Manager: "d", Operation: metav1.ManagedFieldsOperationUpdate, Time: parseTimeOrPanic("2001-01-01T01:00:00Z")},
				{Manager: "b", Operation: metav1.ManagedFieldsOperationApply, Time: nil},
			},
			expected: []metav1.ManagedFieldsEntry{
				{Manager: "a", Operation: metav1.ManagedFieldsOperationApply, Time: nil},
				{Manager: "b", Operation: metav1.ManagedFieldsOperationApply, Time: nil},
				{Manager: "c", Operation: metav1.ManagedFieldsOperationUpdate, Time: parseTimeOrPanic("2001-01-01T01:00:00Z")},
				{Manager: "d", Operation: metav1.ManagedFieldsOperationUpdate, Time: parseTimeOrPanic("2001-01-01T01:00:00Z")},
			},
		},
		{
			name: "manager sorted",
			managedFields: []metav1.ManagedFieldsEntry{
				{Manager: "c", Operation: metav1.ManagedFieldsOperationUpdate, Time: parseTimeOrPanic("2001-01-01T01:00:00Z")},
				{Manager: "a", Operation: metav1.ManagedFieldsOperationApply, Time: nil},
				{Manager: "g", Operation: metav1.ManagedFieldsOperationApply, Time: nil},
				{Manager: "f", Operation: metav1.ManagedFieldsOperationUpdate, Time: parseTimeOrPanic("2002-01-01T01:00:00Z")},
				{Manager: "i", Operation: metav1.ManagedFieldsOperationApply, Time: nil},
				{Manager: "d", Operation: metav1.ManagedFieldsOperationUpdate, Time: parseTimeOrPanic("2002-01-01T01:00:00Z")},
				{Manager: "h", Operation: metav1.ManagedFieldsOperationApply, Time: nil},
				{Manager: "e", Operation: metav1.ManagedFieldsOperationUpdate, Time: parseTimeOrPanic("2003-01-01T01:00:00Z")},
				{Manager: "b", Operation: metav1.ManagedFieldsOperationApply, Time: nil},
			},
			expected: []metav1.ManagedFieldsEntry{
				{Manager: "a", Operation: metav1.ManagedFieldsOperationApply, Time: nil},
				{Manager: "b", Operation: metav1.ManagedFieldsOperationApply, Time: nil},
				{Manager: "g", Operation: metav1.ManagedFieldsOperationApply, Time: nil},
				{Manager: "h", Operation: metav1.ManagedFieldsOperationApply, Time: nil},
				{Manager: "i", Operation: metav1.ManagedFieldsOperationApply, Time: nil},
				{Manager: "c", Operation: metav1.ManagedFieldsOperationUpdate, Time: parseTimeOrPanic("2001-01-01T01:00:00Z")},
				{Manager: "d", Operation: metav1.ManagedFieldsOperationUpdate, Time: parseTimeOrPanic("2002-01-01T01:00:00Z")},
				{Manager: "f", Operation: metav1.ManagedFieldsOperationUpdate, Time: parseTimeOrPanic("2002-01-01T01:00:00Z")},
				{Manager: "e", Operation: metav1.ManagedFieldsOperationUpdate, Time: parseTimeOrPanic("2003-01-01T01:00:00Z")},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			sorted, err := sortEncodedManagedFields(test.managedFields)
			if err != nil {
				t.Fatalf("did not expect error when sorting but got: %v", err)
			}
			if !reflect.DeepEqual(sorted, test.expected) {
				t.Fatalf("expected:\n%v\nbut got:\n%v", test.expected, sorted)
			}
		})
	}
}

func parseTimeOrPanic(s string) *metav1.Time {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		panic(fmt.Sprintf("failed to parse time %s, got: %v", s, err))
	}
	return &metav1.Time{Time: t.UTC()}
}
