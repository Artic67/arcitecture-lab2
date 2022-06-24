package lab2

import "errors"

// Node of stack structure
type Node struct {
	value float64
	next  *Node
}

// Stack structure with node and length variable
type Stack struct {
	top    *Node
	length int
}

// Pop stack method
func (stack *Stack) pop() (*Node, error) {

	// Empty stack check
	if stack.length == 0 {
		return nil, errors.New("Empty stack")
	}
	stack.length--
	node := stack.top
	stack.top = node.next
	return node, nil

}

// Push stack method
func (stack *Stack) push(node *Node) {
	stack.length++
	node.next = stack.top
	stack.top = node
}

