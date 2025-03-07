package types

// Tuple
type Tuple[A any, B any] struct {
	fst A
	snd B
}

func NewTuple[A any, B any](fst A, snd B) Tuple[A, B] {
	return Tuple[A, B]{fst: fst, snd: snd}
}

func (t Tuple[A, B]) Fst() A {
	return t.fst
}

func (t Tuple[A, B]) Snd() B {
	return t.snd
}

// Either
type Either[A any, B any] struct {
	left  Option[A]
	right Option[B]
}

func Left[A any, B any](val A) Either[A, B] {
	return Either[A, B]{
		left:  Some(val),
		right: None[B](),
	}
}

func Right[A any, B any](val B) Either[A, B] {
	return Either[A, B]{
		left:  None[A](),
		right: Some(val),
	}
}

func (e Either[A, B]) Left() (A, bool) {
	return e.left.val, e.left.valid
}

func (e Either[A, B]) Right() (B, bool) {
	return e.right.val, e.right.valid
}

// Option
type Option[A any] struct {
	val   A
	valid bool
}

func Some[A any](val A) Option[A] {
	return Option[A]{val: val, valid: true}
}

func None[A any]() Option[A] {
	return Option[A]{valid: false}
}
