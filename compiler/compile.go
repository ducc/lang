package compiler

import (
	"fmt"
	"io"

	"github.com/ducc/lang/parser"
	"go.uber.org/zap"
)

type Compiler struct {
	logger *zap.SugaredLogger
	debug  bool
}

func New(logger *zap.SugaredLogger, debug bool) *Compiler {
	return &Compiler{logger: logger, debug: debug}
}

func (c *Compiler) Compile(definitions []parser.Instruction) (io.Reader, error) {
	buf := c.newBuffer()

	imports := make([]string, 0)
	if definitions[0].DefineFunction().Name == "import" {
		for i, node := range definitions[0].DefineFunction().Node.Queue().Inner {
			if i == 0 {
				continue
			}
			imports = append(imports, node.ParsedInstruction.DefineString().Value)
		}
	}

	WriteHeader(buf, imports)

	for _, definition := range definitions {
		if err := compileFunction(buf, definition); err != nil {
			return nil, err
		}
	}

	return buf, nil
}

func compileFunction(buf *buffer, instruction parser.Instruction) error {
	if instruction.InstructionType() == parser.InstructionTypeDefineGolangFunction {
		buf.write(instruction.DefineGolangFunction().Content)
		return nil
	}

	definition := instruction.DefineFunction()
	name := definition.Name
	var functionFooter string
	var isMain bool

	if name == "main" {
		buf.writef("func %s() {\nlogger, _ := zap.NewDevelopment()\nzap.ReplaceGlobals(logger)\ndefer logger.Sync()\nscope := util.NewScope()", name)
		functionFooter = "}\n"
		isMain = true
	} else if name == "import" {
		buf.write("// TODO IMPORTS")
		functionFooter = "}\n"
		return nil
	} else {
		buf.writef("func %s(scope *util.Scope) *util.Scope {", name)
		functionFooter = "}\n"
	}

	lastNode := findLastNode(definition.Node)
	fmt.Println("lastNode", lastNode)

	for _, node := range definition.Node.Children {
		compileNodeInstructions(buf, true, isMain, node == lastNode, node, lastNode)
	}

	buf.write(functionFooter)
	return nil
}

func findLastNode(node *parser.Node) *parser.Node {
	if len(node.Children) == 0 {
		fmt.Println("LAST NODE =", node.ParsedInstruction)
		return node
	}

	lastChild := node.Children[len(node.Children)-1]
	fmt.Println("NEXT NODE =", lastChild.ParsedInstruction)
	return findLastNode(lastChild)
}

var idx = 0

func compileNodeInstructions(buf *buffer, initial, isMain, returnScope bool, node, lastNode *parser.Node) {
	// if !initial {
		buf.write("{")
		buf.write("scope := scope.Clone()")
	// }

	idx++
	compileInstruction(buf, idx, node.ParsedInstruction)

	for _, node := range node.Children {
		var returnScope bool
		if node == lastNode && !isMain {
			returnScope = true
		}

		buf.writef("// Node=%s Last=%s Equal=%t Main=%t", node.ParsedInstruction, lastNode.ParsedInstruction, node == lastNode, isMain)
		compileNodeInstructions(buf, false, isMain, returnScope, node, lastNode)
	}

	if returnScope {
		buf.write("return scope")
	}

	// if !initial {
		buf.write("}")
	// }
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
	buf.writef("scope = scope.PushStack(int64(%d))", instruction.Value)
}

func compileDefineString(buf *buffer, idx int, instruction parser.DefineString) {
	buf.writef("scope = scope.PushStack(\"%s\")", instruction.Value)
}

func compileDefineVariable(buf *buffer, idx int, instruction parser.DefineVariable) {
	var before string
	var value string

	switch instruction.Value.InstructionType() {
	case parser.InstructionTypeDefineInt64:
		value = fmt.Sprintf("int64(%d)", instruction.Value.DefineInt64().Value)
	case parser.InstructionTypeDefineString:
		value = "\"" + instruction.Value.DefineString().Value + "\""
	case parser.InstructionTypeCallFunction:
		before = getFunctionCall(instruction.Value.CallFunction().Name) + fmt.Sprintf("\nvalue%d, scope := scope.PopStack()", idx)
		value = fmt.Sprintf("value%d", idx)
	case parser.InstructionTypeDefineFunctionValue:
		value = instruction.Name
	}

	if before != "" {
		buf.write(before)
	}
	buf.writef("scope = scope.SetVar(\"%s\", %s)", instruction.Name, value)
}

func compileDefineFunctionValue(buf *buffer, idx int, instruction parser.DefineFunctionValue) {
	buf.write("/////// MEME ////////")
	buf.writef("scope = scope.GetVar(\"%s\").(func(*util.Scope) *util.Scope)(scope)", instruction.Name)
}

func compileCallFunction(buf *buffer, idx int, instruction parser.CallFunction) {
	buf.debugf("zap.S().Debug(\"call %20s \", scope.String())", instruction.Name)
	buf.write(getFunctionCall(instruction.Name))
}

func compilePushToStack(buf *buffer, idx int, instruction parser.PushToStack) {
	buf.writef("scope = scope.PushStack(scope.GetVar(\"%s\"))", instruction.Variable)
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
	case "pop":
		name = "builtins.Pop"
	}

	return fmt.Sprintf("scope = %s(scope)", name)
}
