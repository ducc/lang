package lang

import (
	"fmt"
	"strings"
)

type InstructionType int

const (
	InstructionTypeUnknown InstructionType = iota

	InstructionTypeDefineInt64
	InstructionTypeDefineFunction
	InstructionTypeCallFunction
	InstructionTypeConditional
)

func (i InstructionType) String() string {
	switch i {
	case InstructionTypeDefineInt64:
		return "DefineInt64"
	case InstructionTypeDefineFunction:
		return "DefineFunction"
	case InstructionTypeCallFunction:
		return "CallFunction"
	case InstructionTypeConditional:
		return "Conditional"
	default:
		return "Unknown"
	}
}

type Instruction interface {
	InstructionType() InstructionType
	DefineInt64() DefineInt64
	DefineFunction() DefineFunction
	CallFunction() CallFunction
	Conditional() Conditional
}

type defaultInstruction struct{}

func (d defaultInstruction) DefineInt64() DefineInt64 {
	panic("Type is not DefineInt64")
}
func (d defaultInstruction) DefineFunction() DefineFunction {
	panic("Type is not DefineFunction")
}
func (d defaultInstruction) CallFunction() CallFunction {
	panic("Type is not CallFunction")
}
func (c defaultInstruction) Conditional() Conditional {
	panic("Type is not Conditional")
}

type DefineInt64 struct {
	defaultInstruction
	Value int64
}

func NewDefineInt64(value int64) DefineInt64 {
	return DefineInt64{Value: value}
}
func (d DefineInt64) String() string {
	return fmt.Sprintf("DefineInt64{Type: %s, Value: %d}", d.InstructionType(), d.Value)
}
func (d DefineInt64) InstructionType() InstructionType {
	return InstructionTypeDefineInt64
}
func (d DefineInt64) DefineInt64() DefineInt64 {
	return d
}

type DefineFunction struct {
	defaultInstruction
	Name         string
	Instructions []Instruction
}

func NewDefineFunction(name string, instructions []Instruction) DefineFunction {
	return DefineFunction{Name: name, Instructions: instructions}
}
func (d DefineFunction) String() string {
	instructions := make([]string, len(d.Instructions))
	for i, instruction := range d.Instructions {
		instructions[i] = fmt.Sprint(instruction)
	}
	return fmt.Sprintf("DefineFunction{Type: %s, Name: %s, Instructions: [%s]}", d.InstructionType(), d.Name, strings.Join(instructions, ", "))
}
func (d DefineFunction) InstructionType() InstructionType {
	return InstructionTypeDefineFunction
}
func (d DefineFunction) DefineFunction() DefineFunction {
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
	return fmt.Sprintf("CallFunction{Type: %s, Name: %s}", c.InstructionType(), c.Name)
}
func (c CallFunction) InstructionType() InstructionType {
	return InstructionTypeCallFunction
}
func (c CallFunction) CallFunction() CallFunction {
	return c
}

type Conditional struct {
	defaultInstruction
	ConditionFunction CallFunction
	TrueFunction      CallFunction
	FalseFunction     CallFunction
}

func NewConditional(conditionFunction, trueFunction, falseFunction CallFunction) Conditional {
	return Conditional{
		ConditionFunction: conditionFunction,
		TrueFunction:      trueFunction,
		FalseFunction:     falseFunction,
	}
}
func (c Conditional) String() string {
	return fmt.Sprintf("Conditional{Type: %s, ConditionalFunction: %s, TrueFunction: %s, FalseFunction: %s}", c.InstructionType(), c.ConditionFunction, c.TrueFunction, c.FalseFunction)
}
func (c Conditional) InstructionType() InstructionType {
	return InstructionTypeConditional
}
func (c Conditional) Conditional() Conditional {
	return c
}
