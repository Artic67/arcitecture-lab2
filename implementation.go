package lab2

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

// ComputePostfix function that parse and calculates all data
func ComputePostfix(input string) (float64, error) {

	stack := &Stack{nil, 0}

	// Filter operators and remove spaces
	operators := Filter(strings.Split(input, " "), func(character string) bool {
		return character != "" && character != " "
	})

	for _, value := range operators {

		if num, e := strconv.ParseFloat(value, 64); e == nil {

			stack.push(&Node{num, nil})

		} else {

			// Left and right operands
			left, _ := stack.pop()
			right, err := stack.pop()

			if err != nil {
				return 0.0, err
			}

			// Checking value for operation symbols
			switch value {

			case "+":

				node := &Node{right.value + left.value, nil}
				stack.push(node)

			case "-":

				node := &Node{right.value - left.value, nil}
				stack.push(node)


			case "^":

				node := &Node{math.Pow(right.value, left.value), nil}
				stack.push(node)

			case "*":

				node := &Node{right.value * left.value, nil}
				stack.push(node)

			case "/":

				if left.value == 0 {
					return 0.0, errors.New("Zero division error!")
				}
				node := &Node{right.value / left.value, nil}
				stack.push(node)

			default:
				return 0.0, errors.New("Unexpected type of operation symbol!")
			}
		}
	}

	// When operators slice is ended we can end the calculation
	if res, err := stack.pop(); err == nil {

		if stack.length != 0 {
			return 0.0, errors.New("Mistake in expresion. Stack isn't empty! ")
		}

		return res.value, nil

	} else {
		return 0.0, err
	}
}
