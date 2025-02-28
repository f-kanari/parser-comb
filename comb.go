package parser

func And[A any, B any](p1 Parser[A], p2 Parser[B]) Parser[Tuple[A, B]] {
	return New(func(input string) (ParseResult[Tuple[A, B]], error) {
		ret1, err := p1.Parse(input)
		if err != nil {
			return ParseResult[Tuple[A, B]]{}, err
		}
		ret2, err := p2.Parse(ret1.Rest)
		if err != nil {
			return ParseResult[Tuple[A, B]]{}, err
		}
		return ParseResult[Tuple[A, B]]{
			Parsed: Tuple[A, B]{fst: ret1.Parsed, snd: ret2.Parsed},
			Rest:   ret2.Rest,
		}, nil
	})
}

func Or[A any, B any](p1 Parser[A], p2 Parser[B]) Parser[Either[A, B]] {
	return New(func(input string) (ParseResult[Either[A, B]], error) {
		ret1, err := p1.Parse(input)
		if err == nil {
			return ParseResult[Either[A, B]]{Parsed: Left[A, B](ret1.Parsed), Rest: ret1.Rest}, nil
		}
		ret2, err := p2.Parse(input)
		if err == nil {
			return ParseResult[Either[A, B]]{Parsed: Right[A](ret2.Parsed), Rest: ret2.Rest}, nil
		}
		return ParseResult[Either[A, B]]{}, err
	})
}

func Repeat[A any](p Parser[A]) Parser[[]A] {
	return New(func(input string) (ParseResult[[]A], error) {
		results := []A{}
		cur := input
		for {
			ret, err := p.Parse(cur)
			if err != nil {
				return ParseResult[[]A]{Parsed: results, Rest: cur}, nil
			}
			cur = ret.Rest
			results = append(results, ret.Parsed)
			if len(cur) == 0 {
				return ParseResult[[]A]{Parsed: results, Rest: ret.Rest}, nil
			}
		}
	})
}
