package runtime

import "fmt"

func (r *FunctionRegistry) registerBuiltinFunctions() {
	register := func(name string, step Step) {
		r.functions[name] = NewFunction([]Step{step})
	}

	register("add", addBuiltin)
	register("print", printPop)
}

func addBuiltin(stack *Stack) {
	b, err := stack.Pop()
	if err != nil {
		panic(fmt.Sprintf("unable to pop b value for add function: %v", err))
	}

	a, err := stack.Pop()
	if err != nil {
		panic(fmt.Sprintf("unable to pop a value for add function: %v", err))
	}

	result := a.(int64) + b.(int64)
	stack.Push(result)
}

func printPop(stack *Stack) {
	v, err := stack.Pop()
	if err != nil {
		panic(fmt.Sprintf("unable to pop for print function: %v", err))
	}

	fmt.Println(v)
}
