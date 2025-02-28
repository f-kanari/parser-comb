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

// パーサー: 文字列を処理し、処理済みと残りに分ける
type Parser[T any] struct {
	fn func(string) (ParseResult[T], error)
}

func (p Parser[T]) Parse(input string) (ParseResult[T], error) {
	return p.fn(input)
}

func New[T any](fn func(string) (ParseResult[T], error)) Parser[T] {
	return Parser[T]{fn: fn}
}

// 期待した文字列か
func Literal(expected string) Parser[string] {
	return New(func(input string) (ParseResult[string], error) {
		if !strings.HasPrefix(input, expected) {
			return ParseResult[string]{}, fmt.Errorf("ParseErr: %s expected but got %s", expected, input)
		}
		return ParseResult[string]{Parsed: input[:len(expected)], Rest: input[len(expected):]}, nil
	})
}

// 数値のパーサー
func Int() Parser[int] {
	return New(func(input string) (ParseResult[int], error) {
		var num []rune
		for _, c := range input {
			if c < '0' || c > '9' {
				break
			}
			num = append(num, c)
		}
		if len(num) == 0 {
			return ParseResult[int]{}, fmt.Errorf("ParseErr: interger expected got %s", input)
		}
		converted, err := strconv.Atoi(string(num))
		if err != nil {
			return ParseResult[int]{}, fmt.Errorf("`strconv.Atoi` failed: %w", err)
		}
		return ParseResult[int]{Parsed: converted, Rest: input[len(num):]}, nil
	})
}
