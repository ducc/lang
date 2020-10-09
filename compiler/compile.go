package compiler

import (
	"fmt"
	"io"

	"github.com/ducc/lang/parser"
	"go.uber.org/zap"
)

type Compiler struct {
	logger *zap.SugaredLogger
}

func New(logger *zap.SugaredLogger) *Compiler {
	return &Compiler{logger: logger}
}

func (c *Compiler) Compile(definitions []parser.DefineFunction) (io.Reader, error) {
	buf := c.newBuffer()
	buf.write(header)

	for _, definition := range definitions {
		if err := compileFunction(buf, definition); err != nil {
			return nil, err
		}
	}

	return buf, nil
}

func compileFunction(buf *buffer, definition parser.DefineFunction) error {
	name := definition.Name

	if name == "main" {
		buf.write("func %s() {\nstack := util.NewStack()", name)
	} else {
		buf.write("func %s(stack *util.Stack) {", name)
	}

	for i, instruction := range definition.Instructions {
		instruction := instruction

		switch instruction.InstructionType() {
		case parser.InstructionTypeDefineInt64:
			compileDefineInt64(buf, instruction.DefineInt64())
		case parser.InstructionTypeCallFunction:
			compileCallFunction(buf, instruction.CallFunction())
		case parser.InstructionTypeConditional:
			compileConditional(buf, instruction.Conditional(), i)
		default:
			return fmt.Errorf("instruction type cannot be inside a function: %s", instruction)
		}
	}

	buf.write("}\n")
	return nil
}

func compileDefineInt64(buf *buffer, instruction parser.DefineInt64) {
	buf.write("stack.Push(int64(%d))", instruction.Value)
}

func compileCallFunction(buf *buffer, instruction parser.CallFunction) {
	buf.write(getFunctionCall(instruction.Name))
}

func compileConditional(buf *buffer, instruction parser.Conditional, tempID int) {
	buf.write(getFunctionCall(instruction.ConditionFunction.Name))
	buf.write("value%d, _ := stack.Pop()", tempID)
	buf.write("if value%d.(bool) {", tempID)
	buf.write(getFunctionCall(instruction.TrueFunction.Name))
	buf.write("} else {")
	buf.write(getFunctionCall(instruction.FalseFunction.Name))
	buf.write("}")
}

func getFunctionCall(name string) string {
	switch name {
	case "add":
		name = "builtins.Add"
	case "sub":
		name = "builtins.Sub"
	case "equal":
		name = "builtins.Equal"
	case "more":
		name = "builtins.More"
	case "take":
		name = "builtins.Take"
	case "drop":
		name = "builtins.Drop"
	case "push":
		name = "builtins.Push"
	case "print":
		name = "builtins.Print"
	case "printascii":
		name = "builtins.PrintASCII"
	}

	return fmt.Sprintf("%s(stack)", name)
}
