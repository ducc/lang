package parser

import (
	"fmt"
	"strconv"
	"strings"
)

func splitIntoFunctions(input string) []string {
	splits := make([]string, 0)

	for _, split := range strings.Split(input, "\n\n") {
		lines := make([]string, 0)
		for _, line := range strings.Split(split, "\n") {
			if strings.HasPrefix(line, "#") {
				continue
			}

			if strings.Contains(line, "#") {
				idx := strings.Index(line, "#")
				line = line[:idx]
			}

			line := strings.TrimSpace(line)
			if line == "" || line == "\n" {
				continue
			}

			lines = append(lines, line)
		}

		if len(lines) == 0 {
			continue
		}

		splits = append(splits, strings.Join(lines, "\n"))
	}

	return splits
}

func splitInstructionParts(line string) []string {
	parts := make([]string, 0)

	var part string
	stringOpen := false
	for _, char := range line {
		if char == '"' {
			stringOpen = !stringOpen
		} else if char == ' ' && part != "" && !stringOpen {
			parts = append(parts, part)
			part = ""
			continue
		}

		part += string(char)
	}

	if part != "" {
		parts = append(parts, part)
	}

	return parts
}

func parseFunctionIntoTree(split string) *Node {
	lines := strings.Split(split, "\n")

	functionName := lines[0]
	if strings.HasPrefix(functionName, "func ") {
		// its a go function
		return nil
	}

	root := newNode([]string{functionName})
	target := root
	var lastNum int

	instructions := parseFunctionInstructions(lines)

	for _, instruction := range instructions {
		if lastNum > instruction.pipes {
			diff := lastNum - instruction.pipes
			target = target.Back(diff + 1)
		}

		n := newNode(instruction.line)
		target.AddChild(n)
		target = n

		lastNum = instruction.pipes
	}

	return root
}

func functionTreeToInstructionTree(root *Node) DefineFunction {
	for _, node := range root.Queue().Inner {
		if node == root {
			node.ParsedInstruction = NewDefineFunction(node.Instruction[0], node)
			continue
		}

		if len(node.Instruction) == 2 {
			// variable assignment
			variableName := node.Instruction[0]
			variableValue := parseInstruction(node.Instruction[1])
			node.ParsedInstruction = NewDefineVariable(variableName, variableValue)
		} else {
			instruction := node.Instruction[0]
			node.ParsedInstruction = parseInstruction(instruction)
		}
	}

	return root.ParsedInstruction.DefineFunction()
}

func parseInstruction(instruction string) Instruction {
	if strings.HasPrefix(instruction, "\"") && strings.HasSuffix(instruction, "\"") {
		// define string variable
		return NewDefineString(instruction[1 : len(instruction)-1])
	} else if intValue, err := strconv.ParseInt(instruction, 10, 64); err == nil {
		// define int variable
		return NewDefineInt64(intValue)
	} else if strings.HasPrefix(instruction, "$") {
		// using function call as a value
		return NewDefineFunctionValue(instruction[1:])
	} else if strings.HasPrefix(instruction, "!") {
		// bringing value onto stack
		return NewPushToStack(instruction[1:])
	} else {
		// function call
		return NewCallFunction(instruction)
	}
}

func ParseFunctions(input string) ([]Instruction, error) {
	definedFunctions := make([]Instruction, 0)

	for _, split := range splitIntoFunctions(input) {
		tree := parseFunctionIntoTree(split)
		if tree == nil {
			// its a go function
			definedFunctions = append(definedFunctions, NewDefineGolangFunction(split))
			continue
		}

		definedFunction := functionTreeToInstructionTree(tree)
		fmt.Println(definedFunction.Node.Queue().String())
		definedFunctions = append(definedFunctions, definedFunction)
	}

	return definedFunctions, nil
}
