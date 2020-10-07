package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input := `

D =
| i64(68)
| push

E =
| i64(69)
| push

H =
| i64(72)
| push

L =
| i64(76)
| push

O =
| i64(79)
| push

R =
| i64(82)
| push

W =
| i64(87)
| push

space =
| i64(8194)
| push

hello =
| H
| E
| L
| L
| O

world =
| W
| O
| R
| L
| D

main =
| hello
| space
| world
| printascii

`

	run(input)
}

func parseRawBlocks(input string) []string {
	var blockRegex = regexp.MustCompile(`(?m)([a-zA-Z0-9]+\s=\n)(?:(?:\|\s.+)\n)+`)
	blockMatches := blockRegex.FindAllString(input, -1)
	for i, blockMatch := range blockMatches {
		fmt.Printf("=== block match %d:\n%s\n", i, blockMatch)
	}
	return blockMatches
}

func parseBlocks(rawBlocks []string) []Block {
	blocks := make([]Block, 0)

	for _, blockMatch := range rawBlocks {
		var functionName string
		blockInstructions := make([]interface{}, 0)

		for i, line := range strings.Split(cleanInput(blockMatch), "\n") {
			if i == 0 {
				functionName = line[:strings.Index(line, "=")-1]
				continue
			} else if strings.HasPrefix(line, "|") {
				line = line[2:]
			}

			var value interface{}
			if strings.HasPrefix(line, "i64(") {
				content := line[4:strings.LastIndex(line, ")")]
				intVal, _ := strconv.ParseInt(content, 10, 64)
				value = intValue(intVal)
			} else if strings.Contains(line, "?") {
				calledFunc := line[:strings.Index(line, "?")-1]
				trueFunc := line[strings.Index(line, "?")+2 : strings.Index(line, ":")-1]

				if strings.Contains(line, ":") {
					falseFunc := line[strings.Index(line, ":")+2:]
					value = functionCallWithConditionalValue(calledFunc, trueFunc, falseFunc)
				} else {
					value = functionCallWithConditionalValue(calledFunc, trueFunc, "")
				}
			} else {
				value = functionCallValue(line)
			}

			blockInstructions = append(blockInstructions, value)
		}

		for _, blockInstruction := range blockInstructions {
			fmt.Println(blockInstruction)
		}

		blocks = append(blocks, newBlock(functionName, blockInstructions))
	}

	return blocks
}

func registerBuiltinFunctions(output map[string]Function, programOutput *[]string) {
	output["add"] = func(v []interface{}) interface{} {
		fmt.Println("add", v)
		return Int64Value{Value: v[0].(Int64Value).Value + v[1].(Int64Value).Value}
	}
	output["sub"] = func(v []interface{}) interface{} {
		fmt.Println("sub", v)
		return Int64Value{Value: v[0].(Int64Value).Value - v[1].(Int64Value).Value}
	}
	output["equal"] = func(v []interface{}) interface{} {
		fmt.Println("equal", v)
		return BoolValue{Value: v[0] == v[1]}
	}
	output["more"] = func(v []interface{}) interface{} {
		fmt.Println("more", v)
		return BoolValue{Value: v[0].(Int64Value).Value > v[1].(Int64Value).Value}
	}
	output["take"] = func(v []interface{}) interface{} {
		fmt.Println("take", v)
		index := v[len(v)-1].(Int64Value).Value
		return v[index]
	}
	output["drop"] = func(v []interface{}) interface{} {
		fmt.Println("drop")
		return None{}
	}
	output["push"] = func(v []interface{}) interface{} {
		fmt.Println("push", v)
		output := make([]interface{}, 0)
		for _, value := range v {
			switch value := value.(type) {
			case []interface{}:
				for _, inner := range value {
					output = append(output, inner)
				}
			default:
				output = append(output, value)
			}
		}
		return output
	}
	output["print"] = func(v []interface{}) interface{} {
		values := make([]string, 0)
		for _, v := range v {
			switch v := v.(type) {
			case Int64Value:
				values = append(values, fmt.Sprint(v.Value))
			case BoolValue:
				var text string
				if v.Value {
					text = "yes"
				} else {
					text = "no"
				}
				values = append(values, text)
			default:
				panic(fmt.Sprintf("unknown type: %s", v))
			}
		}
		*programOutput = append(*programOutput, strings.Join(values, " "))
		return nil
	}
	output["printascii"] = func(v []interface{}) interface{} {
		chars := v[0]
		text := ""
		for _, char := range chars.([]interface{}) {
			text += fmt.Sprintf("%c", rune(char.(Int64Value).Value))
		}
		*programOutput = append(*programOutput, text)
		fmt.Println("printascii", text)
		return nil
	}
}

