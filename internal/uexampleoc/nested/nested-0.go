// Copyright 2023 Google LLC
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
Package nested is a generated package which contains definitions
of structs which generate gNMI paths for a YANG schema.

This package was generated by ygnmi version: (devel): (ygot: (devel))
using the following YANG input files:
  - ../../pathgen/testdata/yang/openconfig-simple.yang
  - ../../pathgen/testdata/yang/openconfig-withlistval.yang
  - ../../pathgen/testdata/yang/openconfig-nested.yang

Imported modules were sourced from:
*/
package nested

import (
	oc "github.com/openconfig/ygnmi/internal/uexampleoc"
	"github.com/openconfig/ygnmi/ygnmi"
	"github.com/openconfig/ygot/ygot"
	"github.com/openconfig/ygot/ytypes"
)

// OpenconfigNested_APath represents the /openconfig-nested/a YANG schema element.
type OpenconfigNested_APath struct {
	*ygnmi.NodePath
}

// OpenconfigNested_APathAny represents the wildcard version of the /openconfig-nested/a YANG schema element.
type OpenconfigNested_APathAny struct {
	*ygnmi.NodePath
}

// B (container):
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "b"
//	Path from root:       "/a/b"
func (n *OpenconfigNested_APath) B() *OpenconfigNested_A_BPath {
	return &OpenconfigNested_A_BPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"b"},
			map[string]interface{}{},
			n,
		),
	}
}

// B (container):
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "b"
//	Path from root:       "/a/b"
func (n *OpenconfigNested_APathAny) B() *OpenconfigNested_A_BPathAny {
	return &OpenconfigNested_A_BPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"b"},
			map[string]interface{}{},
			n,
		),
	}
}

func binarySliceToFloatSlice(in []oc.Binary) []float32 {
	converted := make([]float32, 0, len(in))
	for _, binary := range in {
		converted = append(converted, ygot.BinaryToFloat32(binary))
	}
	return converted
}

// Query returns a Query that can be used in gNMI operations.
func (n *OpenconfigNested_APath) Query() ygnmi.SingletonQuery[*oc.OpenconfigNested_A] {
	return ygnmi.NewSingletonQuery[*oc.OpenconfigNested_A](
		"OpenconfigNested_A",
		true,
		false,
		false,
		false,
		n,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)
}

// Query returns a Query that can be used in gNMI operations.
func (n *OpenconfigNested_APathAny) Query() ygnmi.WildcardQuery[*oc.OpenconfigNested_A] {
	return ygnmi.NewWildcardQuery[*oc.OpenconfigNested_A](
		"OpenconfigNested_A",
		true,
		false,
		false,
		false,
		n,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
	)
}

// OpenconfigNested_A_BPath represents the /openconfig-nested/a/b YANG schema element.
type OpenconfigNested_A_BPath struct {
	*ygnmi.NodePath
}

// OpenconfigNested_A_BPathAny represents the wildcard version of the /openconfig-nested/a/b YANG schema element.
type OpenconfigNested_A_BPathAny struct {
	*ygnmi.NodePath
}

// C (container):
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "c"
//	Path from root:       "/a/b/c"
func (n *OpenconfigNested_A_BPath) C() *OpenconfigNested_A_B_CPath {
	return &OpenconfigNested_A_B_CPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"c"},
			map[string]interface{}{},
			n,
		),
	}
}

// C (container):
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "c"
//	Path from root:       "/a/b/c"
func (n *OpenconfigNested_A_BPathAny) C() *OpenconfigNested_A_B_CPathAny {
	return &OpenconfigNested_A_B_CPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"c"},
			map[string]interface{}{},
			n,
		),
	}
}

// Query returns a Query that can be used in gNMI operations.
func (n *OpenconfigNested_A_BPath) Query() ygnmi.SingletonQuery[*oc.OpenconfigNested_A_B] {
	return ygnmi.NewSingletonQuery[*oc.OpenconfigNested_A_B](
		"OpenconfigNested_A_B",
		true,
		false,
		false,
		false,
		n,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)
}

// Query returns a Query that can be used in gNMI operations.
func (n *OpenconfigNested_A_BPathAny) Query() ygnmi.WildcardQuery[*oc.OpenconfigNested_A_B] {
	return ygnmi.NewWildcardQuery[*oc.OpenconfigNested_A_B](
		"OpenconfigNested_A_B",
		true,
		false,
		false,
		false,
		n,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
	)
}

