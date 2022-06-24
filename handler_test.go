package lab2

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

// WriteTest structure with called boolean.
type WriteTest struct {
	called bool
}

// WriteTest structure method to write data.
func (o *WriteTest) Write(p []byte) (n int, err error) {
	o.called = true
	return 0, nil
}

// Functions that tests Compute Handler when compute is failed. Expression in this example has syntax errors.
func TestHandler_ComputeFailed(t *testing.T) {

	output := WriteTest{}
	expressionFail := " 2 m - 3 * 6 + "
	input := strings.NewReader(expressionFail)

	handler := ComputeHandler{
		Input: input,
		Output: &output,
	}

	err := handler.Compute()
	assert.NotNil(t, err)

}

// Function that tests Compute Handler when compute is successful.
func TestHandler_ComputeSuccessful(t *testing.T) {

	output := WriteTest{}
	expression := " 6 3 - 4 * 7 + "
	input := strings.NewReader(expression)

	handler := ComputeHandler{
		Input: input,
		Output: &output,
	}

	handler.Compute()
	assert.Equal(t, true, output.called)
}