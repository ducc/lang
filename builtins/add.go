package builtins

import (
	"github.com/ducc/lang/util"
)

func Add(scope *util.Scope) *util.Scope {
	b, scope := scope.PopStack()
	a, scope := scope.PopStack()

	result := a.(int64) + b.(int64)
	return scope.PushStack(result)
}
