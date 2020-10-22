package compiler

import (
	"fmt"

	"github.com/ducc/lang/parser"
)

type InstructionScope struct {
	buf *buffer
}

func NewInstructionScope(functionScope *FunctionScope) *InstructionScope {
	return &InstructionScope{buf: functionScope.buf}
}

func (s *InstructionScope) Open() {
	s.buf.OpenScope()
	s.buf.CloneScope()
}

func (s *InstructionScope) Close() {
	s.buf.CloseScope()
}

func compileNodeInstructions(functionScope *FunctionScope, initial, returnScope bool, node, lastNode *parser.Node) {
	scope := NewInstructionScope(functionScope)
	scope.Open()

	compileInstruction(functionScope.buf, functionScope.IncAndGetVariableIndex(), node.ParsedInstruction)

	for _, node := range node.Children {
		var returnScope bool
		if node == lastNode {
			returnScope = true
		}

		compileNodeInstructions(functionScope, false, returnScope, node, lastNode)
	}

	if returnScope {
		scope.buf.ReturnScope()
	}

	scope.Close()
}

func compileInstruction(buf *buffer, idx int, instruction parser.Instruction) {
	switch instruction.InstructionType() {
	case parser.InstructionTypeDefineInt64:
		compileDefineInt64(buf, idx, instruction.DefineInt64())
	case parser.InstructionTypeDefineString:
		compileDefineString(buf, idx, instruction.DefineString())
	case parser.InstructionTypeDefineVariable:
		compileDefineVariable(buf, idx, instruction.DefineVariable())
	case parser.InstructionTypeDefineFunctionValue:
		compileDefineFunctionValue(buf, idx, instruction.DefineFunctionValue())
	case parser.InstructionTypeCallFunction:
		compileCallFunction(buf, idx, instruction.CallFunction())
	case parser.InstructionTypePushToStack:
		compilePushToStack(buf, idx, instruction.PushToStack())
	}
}

func compileDefineInt64(buf *buffer, idx int, instruction parser.DefineInt64) {
	buf.PushStack(fmt.Sprintf("int64(%d)", instruction.Value))
}

func compileDefineString(buf *buffer, idx int, instruction parser.DefineString) {
	buf.PushStack(fmt.Sprintf("\"%s\"", instruction.Value))
}

func compileDefineVariable(buf *buffer, idx int, instruction parser.DefineVariable) {
	var value string

	switch instruction.Value.InstructionType() {
	case parser.InstructionTypeDefineInt64:
		value = fmt.Sprintf("int64(%d)", instruction.Value.DefineInt64().Value)
	case parser.InstructionTypeDefineString:
		value = "\"" + instruction.Value.DefineString().Value + "\""
	case parser.InstructionTypeCallFunction:
		buf.write(getFunctionCall(instruction.Value.CallFunction().Name))
		buf.write(fmt.Sprintf("value%d, scope := scope.PopStack()", idx))
		value = fmt.Sprintf("value%d", idx)
	case parser.InstructionTypeDefineFunctionValue:
		value = instruction.Name
	}

	buf.SetVar(instruction.Name, value)
}

func compileDefineFunctionValue(buf *buffer, idx int, instruction parser.DefineFunctionValue) {
	buf.CallFunctionVar(instruction.Name)
}

func compileCallFunction(buf *buffer, idx int, instruction parser.CallFunction) {
	buf.ZapDebugCall(instruction.Name)
	buf.write(getFunctionCall(instruction.Name))
}

func compilePushToStack(buf *buffer, idx int, instruction parser.PushToStack) {
	buf.PushVarToStack(instruction.Variable)
}
