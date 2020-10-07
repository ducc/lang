package main

import (
	"flag"

	"io/ioutil"

	"github.com/ducc/lang/parser"
	"github.com/ducc/lang/runtime"
	"go.uber.org/zap"
)

var (
	file string
)

func init() {
	flag.StringVar(&file, "f", "", "file to run")
}

func main() {
	flag.Parse()
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	sugar := logger.Sugar()

	if file == "" {
		sugar.Fatal("file must be given with -f parameter")
	}

	fileData, err := ioutil.ReadFile(file)
	if err != nil {
		sugar.With("error", err).Fatal("unable to read file")
	}

	definitions, err := parser.ParseFunctionDefinitions(string(fileData))
	if err != nil {
		sugar.With("error", err).Fatal("parsing function definitions")
	}

	if err := runtime.Run(sugar, definitions); err != nil {
		sugar.With("error", err).Fatal("unable to run runtime")
	}
}
