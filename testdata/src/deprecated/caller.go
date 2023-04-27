package deprecated

import (
	"fmt"
	"io/ioutil"

	"deprecated/pkg"
)

func Caller() {
	ioutil.ReadAll(nil)        // want "using deprecated: As of Go 1.16, this function simply calls io.ReadAll."
	fmt.Println(VarDeprecated) // want "using deprecated: VarDeprecated by GenDecl ValueSpec"
	var (
		_ = VarDeprecated         // want "using deprecated: VarDeprecated by GenDecl ValueSpec"
		_ = VarDeprecated + "..." // want "using deprecated: VarDeprecated by GenDecl ValueSpec"
		_ = vars1                 // want "using deprecated: vars1 by ValueSpec"
		_ = vars2                 // want "using deprecated: vars1/2/3 by GenDecl ValueSpec"
		_ = vars3                 // want "using deprecated: vars1/2/3 by GenDecl ValueSpec"
		_ = ConstDeprecated       // want "using deprecated: ConstDeprecated by GenDecl ValueSpec"
		_ = consts1               // want "using deprecated: consts1 by ValueSpec"
		_ = consts2               // want "using deprecated: consts 1/2/3 by GenDecl ValueSpec"
		_ = consts3               // want "using deprecated: consts 1/2/3 by GenDecl ValueSpec"
	)
	FuncDeprecated() // want "using deprecated: don't use FuncDeprecated by FuncDecl"
	var (
		_  = StructDeprecated{}  // want "using deprecated: use s3 instead of StructDeprecated, by GenDecl TypeSpec"
		_  = StructDeprecated2{} // want "using deprecated: \\(no comment\\)"
		s  = Struct{}
		_  = struct1{} // want "using deprecated: struct1 by TypeSpec"
		s2 = struct2{} // want "using deprecated: struct 1/2/3 by GenDecl TypeSpec"
		_  = s2.F2     // want "using deprecated: F2 by Field"
		s3 = struct3{} // want "using deprecated: struct 1/2/3 by GenDecl TypeSpec"
	)
	s.StructFun() // want "using deprecated: don't use it"
	s3.fun1()
	s3.fun2() // want "using deprecated: fun2 by FuncDecl"

	var (
		_  InterfaceDeprecated // want "using deprecated: InterfaceDeprecated by GenDecl TypeSpec"
		_  interface1          // want "using deprecated: interface1 by TypeSpec"
		_  interface2          // want "using deprecated: interface 1/2/3 by GenDecl TypeSpec"
		i3 interface3          // want "using deprecated: interface 1/2/3 by GenDecl TypeSpec"
	)
	i3.fun2() // want "using deprecated: interface3 fun2 by Field"

	var (
		_ = pkg.StructDeprecated{}  // want "using deprecated: use s3 instead of StructDeprecated, by GenDecl TypeSpec"
		_ = pkg.StructDeprecated2{} // want "using deprecated: \\(no comment\\)"
	)
}
