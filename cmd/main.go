package main

import (
	"fmt"
	"parser"
)

func main() {
	parsed := []any{}
	input := "xxx10"
	literal := parser.Literal("xxx")
	digit := parser.Digit()
	r1, err := literal(input)
	if err != nil {
		fmt.Printf("failed to parse: %v\n", err)
	}
	parsed = append(parsed, r1.Parsed)
	r2, err := digit(r1.Rest)
	if err != nil {
		fmt.Printf("failed to parse: %v\n", err)
	}
	parsed = append(parsed, r2.Parsed)
	fmt.Println(parsed)
}
