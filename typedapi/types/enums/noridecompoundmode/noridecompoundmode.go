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


// Package noridecompoundmode
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/_types/analysis/tokenizers.ts#L74-L78
package noridecompoundmode

import "strings"

type NoriDecompoundMode struct {
	name string
}

var (
	Discard = NoriDecompoundMode{"discard"}

	None = NoriDecompoundMode{"none"}

	Mixed = NoriDecompoundMode{"mixed"}
)

func (n NoriDecompoundMode) MarshalText() (text []byte, err error) {
	return []byte(n.String()), nil
}

func (n *NoriDecompoundMode) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {

	case "discard":
		*n = Discard
	case "none":
		*n = None
	case "mixed":
		*n = Mixed
	default:
		*n = NoriDecompoundMode{string(text)}
	}

	return nil
}

func (n NoriDecompoundMode) String() string {
	return n.name
}