func run(input string) {
	programOutput := make([]string, 0)

	functions := make(map[string]Function)
	registerBuiltinFunctions(functions, &programOutput)

	rawBlocks := parseRawBlocks(input)
	blocks := parseBlocks(rawBlocks)

	fmt.Println("============ registering functions =============")

	names := make([]string, 0)
	for _, block := range blocks {
		block := block
		functions[block.FunctionName] = func(input []interface{}) interface{} {
			return runInstructions(functions, block.Instructions, input)
		}
		names = append(names, block.FunctionName)
	}

	fmt.Printf("registed: %v\n", names)

	fmt.Println("========== running ==========")

	_ = functions["main"]([]interface{}{})

	fmt.Printf("========= output =========\n%s\n", strings.Join(programOutput, "\n"))
}

func runInstructions(functions map[string]Function, instructions []interface{}, input []interface{}) interface{} {
	values := input

	for _, instruction := range instructions {
		switch instruction := instruction.(type) {
		case Int64Value:
			values = append(values, instruction)
		case FunctionCallValue:
			funcs := functions
			fmt.Println("funcs", funcs)
			instr := instruction
			fmt.Println("instr", instr)
			instr_fun := instruction.Function
			fmt.Println("instr_fun", instr_fun)
			fun := funcs[instruction.Function]
			fmt.Println("fun", fun)
			input_values := values
			fmt.Println("input_values", input_values)
			result := fun(input_values)
			values = make([]interface{}, 0)
			values = append(values, result)
		case FunctionCallWithConditionalValue:
			result := functions[instruction.Func.Function](values).(BoolValue).Value
			if result {
				result := functions[instruction.True.Function](values)
				values = make([]interface{}, 0)
				values = append(values, result)
			} else {
				result := functions[instruction.False.Function](values)
				values = make([]interface{}, 0)
				values = append(values, result)
			}
		}
	}

	if len(values) == 0 {
		return None{}
	} else {
		return values[0]
	}
}

func cleanInput(input string) string {
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

type Int64Value struct {
	Value int64
}

func (i Int64Value) String() string {
	return fmt.Sprintf("IntValue{Value: %d}", i.Value)
}
func intValue(v int64) Int64Value {
	return Int64Value{Value: v}
}

type BoolValue struct {
	Value bool
}

func (i BoolValue) String() string {
	return fmt.Sprintf("BoolValue{Value: %t}", i.Value)
}
func boolValue(v bool) BoolValue {
	return BoolValue{Value: v}
}

type FunctionCallValue struct {
	Function string
}

func (i FunctionCallValue) String() string {
	return fmt.Sprintf("FunctionCallValue{Function: %s}", i.Function)
}
func functionCallValue(f string) FunctionCallValue {
	return FunctionCallValue{Function: f}
}

type FunctionCallWithConditionalValue struct {
	Func  FunctionCallValue
	True  FunctionCallValue
	False FunctionCallValue
}

func (i FunctionCallWithConditionalValue) String() string {
	return fmt.Sprintf("FunctionCallWithCondiitionalValue{Func: %s, True: %s, False: %s}", i.Func, i.True, i.False)
}
func functionCallWithConditionalValue(f, trueF, falseF string) FunctionCallWithConditionalValue {
	return FunctionCallWithConditionalValue{
		Func:  functionCallValue(f),
		True:  functionCallValue(trueF),
		False: functionCallValue(falseF),
	}
}

type Function func(v []interface{}) interface{}

type Block struct {
	FunctionName string
	Instructions []interface{}
}

func newBlock(name string, instructions []interface{}) Block {
	return Block{FunctionName: name, Instructions: instructions}
}

type None struct{}
