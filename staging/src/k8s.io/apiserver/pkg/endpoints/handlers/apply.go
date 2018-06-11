/*
Copyright 2017 The Kubernetes Authors.

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

package handlers

import (
	"fmt"

	"github.com/ghodss/yaml"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/apply"
	"k8s.io/apimachinery/pkg/apply/parse"
	"k8s.io/apimachinery/pkg/apply/strategy"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/json"
	"k8s.io/kube-openapi/pkg/util/proto"
)

type applyPatcher struct {
	*patcher

	model proto.Schema
}

// TODO(apelisse): workflowId needs to be passed as a query
// param/header, and a better defaulting needs to be defined too.
const workflowId = "default"

func (p *applyPatcher) convertCurrentVersion(obj runtime.Object) (map[string]interface{}, error) {
	vo, err := p.unsafeConvertor.ConvertToVersion(obj, p.kind.GroupVersion())
	if err != nil {
		return nil, err
	}
	return runtime.DefaultUnstructuredConverter.ToUnstructured(vo)
}

func (p *applyPatcher) extractLastIntent(obj runtime.Object, workflow string) (map[string]interface{}, error) {
	accessor, err := meta.Accessor(obj)
	if err != nil {
		return nil, fmt.Errorf("couldn't get accessor: %v", err)
	}
	last := make(map[string]interface{})
	if accessor.GetLastApplied()[workflow] != "" {
		if err := json.Unmarshal([]byte(accessor.GetLastApplied()[workflow]), &last); err != nil {
			return nil, fmt.Errorf("couldn't unmarshal last applied field: %v", err)
		}
	}
	return last, nil
}

func (p *applyPatcher) getNewIntent() (map[string]interface{}, error) {
	patch := make(map[string]interface{})
	if err := yaml.Unmarshal(p.patchBytes, &patch); err != nil {
		return nil, fmt.Errorf("couldn't unmarshal patch object: %v (patch: %v)", err, string(p.patchBytes))
	}
	return patch, nil
}

func (p *applyPatcher) convertResultToUnversioned(result apply.Result) (runtime.Object, error) {
	voutput, err := p.creater.New(p.kind)
	if err != nil {
		return nil, fmt.Errorf("failed to create empty output object: %v", err)
	}

	err = runtime.DefaultUnstructuredConverter.FromUnstructured(result.MergedResult.(map[string]interface{}), voutput)
	if err != nil {
		return nil, fmt.Errorf("failed to convert merge result back: %v", err)
	}
	p.defaulter.Default(voutput)

	uoutput, err := p.toUnversioned(voutput)
	if err != nil {
		return nil, fmt.Errorf("failed to convert to unversioned: %v", err)
	}

	return uoutput, nil
}

func (p *applyPatcher) saveNewIntent(workflow string, dst runtime.Object) error {
	// We want to consistently save the intent in JSon.
	json, err := yaml.YAMLToJSON(p.patchBytes)
	if err != nil {
		return fmt.Errorf("failed to convert patch to JSon: %v", err)
	}
	accessor, err := meta.Accessor(dst)
	if err != nil {
		return fmt.Errorf("couldn't get accessor: %v", err)
	}
	m := accessor.GetLastApplied()
	if m == nil {
		m = make(map[string]string)
	}
	m[workflow] = string(json)
	accessor.SetLastApplied(m)
	return nil
}

func (p *applyPatcher) applyPatchToCurrentObject(currentObject runtime.Object) (runtime.Object, error) {
	current, err := p.convertCurrentVersion(currentObject)
	if err != nil {
		return nil, fmt.Errorf("failed to convert current object: %v", err)
	}

	lastIntent, err := p.extractLastIntent(currentObject, workflowId)
	if err != nil {
		return nil, fmt.Errorf("failed to extract last intent: %v", err)
	}
	newIntent, err := p.getNewIntent()
	if err != nil {
		return nil, fmt.Errorf("failed to get new intent: %v", err)
	}

	element, err := parse.CreateElement(lastIntent, newIntent, current, p.model)
	if err != nil {
		return nil, fmt.Errorf("failed to parse elements: %v", err)
	}
	result, err := element.Merge(strategy.Create(strategy.Options{}))
	if err != nil {
		return nil, fmt.Errorf("failed to merge elements: %v", err)
	}

	output, err := p.convertResultToUnversioned(result)
	if err != nil {
		return nil, fmt.Errorf("failed to convert merge result: %v", err)
	}

	if err := p.saveNewIntent(workflowId, output); err != nil {
		return nil, fmt.Errorf("failed to save last intent: %v", err)
	}

	// TODO(apelisse): Also update last-applied on the create path
	// TODO(apelisse): Check for conflicts with other lastApplied
	// and report actionable errors to users.

	return output, nil
}

func (p *applyPatcher) createNewObject() (runtime.Object, error) {
	original := p.restPatcher.New()
	objToCreate, gvk, err := p.codec.Decode(p.patchBytes, &p.kind, original)
	if err != nil {
		return nil, transformDecodeError(p.typer, err, original, gvk, p.patchBytes)
	}
	if gvk.GroupVersion() != p.kind.GroupVersion() {
		return nil, errors.NewBadRequest(fmt.Sprintf("the API version in the data (%s) does not match the expected API version (%v)", gvk.GroupVersion().String(), p.kind.GroupVersion().String()))
	}
	return objToCreate, err
}
