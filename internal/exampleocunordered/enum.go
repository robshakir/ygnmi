// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
Package exampleocunordered is a generated package which contains definitions
of structs which represent a YANG schema. The generated schema can be
compressed by a series of transformations (compression was true
in this case).

This package was generated by ygnmi version: (devel): (ygot: v0.29.6)
using the following YANG input files:
  - ../../pathgen/testdata/yang/openconfig-simple.yang
  - ../../pathgen/testdata/yang/openconfig-withlistval.yang
  - ../../pathgen/testdata/yang/openconfig-nested.yang

Imported modules were sourced from:
*/
package exampleocunordered

import (
	"github.com/openconfig/ygot/ygot"
)

// E_Child_Three is a derived int64 type which is used to represent
// the enumerated node Child_Three. An additional value named
// Child_Three_UNSET is added to the enumeration which is used as
// the nil value, indicating that the enumeration was not explicitly set by
// the program importing the generated structures.
type E_Child_Three int64

// IsYANGGoEnum ensures that Child_Three implements the yang.GoEnum
// interface. This ensures that Child_Three can be identified as a
// mapped type for a YANG enumeration.
func (E_Child_Three) IsYANGGoEnum() {}

// ΛMap returns the value lookup map associated with  Child_Three.
func (E_Child_Three) ΛMap() map[string]map[int64]ygot.EnumDefinition { return ΛEnum }

// String returns a logging-friendly string for E_Child_Three.
func (e E_Child_Three) String() string {
	return ygot.EnumLogString(e, int64(e), "E_Child_Three")
}

const (
	// Child_Three_UNSET corresponds to the value UNSET of Child_Three
	Child_Three_UNSET E_Child_Three = 0
	// Child_Three_ONE corresponds to the value ONE of Child_Three
	Child_Three_ONE E_Child_Three = 1
	// Child_Three_TWO corresponds to the value TWO of Child_Three
	Child_Three_TWO E_Child_Three = 2
)
