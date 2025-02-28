package parser

// Tuple
type Tuple[A any, B any] struct {
	fst A
	snd B
}

func (t Tuple[A, B]) Fst() A {
	return t.fst
}

func (t Tuple[A, B]) Snd() B {
	return t.snd
}

// Either
type Either[A any, B any] struct {
	left  Maybe[A]
	right Maybe[B]
}

func Left[A any, B any](val A) Either[A, B] {
	return Either[A, B]{
		left:  Just(val),
		right: Nothing[B](),
	}
}

func Right[A any, B any](val B) Either[A, B] {
	return Either[A, B]{
		left:  Nothing[A](),
		right: Just(val),
	}
}

func (e Either[A, B]) Left() (A, bool) {
	return e.left.val, e.left.valid
}

func (e Either[A, B]) Right() (B, bool) {
	return e.right.val, e.right.valid
}

// Maybe(Optional)
type Maybe[A any] struct {
	val   A
	valid bool
}

func Just[A any](val A) Maybe[A] {
	return Maybe[A]{val: val, valid: true}
}

func Nothing[A any]() Maybe[A] {
	return Maybe[A]{valid: false}
}
