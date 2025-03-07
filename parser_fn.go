package parser

import (
	"fmt"
	"strconv"
	"strings"
)

// Parserの結果を表現する構造体
type ParseResult[T any] struct {
	Parsed T
	Rest   string
}

func (res ParseResult[T]) String() string {
	return fmt.Sprintf("parsed=%v, rest=%s", res.Parsed, res.Rest)
}

func Result[T any](parsed T, rest string) ParseResult[T] {
	return ParseResult[T]{Parsed: parsed, Rest: rest}
}

func Empty[T any]() ParseResult[T] {
	return ParseResult[T]{}
}

// go 1.24から
// ParserFnは、対象をうけとり、解析済みと残りにわける
// type ParserFn[T any] = func(string) (ParseResult[T], error)

// 文字列のパーサー
func Literal(expected string) func(string) (ParseResult[string], error) {
	return func(input string) (ParseResult[string], error) {
		if !strings.HasPrefix(input, expected) {
			return Empty[string](), fmt.Errorf("%s expected but got %s", expected, input)
		}
		return Result(expected, input[len(expected):]), nil
	}
}

func Digit() func(string) (ParseResult[int], error) {
	return func(input string) (ParseResult[int], error) {
		num := []rune{}
		for _, c := range input {
			if c < '0' || '9' < c {
				break
			}
			num = append(num, c)
		}
		val, err := strconv.Atoi(string(num))
		if err != nil {
			return Empty[int](), err
		}
		return Result(val, input[len(num):]), nil
	}
}
