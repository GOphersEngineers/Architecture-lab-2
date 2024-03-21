package lab2

import (
	"bufio"
	"fmt"
	"io"
)

// ComputeHandler should be constructed with input io.Reader and output io.Writer.
// Its Compute() method should read the expression from input and write the computed result to the output.
type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	scanner := bufio.NewScanner(ch.Input)
	scanner.Scan()
	input := scanner.Text()

	result, err := PrefixToPostfix(input)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(ch.Output, result)
	return err
}