// OpenconfigNested_A_B_CPath represents the /openconfig-nested/a/b/c YANG schema element.
type OpenconfigNested_A_B_CPath struct {
	*ygnmi.NodePath
}

// OpenconfigNested_A_B_CPathAny represents the wildcard version of the /openconfig-nested/a/b/c YANG schema element.
type OpenconfigNested_A_B_CPathAny struct {
	*ygnmi.NodePath
}

// D (container):
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "d"
//	Path from root:       "/a/b/c/d"
func (n *OpenconfigNested_A_B_CPath) D() *OpenconfigNested_A_B_C_DPath {
	return &OpenconfigNested_A_B_C_DPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"d"},
			map[string]interface{}{},
			n,
		),
	}
}

// D (container):
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "d"
//	Path from root:       "/a/b/c/d"
func (n *OpenconfigNested_A_B_CPathAny) D() *OpenconfigNested_A_B_C_DPathAny {
	return &OpenconfigNested_A_B_C_DPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"d"},
			map[string]interface{}{},
			n,
		),
	}
}

// Query returns a Query that can be used in gNMI operations.
func (n *OpenconfigNested_A_B_CPath) Query() ygnmi.SingletonQuery[*oc.OpenconfigNested_A_B_C] {
	return ygnmi.NewSingletonQuery[*oc.OpenconfigNested_A_B_C](
		"OpenconfigNested_A_B_C",
		true,
		false,
		false,
		false,
		n,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)
}

// Query returns a Query that can be used in gNMI operations.
func (n *OpenconfigNested_A_B_CPathAny) Query() ygnmi.WildcardQuery[*oc.OpenconfigNested_A_B_C] {
	return ygnmi.NewWildcardQuery[*oc.OpenconfigNested_A_B_C](
		"OpenconfigNested_A_B_C",
		true,
		false,
		false,
		false,
		n,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
	)
}

// OpenconfigNested_A_B_C_DPath represents the /openconfig-nested/a/b/c/d YANG schema element.
type OpenconfigNested_A_B_C_DPath struct {
	*ygnmi.NodePath
}

// OpenconfigNested_A_B_C_DPathAny represents the wildcard version of the /openconfig-nested/a/b/c/d YANG schema element.
type OpenconfigNested_A_B_C_DPathAny struct {
	*ygnmi.NodePath
}

// E (container):
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "e"
//	Path from root:       "/a/b/c/d/e"
func (n *OpenconfigNested_A_B_C_DPath) E() *OpenconfigNested_A_B_C_D_EPath {
	return &OpenconfigNested_A_B_C_D_EPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"e"},
			map[string]interface{}{},
			n,
		),
	}
}

// E (container):
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "e"
//	Path from root:       "/a/b/c/d/e"
func (n *OpenconfigNested_A_B_C_DPathAny) E() *OpenconfigNested_A_B_C_D_EPathAny {
	return &OpenconfigNested_A_B_C_D_EPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"e"},
			map[string]interface{}{},
			n,
		),
	}
}

// Query returns a Query that can be used in gNMI operations.
func (n *OpenconfigNested_A_B_C_DPath) Query() ygnmi.SingletonQuery[*oc.OpenconfigNested_A_B_C_D] {
	return ygnmi.NewSingletonQuery[*oc.OpenconfigNested_A_B_C_D](
		"OpenconfigNested_A_B_C_D",
		true,
		false,
		false,
		false,
		n,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)
}

// Query returns a Query that can be used in gNMI operations.
func (n *OpenconfigNested_A_B_C_DPathAny) Query() ygnmi.WildcardQuery[*oc.OpenconfigNested_A_B_C_D] {
	return ygnmi.NewWildcardQuery[*oc.OpenconfigNested_A_B_C_D](
		"OpenconfigNested_A_B_C_D",
		true,
		false,
		false,
		false,
		n,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
	)
}

// OpenconfigNested_A_B_C_D_EPath represents the /openconfig-nested/a/b/c/d/e YANG schema element.
type OpenconfigNested_A_B_C_D_EPath struct {
	*ygnmi.NodePath
}

// OpenconfigNested_A_B_C_D_EPathAny represents the wildcard version of the /openconfig-nested/a/b/c/d/e YANG schema element.
type OpenconfigNested_A_B_C_D_EPathAny struct {
	*ygnmi.NodePath
}

