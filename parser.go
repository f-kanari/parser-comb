package parser

import (
	"errors"
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

// basic parser
func Literal(expected string) Parser[string] {
	return New(func(s string) (ParseResult[string], error) {
		if !strings.HasPrefix(s, expected) {
			return Empty[string](), fmt.Errorf("%s expected but got %s", expected, s)
		}
		return Result(expected, s[len(expected):]), nil
	})
}

func Digit() Parser[int] {
	return New(func(s string) (ParseResult[int], error) {
		chars := []rune{}
		for _, c := range s {
			if c < '0' || '9' < c {
				break
			}
			chars = append(chars, c)
		}
		if len(chars) == 0 {
			return Empty[int](), fmt.Errorf("integer expected but got %s", s)
		}
		num, err := strconv.Atoi(string(chars))
		if err != nil {
			return Empty[int](), err
		}
		return Result(num, s[len(chars):]), nil
	})
}

func AnyChar() Parser[rune] {
	return New(func(s string) (ParseResult[rune], error) {
		chars := []rune(s[:1])
		if len(chars) == 0 {
			return Empty[rune](), errors.New("empty string")
		}
		return Result(chars[0], s[1:]), nil
	})
}
