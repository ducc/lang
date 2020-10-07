package builtins

import (
	"fmt"

	"github.com/ducc/lang/util"
)

func PrintASCII(stack *util.Stack) {
	chars, err := stack.Pop()
	if err != nil {
		panic(fmt.Sprintf("unable to pop chars value for printascii function: %v", err))
	}

	var text string

	for i := int64(0); i < chars.(int64); i++ {
		char, err := stack.Pop()
		if err != nil {
			panic(fmt.Sprintf("unable to pop char value for printascii function: %v", err))
		}

		text += fmt.Sprintf("%c", rune(char.(int64)))
	}

	fmt.Println(reverse(text))
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
