// Copyright 2021 Google LLC
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

package mk2rbc

// Starlark expression types we use
type starlarkType int

const (
	starlarkTypeUnknown starlarkType = iota
	starlarkTypeList    starlarkType = iota
	starlarkTypeString  starlarkType = iota
	starlarkTypeInt     starlarkType = iota
	starlarkTypeBool    starlarkType = iota
	starlarkTypeVoid    starlarkType = iota
)

type varClass int

const (
	VarClassConfig varClass = iota
	VarClassSoong  varClass = iota
	VarClassLocal  varClass = iota
)

type variableRegistrar interface {
	NewVariable(name string, varClass varClass, valueType starlarkType)
}

// ScopeBase is a dummy implementation of the mkparser.Scope.
// All our scopes are read-only and resolve only simple variables.
type ScopeBase struct{}

func (s ScopeBase) Set(_, _ string) {
	panic("implement me")
}

func (s ScopeBase) Call(_ string, _ []string) []string {
	panic("implement me")
}

func (s ScopeBase) SetFunc(_ string, _ func([]string) []string) {
	panic("implement me")
}