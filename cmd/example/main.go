package main

import (
	"flag"
	"fmt"
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
			fmt.Println("Error")
		}
		input = f
	}

	if *flagO != "" {
		f, err := os.Create(*flagO)
		if err != nil {
			fmt.Println("Error")
		}
		output = f
	} 

	if input == nil {
		fmt.Println("No stdIn defined")
		return 
	} 
	
	handler := &lab2.ComputeHandler {
		Input: input,
		Output: output,
	}
	err := handler.Compute()
	if err != nil {
		fmt.Println(err.Error());
	}
}