// F (container):
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "f"
//	Path from root:       "/a/b/c/d/e/f"
func (n *OpenconfigNested_A_B_C_D_EPath) F() *OpenconfigNested_A_B_C_D_E_FPath {
	return &OpenconfigNested_A_B_C_D_E_FPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"f"},
			map[string]interface{}{},
			n,
		),
	}
}

// F (container):
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "f"
//	Path from root:       "/a/b/c/d/e/f"
func (n *OpenconfigNested_A_B_C_D_EPathAny) F() *OpenconfigNested_A_B_C_D_E_FPathAny {
	return &OpenconfigNested_A_B_C_D_E_FPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"f"},
			map[string]interface{}{},
			n,
		),
	}
}

// Query returns a Query that can be used in gNMI operations.
func (n *OpenconfigNested_A_B_C_D_EPath) Query() ygnmi.SingletonQuery[*oc.OpenconfigNested_A_B_C_D_E] {
	return ygnmi.NewSingletonQuery[*oc.OpenconfigNested_A_B_C_D_E](
		"OpenconfigNested_A_B_C_D_E",
		true,
		false,
		false,
		false,
		n,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)
}

// Query returns a Query that can be used in gNMI operations.
func (n *OpenconfigNested_A_B_C_D_EPathAny) Query() ygnmi.WildcardQuery[*oc.OpenconfigNested_A_B_C_D_E] {
	return ygnmi.NewWildcardQuery[*oc.OpenconfigNested_A_B_C_D_E](
		"OpenconfigNested_A_B_C_D_E",
		true,
		false,
		false,
		false,
		n,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
	)
}

// OpenconfigNested_A_B_C_D_E_FPath represents the /openconfig-nested/a/b/c/d/e/f YANG schema element.
type OpenconfigNested_A_B_C_D_E_FPath struct {
	*ygnmi.NodePath
}

// OpenconfigNested_A_B_C_D_E_FPathAny represents the wildcard version of the /openconfig-nested/a/b/c/d/e/f YANG schema element.
type OpenconfigNested_A_B_C_D_E_FPathAny struct {
	*ygnmi.NodePath
}

// G (container):
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "g"
//	Path from root:       "/a/b/c/d/e/f/g"
func (n *OpenconfigNested_A_B_C_D_E_FPath) G() *OpenconfigNested_A_B_C_D_E_F_GPath {
	return &OpenconfigNested_A_B_C_D_E_F_GPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"g"},
			map[string]interface{}{},
			n,
		),
	}
}

// G (container):
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "g"
//	Path from root:       "/a/b/c/d/e/f/g"
func (n *OpenconfigNested_A_B_C_D_E_FPathAny) G() *OpenconfigNested_A_B_C_D_E_F_GPathAny {
	return &OpenconfigNested_A_B_C_D_E_F_GPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"g"},
			map[string]interface{}{},
			n,
		),
	}
}

// Query returns a Query that can be used in gNMI operations.
func (n *OpenconfigNested_A_B_C_D_E_FPath) Query() ygnmi.SingletonQuery[*oc.OpenconfigNested_A_B_C_D_E_F] {
	return ygnmi.NewSingletonQuery[*oc.OpenconfigNested_A_B_C_D_E_F](
		"OpenconfigNested_A_B_C_D_E_F",
		true,
		false,
		false,
		false,
		n,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)
}

// Query returns a Query that can be used in gNMI operations.
func (n *OpenconfigNested_A_B_C_D_E_FPathAny) Query() ygnmi.WildcardQuery[*oc.OpenconfigNested_A_B_C_D_E_F] {
	return ygnmi.NewWildcardQuery[*oc.OpenconfigNested_A_B_C_D_E_F](
		"OpenconfigNested_A_B_C_D_E_F",
		true,
		false,
		false,
		false,
		n,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
	)
}

// OpenconfigNested_A_B_C_D_E_F_GPath represents the /openconfig-nested/a/b/c/d/e/f/g YANG schema element.
type OpenconfigNested_A_B_C_D_E_F_GPath struct {
	*ygnmi.NodePath
}

// OpenconfigNested_A_B_C_D_E_F_GPathAny represents the wildcard version of the /openconfig-nested/a/b/c/d/e/f/g YANG schema element.
type OpenconfigNested_A_B_C_D_E_F_GPathAny struct {
	*ygnmi.NodePath
}

