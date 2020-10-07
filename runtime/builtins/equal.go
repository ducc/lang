package builtins

import (
	"fmt"

	"github.com/ducc/lang/util"
)

func Equal(stack *util.Stack) {
	b, err := stack.Pop()
	if err != nil {
		panic(fmt.Sprintf("unable to pop b value for equal function: %v", err))
	}

	a, err := stack.Pop()
	if err != nil {
		panic(fmt.Sprintf("unable to pop a value for equal function: %v", err))
	}

	result := a.(int64) == b.(int64)
	stack.Push(result)
}
