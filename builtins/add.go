package builtins

import (
	"fmt"

	"github.com/ducc/lang/util"
)

func Add(stack *util.Stack) {
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
