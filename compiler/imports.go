package compiler

import "github.com/ducc/lang/parser"

const ImportFunctionName = "import"

func getImports(definitions []parser.Instruction) []string {
	imports := make([]string, 0)
	for _, definition := range definitions {
		// imports must be defined in a function
		if definition.InstructionType() != parser.InstructionTypeDefineFunction {
			continue
		}

		definition := definition.DefineFunction()

		if definition.Name == ImportFunctionName {
			for i, node := range definition.Node.Queue().Inner {
				if i == 0 {
					// we want to skip the function name definition
					continue
				}

				imports = append(imports, node.ParsedInstruction.DefineString().Value)
			}
		}
	}
	return imports
}
