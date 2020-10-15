package parser

type functionInstruction struct {
	line  []string
	pipes int
}

func newFunctionInstruction(line []string, pipes int) functionInstruction {
	return functionInstruction{line: line, pipes: pipes}
}

func parseFunctionInstructions(lines []string) []functionInstruction {
	instructions := make([]functionInstruction, 0)

	for _, line := range lines[1:] {
		pipes := 0
		for i, char := range line {
			if char != '|' && char != ' ' {
				line = line[i:]
				break
			}
			if char == '|' {
				pipes++
			}
		}

		instructions = append(instructions, newFunctionInstruction(splitInstructionParts(line), pipes))
	}

	return instructions
}
