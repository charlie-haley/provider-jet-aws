/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	"github.com/pkg/errors"

	"github.com/crossplane/terrajet/pkg/resource"
	"github.com/crossplane/terrajet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this Connection
func (mg *Connection) GetTerraformResourceType() string {
	return "aws_cloudwatch_event_connection"
}

// GetConnectionDetailsMapping for this Connection
func (tr *Connection) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"auth_parameters[*].api_key[*].value": "spec.forProvider.authParameters[*].apiKey[*].valueSecretRef", "auth_parameters[*].basic[*].password": "spec.forProvider.authParameters[*].basic[*].passwordSecretRef", "auth_parameters[*].invocation_http_parameters[*].body[*].value": "spec.forProvider.authParameters[*].invocationHTTPParameters[*].body[*].valueSecretRef", "auth_parameters[*].invocation_http_parameters[*].header[*].value": "spec.forProvider.authParameters[*].invocationHTTPParameters[*].header[*].valueSecretRef", "auth_parameters[*].invocation_http_parameters[*].query_string[*].value": "spec.forProvider.authParameters[*].invocationHTTPParameters[*].queryString[*].valueSecretRef", "auth_parameters[*].oauth[*].client_parameters[*].client_secret": "spec.forProvider.authParameters[*].oauth[*].clientParameters[*].clientSecretSecretRef", "auth_parameters[*].oauth[*].oauth_http_parameters[*].body[*].value": "spec.forProvider.authParameters[*].oauth[*].oauthHTTPParameters[*].body[*].valueSecretRef", "auth_parameters[*].oauth[*].oauth_http_parameters[*].header[*].value": "spec.forProvider.authParameters[*].oauth[*].oauthHTTPParameters[*].header[*].valueSecretRef", "auth_parameters[*].oauth[*].oauth_http_parameters[*].query_string[*].value": "spec.forProvider.authParameters[*].oauth[*].oauthHTTPParameters[*].queryString[*].valueSecretRef"}
}

// GetObservation of this Connection
func (tr *Connection) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this Connection
func (tr *Connection) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this Connection
func (tr *Connection) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this Connection
func (tr *Connection) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this Connection
func (tr *Connection) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this Connection using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *Connection) LateInitialize(attrs []byte) (bool, error) {
	params := &ConnectionParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *Connection) GetTerraformSchemaVersion() int {
	return 0
}
