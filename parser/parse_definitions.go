package parser

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrEmptyFunctionDefinition = errors.New("empty function definition")

func parseFunctionDefinition(input string) (DefineFunction, error) {
	var name string
	instructions := make([]Instruction, 0)

	for i, line := range strings.Split(input, "\n") {
		if i == 0 {
			name = line[:strings.Index(line, "=")-1]
			continue
		} else if strings.HasPrefix(line, "|") {
			line = line[2:]
		}

		instruction, err := parseFunctionDefinitionInstruction(line)
		if err != nil {
			return DefineFunction{}, fmt.Errorf("parsing function definition instruction: %w", err)
		}

		instructions = append(instructions, instruction)
	}

	return DefineFunction{
		Name:         name,
		Instructions: instructions,
	}, nil
}

func parseFunctionDefinitionInstruction(line string) (Instruction, error) {
	if strings.HasPrefix(line, "i64(") {
		instruction, err := parseDefineInt64(line)
		if err != nil {
			return nil, fmt.Errorf("parsing function definition instruction as DefineInt64: %w", err)
		}
		return instruction, nil
	} else if strings.Contains(line, "?") {
		return parseConditional(line), nil
	} else {
		return parseCallFunction(line), nil
	}
}

func parseDefineInt64(line string) (DefineInt64, error) {
	content := line[4:strings.LastIndex(line, ")")]
	intVal, err := strconv.ParseInt(content, 10, 64)
	if err != nil {
		return DefineInt64{}, fmt.Errorf("parsing DefineInt64 content: %w", err)
	}
	return DefineInt64{Value: intVal}, nil
}

func parseConditional(line string) Conditional {
	conditionFunction := line[:strings.Index(line, "?")-1]
	trueFunction := line[strings.Index(line, "?")+2 : strings.Index(line, ":")-1]
	falseFunction := line[strings.Index(line, ":")+2:]

	return Conditional{
		ConditionFunction: parseCallFunction(conditionFunction),
		TrueFunction:      parseCallFunction(trueFunction),
		FalseFunction:     parseCallFunction(falseFunction),
	}
}

func parseCallFunction(line string) CallFunction {
	return CallFunction{Name: line}
}
