package deprecated

// Deprecated: VarDeprecated by GenDecl ValueSpec
var VarDeprecated = ""

// DEPRECATED: vars1/2/3 by GenDecl ValueSpec
var (
	// deprecated. vars1 by ValueSpec
	vars1 = ""
	vars2 = ""
	vars3 = ""
)

// ConstDeprecated
// it's deprecated. ConstDeprecated by GenDecl ValueSpec
const ConstDeprecated = ""

// NOTE: deprecated. consts 1/2/3 by GenDecl ValueSpec
const (
	// deprecated, consts1 by ValueSpec
	consts1 = iota
	consts2
	consts3
)

// FuncDeprecated
//
// Deprecated: don't use FuncDeprecated by FuncDecl
func FuncDeprecated() {
}

type Struct struct{}

// Deprecated, don't use it
func (p Struct) StructFun() {}

// Deprecated: use s3 instead of StructDeprecated, by GenDecl TypeSpec
type StructDeprecated struct{}

func (p StructDeprecated) Fun() {} // want "using deprecated: use s3 instead of StructDeprecated, by GenDecl TypeSpec"

// Deprecated.
type StructDeprecated2 struct{}

// InterfaceDeprecated
//
// Deprecated, InterfaceDeprecated by GenDecl TypeSpec
type InterfaceDeprecated interface{}

// Deprecated struct 1/2/3 by GenDecl TypeSpec
type (
	// Deprecated struct1 by TypeSpec
	struct1 struct{}
	struct2 struct {
		F1 string
		// Deprecated F2 by Field
		F2 string
	}
	struct3 struct{}
)

func (s struct3) fun1() {} // want "using deprecated: struct 1/2/3 by GenDecl TypeSpec"

// Deprecated fun2 by FuncDecl
func (s struct3) fun2() {} // want "using deprecated: struct 1/2/3 by GenDecl TypeSpec"

// Deprecated interface 1/2/3
// by GenDecl TypeSpec
type (
	// Deprecated interface1 by TypeSpec
	interface1 interface{}
	interface2 interface {
	}
	interface3 interface {
		fun1()
		// deprecated. interface3 fun2 by Field
		fun2()
	}
)
