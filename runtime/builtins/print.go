package builtins

import (
	"fmt"

	"github.com/ducc/lang/util"
)

func Print(stack *util.Stack) {
	v, err := stack.Pop()
	if err != nil {
		panic(fmt.Sprintf("unable to pop for print function: %v", err))
	}

	fmt.Println(v)
}
