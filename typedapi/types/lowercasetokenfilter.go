// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.


// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/135ae054e304239743b5777ad8d41cb2c9091d35


package types

// LowercaseTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/_types/analysis/token_filters.ts#L253-L256
type LowercaseTokenFilter struct {
	Language *string        `json:"language,omitempty"`
	Type     string         `json:"type,omitempty"`
	Version  *VersionString `json:"version,omitempty"`
}

// LowercaseTokenFilterBuilder holds LowercaseTokenFilter struct and provides a builder API.
type LowercaseTokenFilterBuilder struct {
	v *LowercaseTokenFilter
}

// NewLowercaseTokenFilter provides a builder for the LowercaseTokenFilter struct.
func NewLowercaseTokenFilterBuilder() *LowercaseTokenFilterBuilder {
	r := LowercaseTokenFilterBuilder{
		&LowercaseTokenFilter{},
	}

	r.v.Type = "lowercase"

	return &r
}

// Build finalize the chain and returns the LowercaseTokenFilter struct
func (rb *LowercaseTokenFilterBuilder) Build() LowercaseTokenFilter {
	return *rb.v
}

func (rb *LowercaseTokenFilterBuilder) Language(language string) *LowercaseTokenFilterBuilder {
	rb.v.Language = &language
	return rb
}

func (rb *LowercaseTokenFilterBuilder) Version(version VersionString) *LowercaseTokenFilterBuilder {
	rb.v.Version = &version
	return rb
}