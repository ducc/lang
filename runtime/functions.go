package runtime

import (
	"fmt"

	"github.com/ducc/lang/lang"
	"go.uber.org/zap"
)

type Step func(stack *Stack)

type Function struct {
	steps []Step
}

func NewFunction(steps []Step) *Function {
	return &Function{steps: steps}
}

func (f *Function) Invoke(stack *Stack) {
	for _, step := range f.steps {
		step(stack)
	}
}

type FunctionRegistry struct {
	logger    *zap.SugaredLogger
	functions map[string]*Function
}

func NewFunctionRegistry(logger *zap.SugaredLogger) *FunctionRegistry {
	registry := &FunctionRegistry{
		logger:    logger,
		functions: make(map[string]*Function),
	}
	registry.registerBuiltinFunctions()
	return registry
}

func (r *FunctionRegistry) Function(name string) (*Function, error) {
	function, ok := r.functions[name]
	if !ok {
		return nil, fmt.Errorf("function '%s' does not exist", name)
	}

	return function, nil
}

func (r *FunctionRegistry) Register(definition lang.DefineFunction) error {
	name := definition.Name
	r.logger.Debugf("registering '%s'", name)

	steps := make([]Step, 0)

	for _, instruction := range definition.Instructions {
		instruction := instruction

		switch instruction.InstructionType() {
		case lang.InstructionTypeDefineInt64:
			steps = append(steps, func(stack *Stack) {
				r.logger.Debugf("%s %s", instruction, stack)
				stack.Push(instruction.DefineInt64().Value) // todo should we be pushing the actual value or a wrapper to the stack?
			})
		case lang.InstructionTypeCallFunction:
			function, err := r.Function(instruction.CallFunction().Name)
			if err != nil {
				return err
			}

			steps = append(steps, func(stack *Stack) {
				r.logger.Debugf("%s %s", instruction, stack)
				function.Invoke(stack)
			})
		case lang.InstructionTypeConditional:
			instruction := instruction.Conditional()

			conditionFunction, err := r.Function(instruction.ConditionFunction.Name)
			if err != nil {
				return fmt.Errorf("getting conditional condition function: %w", err)
			}

			trueFunction, err := r.Function(instruction.TrueFunction.Name)
			if err != nil {
				return fmt.Errorf("getting conditional true function: %w", err)
			}

			falseFunction, err := r.Function(instruction.FalseFunction.Name)
			if err != nil {
				return fmt.Errorf("getting condition false function: %w", err)
			}

			steps = append(steps, func(stack *Stack) {
				r.logger.Debugf("%s %s", instruction, stack)
				conditionFunction.Invoke(stack)

				value, err := stack.Pop()
				if err != nil {
					panic(fmt.Sprintf("enable to pop conditional condition function result: %v", err))
				}

				if value.(bool) {
					trueFunction.Invoke(stack)
				} else {
					falseFunction.Invoke(stack)
				}
			})
		default:
			return fmt.Errorf("instruction type cannot be inside a function: %s", instruction)
		}
	}

	r.functions[name] = NewFunction(steps)
	return nil
}
