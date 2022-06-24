package lab2

import (
	"fmt"
	"io"
	"io/ioutil"
)

// ComputeHandler should be constructed with input io.Reader and output io.Writer.
// Its Compute() method should read the expression from input and write the computed result to the output.


// ComputeHandler structure with reader and writer
type ComputeHandler struct {
	// TODO: Add necessary fields.
	Input io.Reader
	Output io.Writer
}

// ComputeHandler structure method that calls ComputePostfix function and then write it.
func (computeHandler *ComputeHandler) Compute() error {
	data, err := ioutil.ReadAll(computeHandler.Input)
	if err != nil {
		return err
	}

	res, err := ComputePostfix(string(data))

	if err != nil {
		return err
	}

	out := []byte(fmt.Sprintf("%f", res))
	_, err = computeHandler.Output.Write(out)

	return err
}
