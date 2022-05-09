/*
Package withlist is a generated package which contains definitions
of structs which generate gNMI paths for a YANG schema. The generated paths are
based on a compressed form of the schema.

This package was generated by /usr/local/google/home/dgrau/go/pkg/mod/github.com/openconfig/ygot@v0.18.0/genutil/names.go
using the following YANG input files:
	- ../../pathgen/testdata/yang/openconfig-simple.yang
	- ../../pathgen/testdata/yang/openconfig-withlist.yang
Imported modules were sourced from:
*/
package withlist

import (
	oc "github.com/openconfig/ygnmi/internal/exampleoc"
	"github.com/openconfig/ygnmi/ygnmi"
	"github.com/openconfig/ygot/ygot"
	"github.com/openconfig/ygot/ytypes"
)

// Model represents the /openconfig-withlist/model YANG schema element.
type Model struct {
	*ygot.NodePath
}

// ModelAny represents the wildcard version of the /openconfig-withlist/model YANG schema element.
type ModelAny struct {
	*ygot.NodePath
}

// MultiKeyAny (list):
// ----------------------------------------
// Defining module: "openconfig-withlist"
// Instantiating module: "openconfig-withlist"
// Path from parent: "b/multi-key"
// Path from root: "/model/b/multi-key"
func (n *Model) MultiKeyAny() *Model_MultiKeyAny {
	return &Model_MultiKeyAny{
		NodePath: ygot.NewNodePath(
			[]string{"b", "multi-key"},
			map[string]interface{}{"key1": "*", "key2": "*"},
			n,
		),
	}
}

// MultiKeyAny (list):
// ----------------------------------------
// Defining module: "openconfig-withlist"
// Instantiating module: "openconfig-withlist"
// Path from parent: "b/multi-key"
// Path from root: "/model/b/multi-key"
func (n *ModelAny) MultiKeyAny() *Model_MultiKeyAny {
	return &Model_MultiKeyAny{
		NodePath: ygot.NewNodePath(
			[]string{"b", "multi-key"},
			map[string]interface{}{"key1": "*", "key2": "*"},
			n,
		),
	}
}

// WithKey1 sets Model_MultiKeyAny's key "key1" to the specified value.
// Key1: uint32
func (n *Model_MultiKeyAny) WithKey1(Key1 uint32) *Model_MultiKeyAny {
	ygot.ModifyKey(n.NodePath, "key1", Key1)
	return n
}

// WithKey2 sets Model_MultiKeyAny's key "key2" to the specified value.
// Key2: uint64
func (n *Model_MultiKeyAny) WithKey2(Key2 uint64) *Model_MultiKeyAny {
	ygot.ModifyKey(n.NodePath, "key2", Key2)
	return n
}

// SingleKeyAny (list):
// ----------------------------------------
// Defining module: "openconfig-withlist"
// Instantiating module: "openconfig-withlist"
// Path from parent: "a/single-key"
// Path from root: "/model/a/single-key"
// Key (wildcarded): string
func (n *Model) SingleKeyAny() *Model_SingleKeyAny {
	return &Model_SingleKeyAny{
		NodePath: ygot.NewNodePath(
			[]string{"a", "single-key"},
			map[string]interface{}{"key": "*"},
			n,
		),
	}
}

// SingleKeyAny (list):
// ----------------------------------------
// Defining module: "openconfig-withlist"
// Instantiating module: "openconfig-withlist"
// Path from parent: "a/single-key"
// Path from root: "/model/a/single-key"
// Key (wildcarded): string
func (n *ModelAny) SingleKeyAny() *Model_SingleKeyAny {
	return &Model_SingleKeyAny{
		NodePath: ygot.NewNodePath(
			[]string{"a", "single-key"},
			map[string]interface{}{"key": "*"},
			n,
		),
	}
}

// SingleKey (list):
// ----------------------------------------
// Defining module: "openconfig-withlist"
// Instantiating module: "openconfig-withlist"
// Path from parent: "a/single-key"
// Path from root: "/model/a/single-key"
// Key: string
func (n *Model) SingleKey(Key string) *Model_SingleKey {
	return &Model_SingleKey{
		NodePath: ygot.NewNodePath(
			[]string{"a", "single-key"},
			map[string]interface{}{"key": Key},
			n,
		),
	}
}

