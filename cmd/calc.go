package main

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

// TODO: implement calcparser
