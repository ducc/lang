package compiler

import (
	"bytes"
	"fmt"

	"go.uber.org/zap"
)

type buffer struct {
	logger *zap.SugaredLogger
	*bytes.Buffer
	isDebug bool
}

func (c *Compiler) newBuffer() *buffer {
	return &buffer{logger: c.logger, Buffer: bytes.NewBuffer([]byte{}), isDebug: c.debug}
}

func (b buffer) writef(template string, vars ...interface{}) {
	content := fmt.Sprintf(template, vars...)
	b.write(content)
}

func (b buffer) write(content string) {
	_, err := b.WriteString(content + "\n")
	if err != nil {
		b.logger.With("error", err, "content", content).Fatal("writing to buffer")
	}
}

func (b buffer) debugf(template string, vars ...interface{}) {
	if b.isDebug {
		b.writef(template, vars...)
	}
}

func (b buffer) debug(content string) {
	if b.isDebug {
		b.write(content)
	}	
}