// SingleKey (list):
// ----------------------------------------
// Defining module: "openconfig-withlist"
// Instantiating module: "openconfig-withlist"
// Path from parent: "a/single-key"
// Path from root: "/model/a/single-key"
// Key: string
func (n *ModelAny) SingleKey(Key string) *Model_SingleKeyAny {
	return &Model_SingleKeyAny{
		NodePath: ygot.NewNodePath(
			[]string{"a", "single-key"},
			map[string]interface{}{"key": Key},
			n,
		),
	}
}

// State returns a Query that can be used in gNMI operations.
func (n *Model) State() ygnmi.SingletonQuery[*oc.Model] {
	return ygnmi.NewNonLeafSingletonQuery[*oc.Model](
		"Model",
		true,
		n,
		&ytypes.Schema{
			Root:       &oc.Device{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// State returns a Query that can be used in gNMI operations.
func (n *ModelAny) State() ygnmi.WildcardQuery[*oc.Model] {
	return ygnmi.NewNonLeafWildcardQuery[*oc.Model](
		"Model",
		true,
		n,
		&ytypes.Schema{
			Root:       &oc.Device{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// Config returns a Query that can be used in gNMI operations.
func (n *Model) Config() ygnmi.ConfigQuery[*oc.Model] {
	return ygnmi.NewNonLeafConfigQuery[*oc.Model](
		"Model",
		false,
		n,
		&ytypes.Schema{
			Root:       &oc.Device{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// Config returns a Query that can be used in gNMI operations.
func (n *ModelAny) Config() ygnmi.WildcardQuery[*oc.Model] {
	return ygnmi.NewNonLeafWildcardQuery[*oc.Model](
		"Model",
		false,
		n,
		&ytypes.Schema{
			Root:       &oc.Device{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// Model_MultiKey represents the /openconfig-withlist/model/b/multi-key YANG schema element.
type Model_MultiKey struct {
	*ygot.NodePath
}

// Model_MultiKeyAny represents the wildcard version of the /openconfig-withlist/model/b/multi-key YANG schema element.
type Model_MultiKeyAny struct {
	*ygot.NodePath
}

// Model_MultiKey_Key1 represents the /openconfig-withlist/model/b/multi-key/state/key1 YANG schema element.
type Model_MultiKey_Key1 struct {
	parent ygot.PathStruct
}

// Model_MultiKey_Key1Any represents the wildcard version of the /openconfig-withlist/model/b/multi-key/state/key1 YANG schema element.
type Model_MultiKey_Key1Any struct {
	parent ygot.PathStruct
}

// Model_MultiKey_Key2 represents the /openconfig-withlist/model/b/multi-key/state/key2 YANG schema element.
type Model_MultiKey_Key2 struct {
	parent ygot.PathStruct
}

// Model_MultiKey_Key2Any represents the wildcard version of the /openconfig-withlist/model/b/multi-key/state/key2 YANG schema element.
type Model_MultiKey_Key2Any struct {
	parent ygot.PathStruct
}

// Key1 corresponds to an ambiguous path; use .Config() or .State() to get a resolved path for this leaf.
// Note: The returned struct does not implement the PathStruct interface.
func (n *Model_MultiKey) Key1() *Model_MultiKey_Key1 {
	return &Model_MultiKey_Key1{
		parent: n,
	}
}

// Key1 corresponds to an ambiguous path; use .Config() or .State() to get a resolved path for this leaf.
// Note: The returned struct does not implement the PathStruct interface.
func (n *Model_MultiKeyAny) Key1() *Model_MultiKey_Key1Any {
	return &Model_MultiKey_Key1Any{
		parent: n,
	}
}

// Key2 corresponds to an ambiguous path; use .Config() or .State() to get a resolved path for this leaf.
// Note: The returned struct does not implement the PathStruct interface.
func (n *Model_MultiKey) Key2() *Model_MultiKey_Key2 {
	return &Model_MultiKey_Key2{
		parent: n,
	}
}

// Key2 corresponds to an ambiguous path; use .Config() or .State() to get a resolved path for this leaf.
// Note: The returned struct does not implement the PathStruct interface.
func (n *Model_MultiKeyAny) Key2() *Model_MultiKey_Key2Any {
	return &Model_MultiKey_Key2Any{
		parent: n,
	}
}

// State returns a Query that can be used in gNMI operations.
func (n *Model_MultiKey) State() ygnmi.SingletonQuery[*oc.Model_MultiKey] {
	return ygnmi.NewNonLeafSingletonQuery[*oc.Model_MultiKey](
		"Model_MultiKey",
		true,
		n,
		&ytypes.Schema{
			Root:       &oc.Device{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// State returns a Query that can be used in gNMI operations.
func (n *Model_MultiKeyAny) State() ygnmi.WildcardQuery[*oc.Model_MultiKey] {
	return ygnmi.NewNonLeafWildcardQuery[*oc.Model_MultiKey](
		"Model_MultiKey",
		true,
		n,
		&ytypes.Schema{
			Root:       &oc.Device{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// Config returns a Query that can be used in gNMI operations.
func (n *Model_MultiKey) Config() ygnmi.ConfigQuery[*oc.Model_MultiKey] {
	return ygnmi.NewNonLeafConfigQuery[*oc.Model_MultiKey](
		"Model_MultiKey",
		false,
		n,
		&ytypes.Schema{
			Root:       &oc.Device{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// Config returns a Query that can be used in gNMI operations.
func (n *Model_MultiKeyAny) Config() ygnmi.WildcardQuery[*oc.Model_MultiKey] {
	return ygnmi.NewNonLeafWildcardQuery[*oc.Model_MultiKey](
		"Model_MultiKey",
		false,
		n,
		&ytypes.Schema{
			Root:       &oc.Device{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// State returns a Query that can be used in gNMI operations.
// ----------------------------------------
// Defining module: "openconfig-withlist"
// Instantiating module: "openconfig-withlist"
// Path from parent: "state/key1"
// Path from root: "/model/b/multi-key/state/key1"
func (n *Model_MultiKey_Key1) State() ygnmi.SingletonQuery[uint32] {
	return ygnmi.NewLeafSingletonQuery[uint32](
		"Model_MultiKey",
		true,
		true,
		ygot.NewNodePath(
			[]string{"state", "key1"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (uint32, bool) {
			ret := gs.(*oc.Model_MultiKey).Key1
			if ret == nil {
				var zero uint32
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Model_MultiKey) },
		&ytypes.Schema{
			Root:       &oc.Device{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// State returns a Query that can be used in gNMI operations.
// ----------------------------------------
// Defining module: "openconfig-withlist"
// Instantiating module: "openconfig-withlist"
// Path from parent: "state/key1"
// Path from root: "/model/b/multi-key/state/key1"
func (n *Model_MultiKey_Key1Any) State() ygnmi.WildcardQuery[uint32] {
	return ygnmi.NewLeafWildcardQuery[uint32](
		"Model_MultiKey",
		true,
		true,
		ygot.NewNodePath(
			[]string{"state", "key1"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (uint32, bool) {
			ret := gs.(*oc.Model_MultiKey).Key1
			if ret == nil {
				var zero uint32
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Model_MultiKey) },
		&ytypes.Schema{
			Root:       &oc.Device{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// Config returns a Query that can be used in gNMI operations.
// ----------------------------------------
// Defining module: "openconfig-withlist"
// Instantiating module: "openconfig-withlist"
// Path from parent: "config/key1"
// Path from root: "/model/b/multi-key/config/key1"
func (n *Model_MultiKey_Key1) Config() ygnmi.ConfigQuery[uint32] {
	return ygnmi.NewLeafConfigQuery[uint32](
		"Model_MultiKey",
		false,
		true,
		ygot.NewNodePath(
			[]string{"config", "key1"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (uint32, bool) {
			ret := gs.(*oc.Model_MultiKey).Key1
			if ret == nil {
				var zero uint32
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Model_MultiKey) },
		&ytypes.Schema{
			Root:       &oc.Device{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// Config returns a Query that can be used in gNMI operations.
// ----------------------------------------
// Defining module: "openconfig-withlist"
// Instantiating module: "openconfig-withlist"
// Path from parent: "config/key1"
// Path from root: "/model/b/multi-key/config/key1"
func (n *Model_MultiKey_Key1Any) Config() ygnmi.WildcardQuery[uint32] {
	return ygnmi.NewLeafWildcardQuery[uint32](
		"Model_MultiKey",
		false,
		true,
		ygot.NewNodePath(
			[]string{"config", "key1"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (uint32, bool) {
			ret := gs.(*oc.Model_MultiKey).Key1
			if ret == nil {
				var zero uint32
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Model_MultiKey) },
		&ytypes.Schema{
			Root:       &oc.Device{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// State returns a Query that can be used in gNMI operations.
// ----------------------------------------
// Defining module: "openconfig-withlist"
// Instantiating module: "openconfig-withlist"
// Path from parent: "state/key2"
// Path from root: "/model/b/multi-key/state/key2"
func (n *Model_MultiKey_Key2) State() ygnmi.SingletonQuery[uint64] {
	return ygnmi.NewLeafSingletonQuery[uint64](
		"Model_MultiKey",
		true,
		true,
		ygot.NewNodePath(
			[]string{"state", "key2"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (uint64, bool) {
			ret := gs.(*oc.Model_MultiKey).Key2
			if ret == nil {
				var zero uint64
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Model_MultiKey) },
		&ytypes.Schema{
			Root:       &oc.Device{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// State returns a Query that can be used in gNMI operations.
// ----------------------------------------
// Defining module: "openconfig-withlist"
// Instantiating module: "openconfig-withlist"
// Path from parent: "state/key2"
// Path from root: "/model/b/multi-key/state/key2"
func (n *Model_MultiKey_Key2Any) State() ygnmi.WildcardQuery[uint64] {
	return ygnmi.NewLeafWildcardQuery[uint64](
		"Model_MultiKey",
		true,
		true,
		ygot.NewNodePath(
			[]string{"state", "key2"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (uint64, bool) {
			ret := gs.(*oc.Model_MultiKey).Key2
			if ret == nil {
				var zero uint64
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Model_MultiKey) },
		&ytypes.Schema{
			Root:       &oc.Device{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// Config returns a Query that can be used in gNMI operations.
// ----------------------------------------
// Defining module: "openconfig-withlist"
// Instantiating module: "openconfig-withlist"
// Path from parent: "config/key2"
// Path from root: "/model/b/multi-key/config/key2"
func (n *Model_MultiKey_Key2) Config() ygnmi.ConfigQuery[uint64] {
	return ygnmi.NewLeafConfigQuery[uint64](
		"Model_MultiKey",
		false,
		true,
		ygot.NewNodePath(
			[]string{"config", "key2"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (uint64, bool) {
			ret := gs.(*oc.Model_MultiKey).Key2
			if ret == nil {
				var zero uint64
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Model_MultiKey) },
		&ytypes.Schema{
			Root:       &oc.Device{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// Config returns a Query that can be used in gNMI operations.
// ----------------------------------------
// Defining module: "openconfig-withlist"
// Instantiating module: "openconfig-withlist"
// Path from parent: "config/key2"
// Path from root: "/model/b/multi-key/config/key2"
func (n *Model_MultiKey_Key2Any) Config() ygnmi.WildcardQuery[uint64] {
	return ygnmi.NewLeafWildcardQuery[uint64](
		"Model_MultiKey",
		false,
		true,
		ygot.NewNodePath(
			[]string{"config", "key2"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (uint64, bool) {
			ret := gs.(*oc.Model_MultiKey).Key2
			if ret == nil {
				var zero uint64
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Model_MultiKey) },
		&ytypes.Schema{
			Root:       &oc.Device{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// Model_SingleKey represents the /openconfig-withlist/model/a/single-key YANG schema element.
type Model_SingleKey struct {
	*ygot.NodePath
}

// Model_SingleKeyAny represents the wildcard version of the /openconfig-withlist/model/a/single-key YANG schema element.
type Model_SingleKeyAny struct {
	*ygot.NodePath
}

// Model_SingleKey_Key represents the /openconfig-withlist/model/a/single-key/state/key YANG schema element.
type Model_SingleKey_Key struct {
	parent ygot.PathStruct
}

// Model_SingleKey_KeyAny represents the wildcard version of the /openconfig-withlist/model/a/single-key/state/key YANG schema element.
type Model_SingleKey_KeyAny struct {
	parent ygot.PathStruct
}

// Key corresponds to an ambiguous path; use .Config() or .State() to get a resolved path for this leaf.
// Note: The returned struct does not implement the PathStruct interface.
func (n *Model_SingleKey) Key() *Model_SingleKey_Key {
	return &Model_SingleKey_Key{
		parent: n,
	}
}

// Key corresponds to an ambiguous path; use .Config() or .State() to get a resolved path for this leaf.
// Note: The returned struct does not implement the PathStruct interface.
func (n *Model_SingleKeyAny) Key() *Model_SingleKey_KeyAny {
	return &Model_SingleKey_KeyAny{
		parent: n,
	}
}

// State returns a Query that can be used in gNMI operations.
func (n *Model_SingleKey) State() ygnmi.SingletonQuery[*oc.Model_SingleKey] {
	return ygnmi.NewNonLeafSingletonQuery[*oc.Model_SingleKey](
		"Model_SingleKey",
		true,
		n,
		&ytypes.Schema{
			Root:       &oc.Device{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// State returns a Query that can be used in gNMI operations.
func (n *Model_SingleKeyAny) State() ygnmi.WildcardQuery[*oc.Model_SingleKey] {
	return ygnmi.NewNonLeafWildcardQuery[*oc.Model_SingleKey](
		"Model_SingleKey",
		true,
		n,
		&ytypes.Schema{
			Root:       &oc.Device{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// Config returns a Query that can be used in gNMI operations.
func (n *Model_SingleKey) Config() ygnmi.ConfigQuery[*oc.Model_SingleKey] {
	return ygnmi.NewNonLeafConfigQuery[*oc.Model_SingleKey](
		"Model_SingleKey",
		false,
		n,
		&ytypes.Schema{
			Root:       &oc.Device{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// Config returns a Query that can be used in gNMI operations.
func (n *Model_SingleKeyAny) Config() ygnmi.WildcardQuery[*oc.Model_SingleKey] {
	return ygnmi.NewNonLeafWildcardQuery[*oc.Model_SingleKey](
		"Model_SingleKey",
		false,
		n,
		&ytypes.Schema{
			Root:       &oc.Device{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// State returns a Query that can be used in gNMI operations.
// ----------------------------------------
// Defining module: "openconfig-withlist"
// Instantiating module: "openconfig-withlist"
// Path from parent: "state/key"
// Path from root: "/model/a/single-key/state/key"
func (n *Model_SingleKey_Key) State() ygnmi.SingletonQuery[string] {
	return ygnmi.NewLeafSingletonQuery[string](
		"Model_SingleKey",
		true,
		true,
		ygot.NewNodePath(
			[]string{"state", "key"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (string, bool) {
			ret := gs.(*oc.Model_SingleKey).Key
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Model_SingleKey) },
		&ytypes.Schema{
			Root:       &oc.Device{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// State returns a Query that can be used in gNMI operations.
// ----------------------------------------
// Defining module: "openconfig-withlist"
// Instantiating module: "openconfig-withlist"
// Path from parent: "state/key"
// Path from root: "/model/a/single-key/state/key"
func (n *Model_SingleKey_KeyAny) State() ygnmi.WildcardQuery[string] {
	return ygnmi.NewLeafWildcardQuery[string](
		"Model_SingleKey",
		true,
		true,
		ygot.NewNodePath(
			[]string{"state", "key"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (string, bool) {
			ret := gs.(*oc.Model_SingleKey).Key
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Model_SingleKey) },
		&ytypes.Schema{
			Root:       &oc.Device{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// Config returns a Query that can be used in gNMI operations.
// ----------------------------------------
// Defining module: "openconfig-withlist"
// Instantiating module: "openconfig-withlist"
// Path from parent: "config/key"
// Path from root: "/model/a/single-key/config/key"
func (n *Model_SingleKey_Key) Config() ygnmi.ConfigQuery[string] {
	return ygnmi.NewLeafConfigQuery[string](
		"Model_SingleKey",
		false,
		true,
		ygot.NewNodePath(
			[]string{"config", "key"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (string, bool) {
			ret := gs.(*oc.Model_SingleKey).Key
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Model_SingleKey) },
		&ytypes.Schema{
			Root:       &oc.Device{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// Config returns a Query that can be used in gNMI operations.
// ----------------------------------------
// Defining module: "openconfig-withlist"
// Instantiating module: "openconfig-withlist"
// Path from parent: "config/key"
// Path from root: "/model/a/single-key/config/key"
func (n *Model_SingleKey_KeyAny) Config() ygnmi.WildcardQuery[string] {
	return ygnmi.NewLeafWildcardQuery[string](
		"Model_SingleKey",
		false,
		true,
		ygot.NewNodePath(
			[]string{"config", "key"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (string, bool) {
			ret := gs.(*oc.Model_SingleKey).Key
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Model_SingleKey) },
		&ytypes.Schema{
			Root:       &oc.Device{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}
