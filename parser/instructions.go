package parser

import (
	"fmt"
)

type InstructionType int

const (
	InstructionTypeUnknown InstructionType = iota

	InstructionTypeDefineInt64
	InstructionTypeDefineString
	InstructionTypeDefineFunction
	InstructionTypeDefineGolangFunction
	InstructionTypeDefineVariable
	InstructionTypeDefineFunctionValue
	InstructionTypeCallFunction
	InstructionTypePushToStack
)

func (i InstructionType) String() string {
	switch i {
	case InstructionTypeDefineInt64:
		return "DefineInt64"
	case InstructionTypeDefineString:
		return "DefineString"
	case InstructionTypeDefineFunction:
		return "DefineFunction"
	case InstructionTypeDefineGolangFunction:
		return "DefineGolangFunction"
	case InstructionTypeDefineVariable:
		return "DefineVariable"
	case InstructionTypeDefineFunctionValue:
		return "DefineFunctionValue"
	case InstructionTypeCallFunction:
		return "CallFunction"
	case InstructionTypePushToStack:
		return "PushToStack"
	default:
		return "Unknown"
	}
}

type Instruction interface {
	InstructionType() InstructionType
	DefineInt64() DefineInt64
	DefineString() DefineString
	DefineFunction() DefineFunction
	DefineGolangFunction() DefineGolangFunction
	DefineVariable() DefineVariable
	DefineFunctionValue() DefineFunctionValue
	CallFunction() CallFunction
	PushToStack() PushToStack
}

type defaultInstruction struct{}

func (d defaultInstruction) DefineInt64() DefineInt64 {
	panic("Type is not DefineInt64")
}
func (d defaultInstruction) DefineString() DefineString {
	panic("Type is not DefineString")
}
func (d defaultInstruction) DefineFunction() DefineFunction {
	panic("Type is not DefineFunction")
}
func (d defaultInstruction) DefineGolangFunction() DefineGolangFunction {
	panic("Type is not DefineGolangFunction")
}
func (d defaultInstruction) DefineVariable() DefineVariable {
	panic("Type is not DefineVariable")
}
func (d defaultInstruction) DefineFunctionValue() DefineFunctionValue {
	panic("Type is not DefineFunctionValue")
}
func (d defaultInstruction) CallFunction() CallFunction {
	panic("Type is not CallFunction")
}

func (d defaultInstruction) PushToStack() PushToStack {
	panic("Type is not PushToStack")
}

type DefineInt64 struct {
	defaultInstruction
	Value int64
}

func NewDefineInt64(value int64) DefineInt64 {
	return DefineInt64{Value: value}
}
func (d DefineInt64) String() string {
	return fmt.Sprintf("DefineInt64{Value: %d}", d.Value)
}
func (d DefineInt64) InstructionType() InstructionType {
	return InstructionTypeDefineInt64
}
func (d DefineInt64) DefineInt64() DefineInt64 {
	return d
}

type DefineString struct {
	defaultInstruction
	Value string
}

func NewDefineString(value string) DefineString {
	return DefineString{Value: value}
}
func (d DefineString) String() string {
	return fmt.Sprintf("DefineString{Value: %s}", d.Value)
}
func (d DefineString) InstructionType() InstructionType {
	return InstructionTypeDefineString
}
func (d DefineString) DefineString() DefineString {
	return d
}

type DefineFunction struct {
	defaultInstruction
	Name string
	Node *Node
}

func NewDefineFunction(name string, node *Node) DefineFunction {
	return DefineFunction{Name: name, Node: node}
}
func (d DefineFunction) String() string {
	var nodeString string
	if d.Node == nil {
		nodeString = "<nil>"
	} else {
		nodeString = "<present>"
	}
	return fmt.Sprintf("DefineFunction{Name: %s, Node: %s}", d.Name, nodeString)
}
func (d DefineFunction) InstructionType() InstructionType {
	return InstructionTypeDefineFunction
}
func (d DefineFunction) DefineFunction() DefineFunction {
	return d
}

type DefineGolangFunction struct {
	defaultInstruction
	Content string
}

func NewDefineGolangFunction(content string) DefineGolangFunction {
	return DefineGolangFunction{Content: content}
}
func (d DefineGolangFunction) String() string {
	return fmt.Sprintf("DefineGolangFunction{Content: %s}", d.Content)
}
func (d DefineGolangFunction) InstructionType() InstructionType {
	return InstructionTypeDefineGolangFunction
}
func (d DefineGolangFunction) DefineGolangFunction() DefineGolangFunction {
	return d
}

type DefineVariable struct {
	defaultInstruction
	Name  string
	Value Instruction
}

func NewDefineVariable(name string, value Instruction) DefineVariable {
	return DefineVariable{Name: name, Value: value}
}
func (d DefineVariable) String() string {
	return fmt.Sprintf("DefineVariable{Name: %s, Value: %s}", d.Name, d.Value)
}
func (d DefineVariable) InstructionType() InstructionType {
	return InstructionTypeDefineVariable
}
func (d DefineVariable) DefineVariable() DefineVariable {
	return d
}

type DefineFunctionValue struct {
	defaultInstruction
	Name string
}

func NewDefineFunctionValue(name string) DefineFunctionValue {
	return DefineFunctionValue{Name: name}
}
func (d DefineFunctionValue) String() string {
	return fmt.Sprintf("DefineFunctionValue{Name: %s}", d.Name)
}
func (d DefineFunctionValue) InstructionType() InstructionType {
	return InstructionTypeDefineFunctionValue
}
func (d DefineFunctionValue) DefineFunctionValue() DefineFunctionValue {
	return d
}

type CallFunction struct {
	defaultInstruction
	Name string
}

func NewCallFunction(name string) CallFunction {
	return CallFunction{Name: name}
}
func (c CallFunction) String() string {
	return fmt.Sprintf("CallFunction{Name: %s}", c.Name)
}
func (c CallFunction) InstructionType() InstructionType {
	return InstructionTypeCallFunction
}
func (c CallFunction) CallFunction() CallFunction {
	return c
}

type PushToStack struct {
	defaultInstruction
	Variable string
}

func NewPushToStack(variable string) PushToStack {
	return PushToStack{Variable: variable}
}
func (p PushToStack) String() string {
	return fmt.Sprintf("PushToStack{Variable: %s}", p.Variable)
}
func (p PushToStack) InstructionType() InstructionType {
	return InstructionTypePushToStack
}
func (p PushToStack) PushToStack() PushToStack {
	return p
}
