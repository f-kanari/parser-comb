package parser

import (
	"fmt"
	"parser/types"
)

func Pair[A any, B any](p1 Parser[A], p2 Parser[B]) Parser[types.Tuple[A, B]] {
	return New(func(s string) (ParseResult[types.Tuple[A, B]], error) {
		r1, err := p1.Parse(s)
		if err != nil {
			return Empty[types.Tuple[A, B]](), err
		}
		r2, err := p2.Parse(r1.Rest)
		if err != nil {
			return Empty[types.Tuple[A, B]](), err
		}
		return Result(types.NewTuple(r1.Parsed, r2.Parsed), r2.Rest), nil
	})
}

func Map[A any, B any](p Parser[A], mapFn func(A) B) Parser[B] {
	return New(func(s string) (ParseResult[B], error) {
		r, err := p.Parse(s)
		if err != nil {
			return Empty[B](), err
		}
		return Result(mapFn(r.Parsed), r.Rest), nil
	})
}

func Left[A any, B any](p1 Parser[A], p2 Parser[B]) Parser[A] {
	return Map(Pair(p1, p2), func(tuple types.Tuple[A, B]) A {
		return tuple.Fst()
	})
}

func Right[A any, B any](p1 Parser[A], p2 Parser[B]) Parser[B] {
	return Map(Pair(p1, p2), func(tuple types.Tuple[A, B]) B {
		return tuple.Snd()
	})
}

func Repeat[A any](p Parser[A]) Parser[[]A] {
	return New(func(s string) (ParseResult[[]A], error) {
		results := []A{}
		for ret, err := p.Parse(s); err != nil; {
			results = append(results, ret.Parsed)
			s = ret.Rest
		}
		return Result(results, s), nil
	})
}

func Pred[A any](p Parser[A], pred func(A) bool) Parser[A] {
	return New(func(s string) (ParseResult[A], error) {
		r, err := p.Parse(s)
		if err != nil {
			return Empty[A](), err
		}
		if !pred(r.Parsed) {
			return Empty[A](), fmt.Errorf("%v does not match the cond", r.Parsed)
		}
		return r, nil
	})
}

func Either[A any](p1 Parser[A], p2 Parser[A]) Parser[A] {
	return New(func(s string) (ParseResult[A], error) {
		r, err := p1.Parse(s)
		if err != nil {
			return p2.Parse(s)
		}
		return r, nil
	})
}
