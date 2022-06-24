package main

import (
	"flag"
	"io"
	"io/ioutil"
	"os"
	"strings"
	lab2 "github.com/Artic67/arcitecture-lab2"
)

var (
	// TODO: Add other flags support for input and output configuration.
	// Console input arguments in pointer variables
	defaultValue = ""
	inputExpressionPtr = flag.String("e", "", "Expression to compute")
	inputExpFilePtr = flag.String("f", "", "Filename with expression")
	inputOutFilePtr = flag.String("o", "", "Name of output file")
)

func main() {

	flag.Parse()

	// Take values from pointers
	inputExpFile := *inputExpFilePtr
	inputOutFile := *inputOutFilePtr
	inputExpression := *inputExpressionPtr

	var fromFile io.Reader

	// Checking input arguments
	if inputExpression != defaultValue {
		fromFile = strings.NewReader(inputExpression)
	} else if inputExpFile != defaultValue {
		data, error := ioutil.ReadFile(inputExpFile)
		if error != nil {
			os.Stderr.WriteString("Provided file hasn't been found")
			return
		}
		fromFile = strings.NewReader(string(data))
	} else {
		os.Stderr.WriteString("Can't find expression")
		return
	}

	// Creating out file in directory
	var (
		toFile io.Writer
		error error
	)
	if inputOutFile != defaultValue {
		if toFile, error = os.Create(inputOutFile); error != nil {
			os.Stderr.WriteString("An error occupied while creating file")
		}
	} else {
		toFile = os.Stdout
	}

	// Making a handler
	handler := &lab2.ComputeHandler{
		Input: fromFile,
		Output: toFile,
	}

	// Computing
	if error := handler.Compute(); error != nil {
		os.Stderr.WriteString(error.Error())
	}

	// TODO: Change this to accept input from the command line arguments as described in the task and
	//       output the results using the ComputeHandler instance.
	//       handler := &lab2.ComputeHandler{
	//           Input: {construct io.Reader according the command line parameters},
	//           Output: {construct io.Writer according the command line parameters},
	//       }
	//       err := handler.Compute()
}
