// Copyright 2022 Google LLC
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

/*
Package root is a generated package which contains definitions
of structs which generate gNMI paths for a YANG schema. The generated paths are
based on a compressed form of the schema.

This package was generated by ygnmi version: (devel): (ygot: v0.23.0)
using the following YANG input files:
	- ../../pathgen/testdata/yang/openconfig-simple.yang
	- ../../pathgen/testdata/yang/openconfig-withlistval.yang
	- ../../pathgen/testdata/yang/openconfig-nested.yang
Imported modules were sourced from:
*/
package root

import (
	oc "github.com/openconfig/ygnmi/internal/exampleoc"
	"github.com/openconfig/ygnmi/internal/exampleoc/nested"
	"github.com/openconfig/ygnmi/internal/exampleoc/simple"
	"github.com/openconfig/ygnmi/internal/exampleoc/withlistval"
	"github.com/openconfig/ygnmi/ygnmi"
	"github.com/openconfig/ygot/ytypes"
)

// RootPath represents the /root YANG schema element.
type RootPath struct {
	*ygnmi.DeviceRootBase
}

// New returns a new path object from which YANG paths can be constructed.
func New() *RootPath {
	return &RootPath{ygnmi.NewDeviceRootBase()}
}

// A (container):
// 	Defining module:      "openconfig-nested"
// 	Instantiating module: "openconfig-nested"
// 	Path from parent:     "a"
// 	Path from root:       "/a"
func (n *RootPath) A() *nested.APath {
	return &nested.APath{
		NodePath: ygnmi.NewNodePath(
			[]string{"a"},
			map[string]interface{}{},
			n,
		),
	}
}

// Model (container):
// 	Defining module:      "openconfig-withlistval"
// 	Instantiating module: "openconfig-withlistval"
// 	Path from parent:     "model"
// 	Path from root:       "/model"
func (n *RootPath) Model() *withlistval.ModelPath {
	return &withlistval.ModelPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"model"},
			map[string]interface{}{},
			n,
		),
	}
}

// Parent (container): I am a parent container
// that has 4 children.
// 	Defining module:      "openconfig-simple"
// 	Instantiating module: "openconfig-simple"
// 	Path from parent:     "parent"
// 	Path from root:       "/parent"
func (n *RootPath) Parent() *simple.ParentPath {
	return &simple.ParentPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"parent"},
			map[string]interface{}{},
			n,
		),
	}
}

// RemoteContainer (container):
// 	Defining module:      "openconfig-remote"
// 	Instantiating module: "openconfig-simple"
// 	Path from parent:     "remote-container"
// 	Path from root:       "/remote-container"
func (n *RootPath) RemoteContainer() *simple.RemoteContainerPath {
	return &simple.RemoteContainerPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"remote-container"},
			map[string]interface{}{},
			n,
		),
	}
}

// State returns a Query that can be used in gNMI operations.
func (n *RootPath) State() ygnmi.SingletonQuery[*oc.Root] {
	return ygnmi.NewNonLeafSingletonQuery[*oc.Root](
		"Root",
		true,
		n,
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// Config returns a Query that can be used in gNMI operations.
func (n *RootPath) Config() ygnmi.ConfigQuery[*oc.Root] {
	return ygnmi.NewNonLeafConfigQuery[*oc.Root](
		"Root",
		false,
		n,
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}
