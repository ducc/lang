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

func (b *buffer) FunctionDefinition(name string) {
	b.WriteString(fmt.Sprintf("func %s(scope *util.Scope) *util.Scope", name))
}
func (b *buffer) OpenScope()              { b.write("{") }
func (b *buffer) CloneScope()             { b.write("scope := scope.Clone()") }
func (b *buffer) CloseScope()             { b.write("}") }
func (b *buffer) ReturnScope()            { b.write("return scope") }
func (b *buffer) PushStack(v string)      { b.writef("scope = scope.PushStack(%v)", v) }
func (b *buffer) PushVarToStack(v string) { b.PushStack(fmt.Sprintf("scope.GetVar(\"%s\")", v)) }
func (b *buffer) SetVar(k, v string)      { b.writef("scope = scope.SetVar(\"%s\", %s)", k, v) }
func (b *buffer) CallFunctionVar(v string) {
	b.writef("scope = scope.GetVar(\"%s\").(func(*util.Scope) *util.Scope)(scope)", v)
}
func (b *buffer) ZapDebugCall(v string) { b.writef("zap.S().Debug(\"call %20s \", scope.String())", v) }
