package main

import (
	"flag"
	"io"
	"os"
	"strings"

	lab2 "github.com/roman-mazur/architecture-lab-2"
)

var (
	flagE = flag.String("e", "", "Expression to compute")
	flagF = flag.String("f", "", "Input file")
	flagO = flag.String("o", "", "Output file")
)

func main() {
	flag.Parse()

	var input io.Reader = nil
	var output = os.Stdout

	if *flagE != "" {
		input = strings.NewReader(*flagE)
	}

	if *flagF != "" {
		f, err := os.Open(*flagF)
		if err != nil {
			os.Stderr.WriteString("Error")
		}
		defer f.Close()
		input = f
	}

	if *flagO != "" {
		f, err := os.Create(*flagO)
		if err != nil {
			os.Stderr.WriteString("Error")
		}
		defer f.Close()
		output = f
	}

	if input == nil {
		os.Stderr.WriteString("No stdIn defined")
		return
	}

	handler := &lab2.ComputeHandler{
		Input:  input,
		Output: output,
	}
	err := handler.Compute()
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
}
