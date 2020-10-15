package main

import (
	"flag"

	"io/ioutil"

	"github.com/ducc/lang/compiler"
	"github.com/ducc/lang/parser"
	"go.uber.org/zap"
)

var (
	inputFile  string
	outputFile string
	debug bool 
)

func init() {
	flag.StringVar(&inputFile, "f", "", "file to run")
	flag.StringVar(&outputFile, "o", "output.go", "where to write compiled go")
	flag.BoolVar(&debug, "debug", true, "enable debug logging in output file")
}

func main() {
	flag.Parse()
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
	defer logger.Sync()

	sugar := logger.Sugar()

	if inputFile == "" {
		sugar.Fatal("file must be given with --f parameter")
	}

	if outputFile == "" {
		sugar.Fatal("output file msut be given with --o parameter")
	}

	fileData, err := ioutil.ReadFile(inputFile)
	if err != nil {
		sugar.With("error", err).Fatal("unable to read file")
	}

	definitions, err := parser.ParseFunctions(string(fileData))
	if err != nil {
		sugar.With("error", err).Fatal("parsing function definitions")
	}

	outputReader, err := compiler.New(sugar, debug).Compile(definitions)
	if err != nil {
		sugar.With("error", err).Fatal("unable to compile")
	}

	outputData, err := ioutil.ReadAll(outputReader)
	if err != nil {
		sugar.With("error", err).Fatal("unable to read output")
	}

	if err := ioutil.WriteFile(outputFile, outputData, 0600); err != nil {
		sugar.With("error", err).Fatal("unable to write output file")
	}
}
