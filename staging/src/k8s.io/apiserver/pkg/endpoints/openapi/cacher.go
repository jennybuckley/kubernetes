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
	"fmt"
	"net/http"
	"sync"

	golangproto "github.com/golang/protobuf/proto"
	openapi_v2 "github.com/googleapis/gnostic/OpenAPIv2"

	"k8s.io/apimachinery/pkg/runtime/schema"
	utilresponsewriter "k8s.io/apiserver/pkg/util/responsewriter"
	"k8s.io/kube-openapi/pkg/util/proto"
)

// SchemaCacher fetches the openapi schema once and then caches it in memory
type SchemaCacher struct {
	rwMutex sync.RWMutex

	openAPIHandler http.Handler
	openAPIPath    string

	openAPISchema Resources
	lastEtag      string
}

var _ Getter = &SchemaCacher{}

// SetSource sets the source for the OpenAPI schema
func (s *SchemaCacher) SetSource(handler http.Handler, path string) {
	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()
	s.openAPIHandler = handler
	s.openAPIPath = path
}

// Get implements Getter
func (s *SchemaCacher) Get() (Resources, error) {
	err := s.updateSchema()
	if err != nil {
		return nil, err
	}

	s.rwMutex.RLock()
	defer s.rwMutex.RUnlock()
	return s.openAPISchema, nil
}

// updateSchema will parse the OpenAPI schema if it has changed since the last
// time updateSchema was called, and cache it.
func (s *SchemaCacher) updateSchema() error {
	schemaDoc, newEtag, httpStatus, err := s.downloadSchema()
	if err != nil {
		return err
	}

	// The schema hasn't changed so there's nothing to update
	if httpStatus == http.StatusNotModified {
		return nil
	}

	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()

	// Check etag again after taking the write lock in case lastEtag
	// was updated while waiting for the write lock
	if newEtag == s.lastEtag {
		return nil
	}

	// Convert openapi_v2.Document to our Resource type
	schema, err := NewOpenAPIData(schemaDoc)
	if err != nil {
		return err
	}

	s.openAPISchema = schema
	s.lastEtag = newEtag
	return nil
}

// downloadSchema will attempt to get the OpenAPISchema from the handler and path
// that this SchemaCacher uses as a source, if it has changed.
func (s *SchemaCacher) downloadSchema() (returnSchemaDoc *openapi_v2.Document, newEtag string, httpStatus int, err error) {
	s.rwMutex.RLock()
	defer s.rwMutex.RUnlock()

	if s.openAPIHandler == nil {
		return nil, "", 0, fmt.Errorf("No OpenAPI handler specified")
	}

	req, err := http.NewRequest("GET", s.openAPIPath, nil)
	if err != nil {
		return nil, "", 0, err
	}
	req.Header.Add("Accept", "application/com.github.proto-openapi.spec.v2@v1.0+protobuf")
	if len(s.lastEtag) > 0 {
		req.Header.Add("If-None-Match", s.lastEtag)
	}

	writer := utilresponsewriter.NewInMemoryResponseWriter()
	s.openAPIHandler.ServeHTTP(writer, req)

	switch writer.RespCode {
	case http.StatusNotModified:
		if len(s.lastEtag) == 0 {
			return nil, s.lastEtag, http.StatusNotModified, fmt.Errorf("http.StatusNotModified is not allowed in absence of etag")
		}
		return nil, s.lastEtag, http.StatusNotModified, nil
	case http.StatusOK:
		returnSchemaDoc = &openapi_v2.Document{}
		if err := golangproto.Unmarshal(writer.Data, returnSchemaDoc); err != nil {
			return nil, "", 0, err
		}
		newEtag = writer.Header().Get("Etag")
		return returnSchemaDoc, newEtag, http.StatusOK, nil
	default:
		return nil, "", 0, fmt.Errorf("failed to retrieve openAPI spec, http error: %s", writer.String())
	}
}

// groupVersionKindExtensionKey is the key used to lookup the
// GroupVersionKind value for an object definition from the
// definition's "extensions" map.
const groupVersionKindExtensionKey = "x-kubernetes-group-version-kind"

// document is an implementation of `Resources`. It looks for
// resources in an openapi Schema.
type document struct {
	// Maps gvk to model name
	resources map[schema.GroupVersionKind]string
	models    proto.Models
}

var _ Resources = &document{}

func NewOpenAPIData(doc *openapi_v2.Document) (Resources, error) {
	models, err := proto.NewOpenAPIData(doc)
	if err != nil {
		return nil, err
	}

	resources := map[schema.GroupVersionKind]string{}
	for _, modelName := range models.ListModels() {
		model := models.LookupModel(modelName)
		if model == nil {
			panic("ListModels returns a model that can't be looked-up.")
		}
		gvkList := parseGroupVersionKind(model)
		for _, gvk := range gvkList {
			if len(gvk.Kind) > 0 {
				resources[gvk] = modelName
			}
		}
	}

	return &document{
		resources: resources,
		models:    models,
	}, nil
}

func (d *document) LookupResource(gvk schema.GroupVersionKind) proto.Schema {
	modelName, found := d.resources[gvk]
	if !found {
		return nil
	}
	return d.models.LookupModel(modelName)
}

// Get and parse GroupVersionKind from the extension. Returns empty if it doesn't have one.
func parseGroupVersionKind(s proto.Schema) []schema.GroupVersionKind {
	extensions := s.GetExtensions()

	gvkListResult := []schema.GroupVersionKind{}

	// Get the extensions
	gvkExtension, ok := extensions[groupVersionKindExtensionKey]
	if !ok {
		return []schema.GroupVersionKind{}
	}

	// gvk extension must be a list of at least 1 element.
	gvkList, ok := gvkExtension.([]interface{})
	if !ok {
		return []schema.GroupVersionKind{}
	}

	for _, gvk := range gvkList {
		// gvk extension list must be a map with group, version, and
		// kind fields
		gvkMap, ok := gvk.(map[interface{}]interface{})
		if !ok {
			continue
		}
		group, ok := gvkMap["group"].(string)
		if !ok {
			continue
		}
		version, ok := gvkMap["version"].(string)
		if !ok {
			continue
		}
		kind, ok := gvkMap["kind"].(string)
		if !ok {
			continue
		}

		gvkListResult = append(gvkListResult, schema.GroupVersionKind{
			Group:   group,
			Version: version,
			Kind:    kind,
		})
	}

	return gvkListResult
}