// H (container):
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "h"
//	Path from root:       "/a/b/c/d/e/f/g/h"
func (n *OpenconfigNested_A_B_C_D_E_F_GPath) H() *OpenconfigNested_A_B_C_D_E_F_G_HPath {
	return &OpenconfigNested_A_B_C_D_E_F_G_HPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"h"},
			map[string]interface{}{},
			n,
		),
	}
}

// H (container):
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "h"
//	Path from root:       "/a/b/c/d/e/f/g/h"
func (n *OpenconfigNested_A_B_C_D_E_F_GPathAny) H() *OpenconfigNested_A_B_C_D_E_F_G_HPathAny {
	return &OpenconfigNested_A_B_C_D_E_F_G_HPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"h"},
			map[string]interface{}{},
			n,
		),
	}
}

// Query returns a Query that can be used in gNMI operations.
func (n *OpenconfigNested_A_B_C_D_E_F_GPath) Query() ygnmi.SingletonQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G] {
	return ygnmi.NewSingletonQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G](
		"OpenconfigNested_A_B_C_D_E_F_G",
		true,
		false,
		false,
		false,
		n,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)
}

// Query returns a Query that can be used in gNMI operations.
func (n *OpenconfigNested_A_B_C_D_E_F_GPathAny) Query() ygnmi.WildcardQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G] {
	return ygnmi.NewWildcardQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G](
		"OpenconfigNested_A_B_C_D_E_F_G",
		true,
		false,
		false,
		false,
		n,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
	)
}

// OpenconfigNested_A_B_C_D_E_F_G_HPath represents the /openconfig-nested/a/b/c/d/e/f/g/h YANG schema element.
type OpenconfigNested_A_B_C_D_E_F_G_HPath struct {
	*ygnmi.NodePath
}

// OpenconfigNested_A_B_C_D_E_F_G_HPathAny represents the wildcard version of the /openconfig-nested/a/b/c/d/e/f/g/h YANG schema element.
type OpenconfigNested_A_B_C_D_E_F_G_HPathAny struct {
	*ygnmi.NodePath
}

// I (container):
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "i"
//	Path from root:       "/a/b/c/d/e/f/g/h/i"
func (n *OpenconfigNested_A_B_C_D_E_F_G_HPath) I() *OpenconfigNested_A_B_C_D_E_F_G_H_IPath {
	return &OpenconfigNested_A_B_C_D_E_F_G_H_IPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"i"},
			map[string]interface{}{},
			n,
		),
	}
}

// I (container):
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "i"
//	Path from root:       "/a/b/c/d/e/f/g/h/i"
func (n *OpenconfigNested_A_B_C_D_E_F_G_HPathAny) I() *OpenconfigNested_A_B_C_D_E_F_G_H_IPathAny {
	return &OpenconfigNested_A_B_C_D_E_F_G_H_IPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"i"},
			map[string]interface{}{},
			n,
		),
	}
}

// Query returns a Query that can be used in gNMI operations.
func (n *OpenconfigNested_A_B_C_D_E_F_G_HPath) Query() ygnmi.SingletonQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G_H] {
	return ygnmi.NewSingletonQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G_H](
		"OpenconfigNested_A_B_C_D_E_F_G_H",
		true,
		false,
		false,
		false,
		n,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)
}

// Query returns a Query that can be used in gNMI operations.
func (n *OpenconfigNested_A_B_C_D_E_F_G_HPathAny) Query() ygnmi.WildcardQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G_H] {
	return ygnmi.NewWildcardQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G_H](
		"OpenconfigNested_A_B_C_D_E_F_G_H",
		true,
		false,
		false,
		false,
		n,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
	)
}

// OpenconfigNested_A_B_C_D_E_F_G_H_IPath represents the /openconfig-nested/a/b/c/d/e/f/g/h/i YANG schema element.
type OpenconfigNested_A_B_C_D_E_F_G_H_IPath struct {
	*ygnmi.NodePath
}

// OpenconfigNested_A_B_C_D_E_F_G_H_IPathAny represents the wildcard version of the /openconfig-nested/a/b/c/d/e/f/g/h/i YANG schema element.
type OpenconfigNested_A_B_C_D_E_F_G_H_IPathAny struct {
	*ygnmi.NodePath
}

