package runtime

import (
	"github.com/ducc/lang/parser"
	"github.com/ducc/lang/util"
	"go.uber.org/zap"
)

func Run(logger *zap.SugaredLogger, definedFunctions []parser.DefineFunction) error {
	functionRegistry := NewFunctionRegistry(logger)
	for _, definedFunction := range definedFunctions {
		if err := functionRegistry.Register(definedFunction); err != nil {
			return err
		}
	}

	function, err := functionRegistry.Function("main")
	if err != nil {
		return err
	}

	function.Invoke(util.NewStack())
	return nil
}
