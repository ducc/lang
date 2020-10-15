package builtins

import (
	"fmt"

	"github.com/ducc/lang/util"
)

func Print(scope *util.Scope) *util.Scope {
	v, scope := scope.PopStack()
	fmt.Println(v)
	return scope
}