// J (container):
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "j"
//	Path from root:       "/a/b/c/d/e/f/g/h/i/j"
func (n *OpenconfigNested_A_B_C_D_E_F_G_H_IPath) J() *OpenconfigNested_A_B_C_D_E_F_G_H_I_JPath {
	return &OpenconfigNested_A_B_C_D_E_F_G_H_I_JPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"j"},
			map[string]interface{}{},
			n,
		),
	}
}

// J (container):
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "j"
//	Path from root:       "/a/b/c/d/e/f/g/h/i/j"
func (n *OpenconfigNested_A_B_C_D_E_F_G_H_IPathAny) J() *OpenconfigNested_A_B_C_D_E_F_G_H_I_JPathAny {
	return &OpenconfigNested_A_B_C_D_E_F_G_H_I_JPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"j"},
			map[string]interface{}{},
			n,
		),
	}
}

// Query returns a Query that can be used in gNMI operations.
func (n *OpenconfigNested_A_B_C_D_E_F_G_H_IPath) Query() ygnmi.SingletonQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G_H_I] {
	return ygnmi.NewSingletonQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G_H_I](
		"OpenconfigNested_A_B_C_D_E_F_G_H_I",
		true,
		false,
		false,
		false,
		n,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)
}

// Query returns a Query that can be used in gNMI operations.
func (n *OpenconfigNested_A_B_C_D_E_F_G_H_IPathAny) Query() ygnmi.WildcardQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G_H_I] {
	return ygnmi.NewWildcardQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G_H_I](
		"OpenconfigNested_A_B_C_D_E_F_G_H_I",
		true,
		false,
		false,
		false,
		n,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
	)
}

// OpenconfigNested_A_B_C_D_E_F_G_H_I_JPath represents the /openconfig-nested/a/b/c/d/e/f/g/h/i/j YANG schema element.
type OpenconfigNested_A_B_C_D_E_F_G_H_I_JPath struct {
	*ygnmi.NodePath
}

// OpenconfigNested_A_B_C_D_E_F_G_H_I_JPathAny represents the wildcard version of the /openconfig-nested/a/b/c/d/e/f/g/h/i/j YANG schema element.
type OpenconfigNested_A_B_C_D_E_F_G_H_I_JPathAny struct {
	*ygnmi.NodePath
}

// K (container):
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "k"
//	Path from root:       "/a/b/c/d/e/f/g/h/i/j/k"
func (n *OpenconfigNested_A_B_C_D_E_F_G_H_I_JPath) K() *OpenconfigNested_A_B_C_D_E_F_G_H_I_J_KPath {
	return &OpenconfigNested_A_B_C_D_E_F_G_H_I_J_KPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"k"},
			map[string]interface{}{},
			n,
		),
	}
}

// K (container):
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "k"
//	Path from root:       "/a/b/c/d/e/f/g/h/i/j/k"
func (n *OpenconfigNested_A_B_C_D_E_F_G_H_I_JPathAny) K() *OpenconfigNested_A_B_C_D_E_F_G_H_I_J_KPathAny {
	return &OpenconfigNested_A_B_C_D_E_F_G_H_I_J_KPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"k"},
			map[string]interface{}{},
			n,
		),
	}
}

// Query returns a Query that can be used in gNMI operations.
func (n *OpenconfigNested_A_B_C_D_E_F_G_H_I_JPath) Query() ygnmi.SingletonQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G_H_I_J] {
	return ygnmi.NewSingletonQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G_H_I_J](
		"OpenconfigNested_A_B_C_D_E_F_G_H_I_J",
		true,
		false,
		false,
		false,
		n,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)
}

// Query returns a Query that can be used in gNMI operations.
func (n *OpenconfigNested_A_B_C_D_E_F_G_H_I_JPathAny) Query() ygnmi.WildcardQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G_H_I_J] {
	return ygnmi.NewWildcardQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G_H_I_J](
		"OpenconfigNested_A_B_C_D_E_F_G_H_I_J",
		true,
		false,
		false,
		false,
		n,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
	)
}

// OpenconfigNested_A_B_C_D_E_F_G_H_I_J_KPath represents the /openconfig-nested/a/b/c/d/e/f/g/h/i/j/k YANG schema element.
type OpenconfigNested_A_B_C_D_E_F_G_H_I_J_KPath struct {
	*ygnmi.NodePath
}

