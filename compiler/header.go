package compiler

const MainFunctionName = "main"

func WriteHeader(buf *buffer, imports []string, withMain bool) {
	buf.write("// auto generated - changes will not be persisted")
	buf.write("package main")
	buf.write("import (")
	buf.write("\"fmt\"")
	buf.write("\"github.com/ducc/lang/util\"")
	buf.write("\"github.com/ducc/lang/builtins\"")
	buf.write("\"go.uber.org/zap\"")
	for _, i := range imports {
		buf.writef("\"%s\"", i)
	}
	buf.write(")")

	buf.write("// incase imports are not used")
	buf.write("var _ = fmt.Println")
	buf.write("var _ = util.NewScope")
	buf.write("var _ = builtins.Add")
	buf.write("var _ = zap.S")

	if withMain {
		buf.write("func main() {")
		buf.write("logger, _ := zap.NewDevelopment()")
		buf.write("zap.ReplaceGlobals(logger)")
		buf.write("defer logger.Sync()")
		buf.write("scope := util.NewScope()")
		buf.write("scope = ducc_main(scope)")
		buf.write("}")
	}
}
