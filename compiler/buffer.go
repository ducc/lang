package compiler

import (
	"bytes"
	"fmt"

	"go.uber.org/zap"
)

type buffer struct {
	logger *zap.SugaredLogger
	*bytes.Buffer
}

func (c *Compiler) newBuffer() *buffer {
	return &buffer{logger: c.logger, Buffer: bytes.NewBuffer([]byte{})}
}

func (b buffer) write(template string, vars ...interface{}) {
	content := fmt.Sprintf(template, vars...)
	_, err := b.WriteString(content + "\n")
	if err != nil {
		b.logger.With("error", err, "content", content).Fatal("writing to buffer")
	}
}
