package compiler

import (
	"fmt"

	"github.com/ducc/lang/parser"
)

type FunctionScope struct {
	buf           *buffer
	variableIndex int
}

func NewFunctionScope(buf *buffer) *FunctionScope {
	return &FunctionScope{buf: buf}
}

func (f *FunctionScope) IncAndGetVariableIndex() int {
	f.variableIndex++
	return f.variableIndex
}

func compileFunction(buf *buffer, instruction parser.Instruction) error {
	if instruction.InstructionType() == parser.InstructionTypeDefineGolangFunction {
		buf.write(instruction.DefineGolangFunction().Content)
		return nil
	}

	definition := instruction.DefineFunction()
	name := definition.Name

	isMain := name == MainFunctionName

	if isMain {
		name = "ducc_main"
	} else if name == ImportFunctionName {
		return nil // we have already handled imports in WriteHeader
	}

	buf.FunctionDefinition(name)
	buf.OpenScope()

	lastNode := definition.Node.Last()
	scope := NewFunctionScope(buf)

	for _, node := range definition.Node.Children {
		compileNodeInstructions(scope, true, node == lastNode, node, lastNode)
	}

	buf.CloseScope()
	return nil
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
	case "print":
		name = "builtins.Print"
	}

	return fmt.Sprintf("scope = %s(scope)", name)
}