// OpenconfigNested_A_B_C_D_E_F_G_H_I_J_KPathAny represents the wildcard version of the /openconfig-nested/a/b/c/d/e/f/g/h/i/j/k YANG schema element.
type OpenconfigNested_A_B_C_D_E_F_G_H_I_J_KPathAny struct {
	*ygnmi.NodePath
}

// L (container):
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "l"
//	Path from root:       "/a/b/c/d/e/f/g/h/i/j/k/l"
func (n *OpenconfigNested_A_B_C_D_E_F_G_H_I_J_KPath) L() *OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_LPath {
	return &OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_LPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"l"},
			map[string]interface{}{},
			n,
		),
	}
}

// L (container):
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "l"
//	Path from root:       "/a/b/c/d/e/f/g/h/i/j/k/l"
func (n *OpenconfigNested_A_B_C_D_E_F_G_H_I_J_KPathAny) L() *OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_LPathAny {
	return &OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_LPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"l"},
			map[string]interface{}{},
			n,
		),
	}
}

// Query returns a Query that can be used in gNMI operations.
func (n *OpenconfigNested_A_B_C_D_E_F_G_H_I_J_KPath) Query() ygnmi.SingletonQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K] {
	return ygnmi.NewSingletonQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K](
		"OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K",
		true,
		false,
		false,
		false,
		n,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)
}

// Query returns a Query that can be used in gNMI operations.
func (n *OpenconfigNested_A_B_C_D_E_F_G_H_I_J_KPathAny) Query() ygnmi.WildcardQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K] {
	return ygnmi.NewWildcardQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K](
		"OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K",
		true,
		false,
		false,
		false,
		n,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
	)
}

// OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_LPath represents the /openconfig-nested/a/b/c/d/e/f/g/h/i/j/k/l YANG schema element.
type OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_LPath struct {
	*ygnmi.NodePath
}

// OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_LPathAny represents the wildcard version of the /openconfig-nested/a/b/c/d/e/f/g/h/i/j/k/l YANG schema element.
type OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_LPathAny struct {
	*ygnmi.NodePath
}

// M (container):
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "m"
//	Path from root:       "/a/b/c/d/e/f/g/h/i/j/k/l/m"
func (n *OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_LPath) M() *OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_MPath {
	return &OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_MPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"m"},
			map[string]interface{}{},
			n,
		),
	}
}

// M (container):
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "m"
//	Path from root:       "/a/b/c/d/e/f/g/h/i/j/k/l/m"
func (n *OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_LPathAny) M() *OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_MPathAny {
	return &OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_MPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"m"},
			map[string]interface{}{},
			n,
		),
	}
}

// Query returns a Query that can be used in gNMI operations.
func (n *OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_LPath) Query() ygnmi.SingletonQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L] {
	return ygnmi.NewSingletonQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L](
		"OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L",
		true,
		false,
		false,
		false,
		n,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)
}

// Query returns a Query that can be used in gNMI operations.
func (n *OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_LPathAny) Query() ygnmi.WildcardQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L] {
	return ygnmi.NewWildcardQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L](
		"OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L",
		true,
		false,
		false,
		false,
		n,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
	)
}

// OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_MPath represents the /openconfig-nested/a/b/c/d/e/f/g/h/i/j/k/l/m YANG schema element.
type OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_MPath struct {
	*ygnmi.NodePath
}

// OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_MPathAny represents the wildcard version of the /openconfig-nested/a/b/c/d/e/f/g/h/i/j/k/l/m YANG schema element.
type OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_MPathAny struct {
	*ygnmi.NodePath
}

// State (container):
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "state"
//	Path from root:       "/a/b/c/d/e/f/g/h/i/j/k/l/m/state"
func (n *OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_MPath) State() *OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_StatePath {
	return &OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_StatePath{
		NodePath: ygnmi.NewNodePath(
			[]string{"state"},
			map[string]interface{}{},
			n,
		),
	}
}

// State (container):
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "state"
//	Path from root:       "/a/b/c/d/e/f/g/h/i/j/k/l/m/state"
func (n *OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_MPathAny) State() *OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_StatePathAny {
	return &OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_StatePathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"state"},
			map[string]interface{}{},
			n,
		),
	}
}

// Query returns a Query that can be used in gNMI operations.
func (n *OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_MPath) Query() ygnmi.SingletonQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M] {
	return ygnmi.NewSingletonQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M](
		"OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M",
		true,
		false,
		false,
		false,
		n,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)
}

