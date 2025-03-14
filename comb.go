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

func Many0[A any](p Parser[A]) Parser[[]A] {
	return New(func(s string) (ParseResult[[]A], error) {
		results := []A{}
		cur := s
		for {
			ret, err := p.Parse(cur)
			if err != nil {
				break
			}
			results = append(results, ret.Parsed)
			cur = ret.Rest
		}
		return Result(results, cur), nil
	})
}

func Many1[A any](p Parser[A]) Parser[[]A] {
	return New(func(s string) (ParseResult[[]A], error) {
		results := []A{}
		ret, err := p.Parse(s)
		if err != nil {
			return Empty[[]A](), err
		}

		cur := ret.Rest
		for {
			ret, err := p.Parse(cur)
			if err != nil {
				break
			}
			results = append(results, ret.Parsed)
			cur = ret.Rest
		}
		return Result(results, cur), nil
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

func AndThen[A any, B any](p Parser[A], fn func(A) Parser[B]) Parser[B] {
	return New(func(s string) (ParseResult[B], error) {
		r, err := p.Parse(s)
		if err != nil {
			return Empty[B](), err
		}
		return fn(r.Parsed).Parse(r.Rest)
	})
}

func OneOf[A any](ps ...Parser[A]) Parser[A] {
	return New(func(s string) (ParseResult[A], error) {
		for _, p := range ps {
			r, err := p.Parse(s)
			if err != nil {
				continue
			}
			return r, nil
		}
		return Empty[A](), fmt.Errorf("all parsers failed: %v", ps)
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

func Opt[A any](p Parser[A]) Parser[types.Option[A]] {
	return New(func(s string) (ParseResult[types.Option[A]], error) {
		ret, err := p.Parse(s)
		if err != nil {
			return Result(types.None[A](), s), nil
		}
		return Result(types.Some(ret.Parsed), ret.Rest), err
	})
}
