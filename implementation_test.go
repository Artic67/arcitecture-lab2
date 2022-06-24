package lab2

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Functions that tests Computing of Postfix expression when compute is failed. Expression in this example has syntax errors.
func TestComputePostfixFailed(t *testing.T) {

	_, error := ComputePostfix(" 4 + ")
	assert.NotNil(t, error)

	_, error = ComputePostfix(" - ")
	assert.NotNil(t, error)

	_, error = ComputePostfix(" 9 + ")
	assert.NotNil(t, error)

	_, error = ComputePostfix("1 7")
	assert.NotNil(t, error)

	_, error = ComputePostfix("qwertyuiop")
	assert.NotNil(t, error)
}

// Function that tests Computing of Postfix expression when compute is successful.
func TestComputePostfixSuccess(t *testing.T) {

	result, error := ComputePostfix(" 6 4 - 5 * 9 + ")
	if assert.Nil(t, error) {
		assert.Equal(t, 19.0, result)
	}

	result, error = ComputePostfix(" 2 4 ^ 18 36 / +")
	if assert.Nil(t, error) {
		assert.Equal(t, 16.5, result)
	}

	result, error = ComputePostfix(" 6 18 9 / 3 2 - ^ * ")
	if assert.Nil(t, error) {
		assert.Equal(t, 12.0, result)
	}

	result, error = ComputePostfix(" 9 4 * 197 107 - 5 / - 45 9 / + 7 24 + 12 - - ")
	if assert.Nil(t, error) {
		assert.Equal(t, 4.0, result)
	}
}

//Example of function that is computing infix from postfix expression
func ExampleComputePostfix() {

	result, _ := ComputePostfix(" 8 4 +")
	fmt.Println(result)
	// Output: 12

}