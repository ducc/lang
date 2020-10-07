package main

import (
	"flag"
	"fmt"

	"io/ioutil"

	"github.com/ducc/lang/lang"
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

	_ = fileData

	definitions, err := lang.ParseFunctionDefinitions(string(fileData))
	if err != nil {
		sugar.With("error", err).Fatal("parsing function definitions")
	}

	fmt.Println(definitions)

	if err := runtime.Run(sugar, definitions); err != nil {
		sugar.With("error", err).Fatal("unable to run runtime")
	}
}
