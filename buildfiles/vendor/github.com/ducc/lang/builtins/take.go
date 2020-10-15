package builtins

import (
	"fmt"

	"github.com/ducc/lang/util"
)

func Take(stack *util.Stack) {
	idx, err := stack.Pop()
	if err != nil {
		panic(fmt.Sprintf("unable to pop b value for take function: %v", err))
	}

	value, err := stack.Get(idx.(int64))
	if err != nil {
		panic(fmt.Sprintf("unable to pop a value for take function: %v", err))
	}

	stack.Push(value)
}
