package runtime

import (
	"github.com/ducc/lang/runtime/builtins"
)

func (r *FunctionRegistry) registerBuiltinFunctions() {
	register := func(name string, step Step) {
		r.functions[name] = NewFunction([]Step{step})
	}

	register("add", builtins.Add)
	register("sub", builtins.Sub)
	register("equal", builtins.Equal)
	register("more", builtins.More)
	register("take", builtins.Take)
	register("drop", builtins.Drop)
	register("push", builtins.Push)
	register("print", builtins.Print)
	register("printascii", builtins.PrintASCII)
}
