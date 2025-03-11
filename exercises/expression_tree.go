package exercises

import (
	"github.com/gaarutyunov/galgo/numbers"
	"github.com/gaarutyunov/galgo/stack"
)

type (
	Node interface {
		Evaluate() int
	}
	TreeBuilder interface {
		BuildTree(postfix []string) Node
	}
)

type (
	ValueNode struct {
		value int
	}
	operationNode struct {
		left  Node
		right Node
	}
	AddNode struct {
		*operationNode
	}
	SubtractNode struct {
		*operationNode
	}
	MultiplyNode struct {
		*operationNode
	}
	DivideNode struct {
		*operationNode
	}
)

type (
	Operations  map[string]func(left, right Node) Node
	BuilderFunc func(postfix []string) Node
)

func Value(v int) *ValueNode {
	return &ValueNode{value: v}
}

func Add(left, right Node) Node {
	return &AddNode{&operationNode{left, right}}
}

func Subtract(left, right Node) Node {
	return &SubtractNode{&operationNode{left: left, right: right}}
}

func Multiply(left, right Node) Node {
	return &MultiplyNode{&operationNode{left: left, right: right}}
}

func Divide(left, right Node) Node {
	return &DivideNode{&operationNode{left: left, right: right}}
}

func (n *ValueNode) Evaluate() int {
	return n.value
}

func (d *DivideNode) Evaluate() int {
	return d.left.Evaluate() / d.right.Evaluate()
}

func (m *MultiplyNode) Evaluate() int {
	return m.left.Evaluate() * m.right.Evaluate()
}

func (s *SubtractNode) Evaluate() int {
	return s.left.Evaluate() - s.right.Evaluate()
}

func (a *AddNode) Evaluate() int {
	return a.left.Evaluate() + a.right.Evaluate()
}

func (b BuilderFunc) BuildTree(postfix []string) Node {
	return b(postfix)
}

func (o Operations) Node(s string, left, right Node) Node {
	for key, operation := range o {
		if key == s {
			return operation(left, right)
		}
	}

	panic("unknown operation")
}

func ExpressionTreeSolution(builder TreeBuilder) func(postfix []string) int {
	return func(postfix []string) int {
		return builder.BuildTree(postfix).Evaluate()
	}
}

// ExpressionTreeV1 https://leetcode.com/problems/design-an-expression-tree-with-evaluate-function
func ExpressionTreeV1(postfix []string) (n Node) {
	ops := Operations{
		"+": Add,
		"-": Subtract,
		"*": Multiply,
		"/": Divide,
	}

	s := stack.New[Node]()

	for _, v := range postfix {
		if _, ok := ops[v]; ok {
			right, _ := s.Pop()
			left, _ := s.Pop()
			s.Push(ops.Node(v, left, right))
		} else {
			s.Push(Value(numbers.FromString(v)))
		}
	}

	v, _ := s.Head()

	return v
}