// Query returns a Query that can be used in gNMI operations.
func (n *OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_MPathAny) Query() ygnmi.WildcardQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M] {
	return ygnmi.NewWildcardQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M](
		"OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M",
		true,
		false,
		false,
		false,
		n,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
	)
}

// OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_State_FooPath represents the /openconfig-nested/a/b/c/d/e/f/g/h/i/j/k/l/m/state/foo YANG schema element.
type OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_State_FooPath struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_State_FooPathAny represents the wildcard version of the /openconfig-nested/a/b/c/d/e/f/g/h/i/j/k/l/m/state/foo YANG schema element.
type OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_State_FooPathAny struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// Query returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "foo"
//	Path from root:       "/a/b/c/d/e/f/g/h/i/j/k/l/m/state/foo"
func (n *OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_State_FooPath) Query() ygnmi.SingletonQuery[string] {
	return ygnmi.NewSingletonQuery[string](
		"OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_State",
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"foo"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (string, bool) {
			ret := gs.(*oc.OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_State).Foo
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_State) },
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)
}

// Query returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "foo"
//	Path from root:       "/a/b/c/d/e/f/g/h/i/j/k/l/m/state/foo"
func (n *OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_State_FooPathAny) Query() ygnmi.WildcardQuery[string] {
	return ygnmi.NewWildcardQuery[string](
		"OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_State",
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"foo"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (string, bool) {
			ret := gs.(*oc.OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_State).Foo
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_State) },
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
	)
}

// OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_StatePath represents the /openconfig-nested/a/b/c/d/e/f/g/h/i/j/k/l/m/state YANG schema element.
type OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_StatePath struct {
	*ygnmi.NodePath
}

// OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_StatePathAny represents the wildcard version of the /openconfig-nested/a/b/c/d/e/f/g/h/i/j/k/l/m/state YANG schema element.
type OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_StatePathAny struct {
	*ygnmi.NodePath
}

// Foo (leaf):
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "foo"
//	Path from root:       "/a/b/c/d/e/f/g/h/i/j/k/l/m/state/foo"
func (n *OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_StatePath) Foo() *OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_State_FooPath {
	return &OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_State_FooPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"foo"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
}

// Foo (leaf):
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "foo"
//	Path from root:       "/a/b/c/d/e/f/g/h/i/j/k/l/m/state/foo"
func (n *OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_StatePathAny) Foo() *OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_State_FooPathAny {
	return &OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_State_FooPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"foo"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
}

// Query returns a Query that can be used in gNMI operations.
func (n *OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_StatePath) Query() ygnmi.SingletonQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_State] {
	return ygnmi.NewSingletonQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_State](
		"OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_State",
		true,
		false,
		false,
		false,
		n,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)
}

// Query returns a Query that can be used in gNMI operations.
func (n *OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_StatePathAny) Query() ygnmi.WildcardQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_State] {
	return ygnmi.NewWildcardQuery[*oc.OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_State](
		"OpenconfigNested_A_B_C_D_E_F_G_H_I_J_K_L_M_State",
		true,
		false,
		false,
		false,
		n,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
	)
}

// OpenconfigNested_ContainerPath represents the /openconfig-nested/container YANG schema element.
type OpenconfigNested_ContainerPath struct {
	*ygnmi.NodePath
}

// OpenconfigNested_ContainerPathAny represents the wildcard version of the /openconfig-nested/container YANG schema element.
type OpenconfigNested_ContainerPathAny struct {
	*ygnmi.NodePath
}

// Query returns a Query that can be used in gNMI operations.
func (n *OpenconfigNested_ContainerPath) Query() ygnmi.SingletonQuery[*oc.OpenconfigNested_Container] {
	return ygnmi.NewSingletonQuery[*oc.OpenconfigNested_Container](
		"OpenconfigNested_Container",
		true,
		false,
		false,
		false,
		n,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)
}

// Query returns a Query that can be used in gNMI operations.
func (n *OpenconfigNested_ContainerPathAny) Query() ygnmi.WildcardQuery[*oc.OpenconfigNested_Container] {
	return ygnmi.NewWildcardQuery[*oc.OpenconfigNested_Container](
		"OpenconfigNested_Container",
		true,
		false,
		false,
		false,
		n,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
	)
}
