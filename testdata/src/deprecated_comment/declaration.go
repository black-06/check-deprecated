package deprecated_comment

// Deprecated: VarDeprecated by GenDecl ValueSpec
var VarDeprecated = ""

// DEPRECATED: vars1/2/3 by GenDecl ValueSpec
var (
	// deprecated. vars1 by ValueSpec
	vars1 = "" // want "malformed deprecated header: the proper format is `Deprecated: <text>`"
	vars2 = "" // want "malformed deprecated header: use `Deprecated: ` \\(note the casing\\) instead of `DEPRECATED: `"
	vars3 = "" // want "malformed deprecated header: use `Deprecated: ` \\(note the casing\\) instead of `DEPRECATED: `"
)

// ConstDeprecated
// it's deprecated. ConstDeprecated by GenDecl ValueSpec
const ConstDeprecated = "" // want "malformed deprecated header: the proper format is `Deprecated: <text>`"

// NOTE: deprecated. consts 1/2/3 by GenDecl ValueSpec
const (
	// deprecated, consts1 by ValueSpec
	consts1 = iota // want "malformed deprecated header: the proper format is `Deprecated: <text>`"
	consts2        // want "malformed deprecated header: the proper format is `Deprecated: <text>`"
	consts3        // want "malformed deprecated header: the proper format is `Deprecated: <text>`"
)

// FuncDeprecated
//
// Deprecated: don't use FuncDeprecated by FuncDecl
func FuncDeprecated() {
}

type Struct struct{}

// Deprecated, don't use it
func (p Struct) StructFun() {} // want "malformed deprecated header: the proper format is `Deprecated: <text>`"

// Deprecated: use s3 instead of StructDeprecated, by GenDecl TypeSpec
type StructDeprecated struct{}

func (p StructDeprecated) Fun() {}

// Deprecated.
type StructDeprecated2 struct{} // want "malformed deprecated header: the proper format is `Deprecated: <text>`"

// InterfaceDeprecated
//
// Deprecated, InterfaceDeprecated by GenDecl TypeSpec
type InterfaceDeprecated interface{} // want "malformed deprecated header: the proper format is `Deprecated: <text>`"

// Deprecated struct 1/2/3 by GenDecl TypeSpec
type (
	// Deprecated struct1 by TypeSpec
	struct1 struct{} // want "malformed deprecated header: the proper format is `Deprecated: <text>`"
	struct2 struct { // want "malformed deprecated header: the proper format is `Deprecated: <text>`"
		F1 string
		// Deprecated F2 by Field
		F2 string // want "malformed deprecated header: the proper format is `Deprecated: <text>`"
	}
	struct3 struct{} // want "malformed deprecated header: the proper format is `Deprecated: <text>`"
)

func (s struct3) fun1() {}

// Deprecated fun2 by FuncDecl
func (s struct3) fun2() {} // want "malformed deprecated header: the proper format is `Deprecated: <text>`"

// Deprecated interface 1/2/3
// by GenDecl TypeSpec
type (
	// Deprecated interface1 by TypeSpec
	interface1 interface{} // want "malformed deprecated header: the proper format is `Deprecated: <text>`"
	interface2 interface { // want "malformed deprecated header: the proper format is `Deprecated: <text>`"
	}
	interface3 interface { // want "malformed deprecated header: the proper format is `Deprecated: <text>`"
		fun1()
		// deprecated. interface3 fun2 by Field
		fun2() // want "malformed deprecated header: the proper format is `Deprecated: <text>`"
	}
)
