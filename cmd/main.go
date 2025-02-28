package main

import (
	"fmt"
	"parser"
)

func main() {
	input := "++++xxxx"
	p := parser.Literal("+")
	repeated := parser.Repeat(p)
	ret, err := repeated.Parse(input)
	fmt.Printf("ret=%s, err=%v", ret, err)
}
