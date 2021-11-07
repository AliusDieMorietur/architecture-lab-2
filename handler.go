package lab2

import (
	"io"
)

// ComputeHandler should be constructed with input io.Reader and output io.Writer.
// Its Compute() method should read the expression from input and write the computed result to the output.
type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
	// TODO: Add necessary fields.
}

func (ch *ComputeHandler) Compute() error {
	s := make([]byte, 0)
	p := make([]byte, 8)
	for {
		n, err := ch.Input.Read(p)
		s = append(s, p[:n]...)
		if err == io.EOF {
			break
		}
	}
	result, err := PrefixToInfix(string(s))
	if err != nil {
		return err
	}
	ch.Output.Write([]byte(result))
	return nil
}
