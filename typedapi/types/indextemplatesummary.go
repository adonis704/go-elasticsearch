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
// https://github.com/elastic/elasticsearch-specification/tree/76e25d34bff1060e300c95f4be468ef88e4f3465

package types

// IndexTemplateSummary type.
//
// https://github.com/elastic/elasticsearch-specification/blob/76e25d34bff1060e300c95f4be468ef88e4f3465/specification/indices/_types/IndexTemplate.ts#L56-L65
type IndexTemplateSummary struct {
	Aliases   map[string]Alias           `json:"aliases,omitempty"`
	Lifecycle *DataLifecycleWithRollover `json:"lifecycle,omitempty"`
	Mappings  *TypeMapping               `json:"mappings,omitempty"`
	Settings  *IndexSettings             `json:"settings,omitempty"`
}

// NewIndexTemplateSummary returns a IndexTemplateSummary.
func NewIndexTemplateSummary() *IndexTemplateSummary {
	r := &IndexTemplateSummary{
		Aliases: make(map[string]Alias, 0),
	}

	return r
}
