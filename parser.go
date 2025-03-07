package parser

import (
	"fmt"
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

// パーサー: 文字列を処理し、処理済みと残りに分ける
type Parser[T any] struct {
	fn func(string) (ParseResult[T], error)
}

func New[T any](fn func(string) (ParseResult[T], error)) Parser[T] {
	return Parser[T]{fn: fn}
}

func (p Parser[T]) Parse(input string) (ParseResult[T], error) {
	return p.fn(input)
}
