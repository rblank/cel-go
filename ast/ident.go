// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ast

import "celgo/common"

type IdentExpression struct {
	BaseExpression

	Name string
}

func (e *IdentExpression) String() string {
	return ToDebugString(e)
}

func (e *IdentExpression) writeDebugString(w *debugWriter) {
	w.append(e.Name)
	w.adorn(e)
}

func NewIdent(id int64, l common.Location, name string) *IdentExpression {
	return &IdentExpression{
		BaseExpression: BaseExpression{id: id, location: l},
		Name:           name,
	}
}
