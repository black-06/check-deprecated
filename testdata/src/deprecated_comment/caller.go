package deprecated_comment

import (
	"fmt"
	"io/ioutil"
)

func Caller() {
	ioutil.ReadAll(nil)
	fmt.Println(VarDeprecated)
	var (
		_ = VarDeprecated
		_ = VarDeprecated + "..."
		_ = vars1
		_ = vars2
		_ = vars3
		_ = ConstDeprecated
		_ = consts1
		_ = consts2
		_ = consts3
	)
	FuncDeprecated()
	var (
		_  = StructDeprecated{}
		_  = StructDeprecated2{}
		s  = Struct{}
		_  = struct1{}
		s2 = struct2{}
		_  = s2.F2
		s3 = struct3{}
	)
	s.StructFun()
	s3.fun1()
	s3.fun2()

	var (
		_  InterfaceDeprecated
		_  interface1
		_  interface2
		i3 interface3
	)
	i3.fun2()
}
