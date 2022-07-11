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

// CustomCategorizeTextAnalyzer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/_types/aggregations/bucket.ts#L497-L501
type CustomCategorizeTextAnalyzer struct {
	CharFilter []string `json:"char_filter,omitempty"`
	Filter     []string `json:"filter,omitempty"`
	Tokenizer  *string  `json:"tokenizer,omitempty"`
}

// CustomCategorizeTextAnalyzerBuilder holds CustomCategorizeTextAnalyzer struct and provides a builder API.
type CustomCategorizeTextAnalyzerBuilder struct {
	v *CustomCategorizeTextAnalyzer
}

// NewCustomCategorizeTextAnalyzer provides a builder for the CustomCategorizeTextAnalyzer struct.
func NewCustomCategorizeTextAnalyzerBuilder() *CustomCategorizeTextAnalyzerBuilder {
	r := CustomCategorizeTextAnalyzerBuilder{
		&CustomCategorizeTextAnalyzer{},
	}

	return &r
}

// Build finalize the chain and returns the CustomCategorizeTextAnalyzer struct
func (rb *CustomCategorizeTextAnalyzerBuilder) Build() CustomCategorizeTextAnalyzer {
	return *rb.v
}

func (rb *CustomCategorizeTextAnalyzerBuilder) CharFilter(char_filter ...string) *CustomCategorizeTextAnalyzerBuilder {
	rb.v.CharFilter = char_filter
	return rb
}

func (rb *CustomCategorizeTextAnalyzerBuilder) Filter(filter ...string) *CustomCategorizeTextAnalyzerBuilder {
	rb.v.Filter = filter
	return rb
}

func (rb *CustomCategorizeTextAnalyzerBuilder) Tokenizer(tokenizer string) *CustomCategorizeTextAnalyzerBuilder {
	rb.v.Tokenizer = &tokenizer
	return rb
}