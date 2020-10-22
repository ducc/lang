package compiler

import (
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

	var withMain bool
	for _, definition := range definitions {
		if definition.InstructionType() == parser.InstructionTypeDefineFunction && definition.DefineFunction().Name == MainFunctionName {
			withMain = true
			break
		}
	}

	// package, imports, unused variable hacks
	WriteHeader(buf, getImports(definitions), withMain)

	for _, definition := range definitions {
		funcBuf := c.newBuffer()

		if err := compileFunction(funcBuf, definition); err != nil {
			return nil, err
		}

		_, err := buf.Write(funcBuf.Bytes())
		if err != nil {
			return nil, err
		}
	}

	return buf, nil
}
