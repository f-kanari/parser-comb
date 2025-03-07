package parser

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
