package main

import "parser"

// GOAL: 四則演算のASTをビルドする

type NodeKind int

const (
	IntNode NodeKind = iota
	OpNode
)

type CalcNode struct {
	kind  NodeKind
	val   any // IntNodeならint, OpNodeなら演算子(string)
	left  *CalcNode
	right *CalcNode
}

func NewIntNode(val int) CalcNode {
	return CalcNode{kind: IntNode, val: val}
}

func NewOpNode(op string, left *CalcNode, right *CalcNode) CalcNode {
	return CalcNode{
		kind:  OpNode,
		val:   op,
		left:  left,
		right: right,
	}
}

func IntNodeParser() parser.Parser[CalcNode] {
	return parser.New(func(input string) (parser.ParseResult[CalcNode], error) {
		val, err := parser.Int().Parse(input)
		if err != nil {
			return parser.ParseResult[CalcNode]{}, err
		}
		return parser.ParseResult[CalcNode]{Parsed: NewIntNode(val.Parsed), Rest: val.Rest}, nil
	})
}
