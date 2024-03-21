package main

import (
	"fmt"
	"os"
	"flag"
	"io"
	"strings"

	lab2 "github.com/GOphersEngineers/Architecture-lab-2"
)

func main() {
	var expr string
	var inFile string
	var outFile string

	flag.StringVar(&expr, "e", "", "Expression to convert from prefix to postfix")
	flag.StringVar(&inFile, "f", "", "File with expression to convert from prefix to postfix")
	flag.StringVar(&outFile, "o", "", "File to write the result of conversion")

	flag.Parse()

	var reader io.Reader
	var writer io.Writer

	if expr != "" {
		reader = strings.NewReader(expr)
	} else if inFile != "" {
		file, err := os.Open(inFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer file.Close()
		reader = file
	} else {
		reader = os.Stdin
	}

	if outFile != "" {
		file, err := os.Create(outFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer file.Close()
		writer = file
	} else {
		writer = os.Stdout
	}

	handler := &lab2.ComputeHandler{
		Input:  reader,
		Output: writer,
	}

	err := handler.Compute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}