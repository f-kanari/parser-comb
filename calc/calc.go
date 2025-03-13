package calc

import (
	"fmt"
	"parser"
	"parser/types"
)

type OperatorType int

const (
	ADD OperatorType = iota
	SUB
	MUL
	DIV
)

// OperatorTypeのString実装
func (o OperatorType) String() string {
	switch o {
	case ADD:
		return "+"
	case SUB:
		return "-"
	case MUL:
		return "*"
	case DIV:
		return "/"
	default:
		return "unknown"
	}
}

type NodeType int

const (
	OPERATOR NodeType = iota
	NUMBER
)

type Node struct {
	Type     NodeType
	Value    int
	Operator OperatorType
	Left     *Node
	Right    *Node
}

// Nodeのstring実装
func (n *Node) String() string {
	if n == nil {
		return ""
	}

	if n.Type == NUMBER {
		return fmt.Sprintf("%d", n.Value)
	}

	return fmt.Sprintf("(%s %s %s)",
		n.Left.String(),
		n.Operator.String(),
		n.Right.String())
}

func NewNumber(val int) *Node {
	return &Node{Type: NUMBER, Value: val}
}

func NewOperator(operator OperatorType, left, right *Node) *Node {
	return &Node{
		Type:     OPERATOR,
		Operator: operator,
		Left:     left,
		Right:    right,
	}
}

func (node *Node) Eval() int {
	if node.Type == NUMBER {
		return node.Value
	}

	lhs := node.Left.Eval()
	rhs := node.Right.Eval()

	switch node.Operator {
	case ADD:
		return lhs + rhs
	case SUB:
		return lhs - rhs
	case MUL:
		return lhs * rhs
	case DIV:
		if rhs == 0 {
			panic("div by 0")
		}
		return lhs / rhs
	default:
		panic("unknown operator")
	}
}

// expr := number(operator number)*
func CalcParser() parser.Parser[*Node] {
	symAndNums := parser.Many0(parser.Pair(OperatorParser(), NumberParser()))
	p := parser.Pair(NumberParser(), symAndNums)
	return parser.Map(p, func(tup types.Tuple[*Node, []types.Tuple[OperatorType, *Node]]) *Node {
		node := tup.Fst()
		for _, opAndNum := range tup.Snd() {
			node = NewOperator(opAndNum.Fst(), node, opAndNum.Snd())
		}
		return node
	})
}

// number:= [0-9]*
func NumberParser() parser.Parser[*Node] {
	return parser.Map(parser.Digit(), func(val int) *Node {
		return NewNumber(val)
	})
}

// operator := (+|-|*|/)
func OperatorParser() parser.Parser[OperatorType] {
	return parser.Map(parser.OneOf(
		parser.Literal("+"),
		parser.Literal("-"),
		parser.Literal("*"),
		parser.Literal("/"),
	), func(op string) OperatorType {
		switch op {
		case "+":
			return ADD
		case "-":
			return SUB
		case "*":
			return MUL
		case "/":
			return DIV
		default:
			panic("unexpcted operator")
		}
	})
}
