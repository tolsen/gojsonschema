// Copyright 2013 sigu-399 ( https://github.com/sigu-399 )
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// author           sigu-399
// author-github    https://github.com/sigu-399
// author-mail      sigu.399@gmail.com
//
// repository-name  gojsonschema
// repository-desc  An implementation of JSON Schema, based on IETF's draft v4 - Go language.
//
// description      Pool of referenced schemas.
//
// created          25-06-2013

package gojsonschema

import ()

type schemaReferencePool struct {
	schemaPoolDocuments map[string]*jsonSchema
}

func newSchemaReferencePool() *schemaReferencePool {
	p := &schemaReferencePool{}
	p.schemaPoolDocuments = make(map[string]*jsonSchema)
	return p
}

func (p *schemaReferencePool) GetSchema(ref string) (r *jsonSchema, o bool) {

	if sch, ok := p.schemaPoolDocuments[ref]; ok {
		return sch, true
	}
	return nil, false
}

func (p *schemaReferencePool) AddSchema(ref string, sch *jsonSchema) {
	p.schemaPoolDocuments[ref] = sch
}
