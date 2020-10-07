package lang

import (
	"fmt"
	"regexp"
	"strings"
)

func ParseFunctionDefinitions(input string) ([]DefineFunction, error) {
	splits := splitIntoFunctionDefinitions(input)
	cleanSplits := cleanFunctionDefinitions(splits)

	definedFunctions := make([]DefineFunction, 0, len(cleanSplits))
	for _, definition := range cleanSplits {
		definedFunction, err := parseFunctionDefinition(definition)
		if err != nil {
			return nil, fmt.Errorf("parsing function definition: %w", err)
		}

		definedFunctions = append(definedFunctions, definedFunction)
	}

	return definedFunctions, nil
}

func splitIntoFunctionDefinitions(input string) []string {
	var r = regexp.MustCompile(`(?m)([a-zA-Z0-9]+\s=\n)(?:(?:\|\s.+)\n)+`)
	return r.FindAllString(input, -1)
}

func cleanFunctionDefinitions(definitions []string) []string {
	output := make([]string, len(definitions))
	for i, definition := range definitions {
		output[i] = cleanFunctionDefinition(definition)
	}
	return output
}

func cleanFunctionDefinition(input string) string {
	lines := strings.Split(input, "\n")
	instructions := make([]string, 0)
	for _, line := range lines {
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "#") {
			continue
		}
		instructions = append(instructions, line)
	}
	return strings.Join(instructions, "\n")
}
