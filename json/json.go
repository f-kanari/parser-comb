package json

import (
	"fmt"
	"parser"
	"parser/types"
	"sort"
	"strings"
)

// json = "{" ( member ( "," member )* )? "}"
// member = string ":" value
// value = string | number | "true" | "false" | "null"
// array = "[" ( value ( "," value )* )? "]"
// string = '"' character* '"'
// character = [a-zA-Z_]*
// number = 0 | [1-9] [0-9]*

type ValueType int

const (
	TypeObject ValueType = iota
	TypeArray
	TypeString
	TypeNumber
	TypeBoolean
	TypeNull
)

// JSON値を構造体で表現する
type JSON struct {
	Type    ValueType
	Object  map[string]JSON
	Array   []JSON
	Str     string
	Number  int
	Boolean bool
}

func (j JSON) String() string {
	switch j.Type {
	case TypeObject:
		if len(j.Object) == 0 {
			return "{}"
		}
		var pairs []string
		for k, v := range j.Object {
			pairs = append(pairs, fmt.Sprintf("key=%q:val=%s", k, v.String()))
		}
		sort.Strings(pairs) // マップの順序を一定にする
		return strings.Join(pairs, "\n")
	case TypeArray:
		if len(j.Array) == 0 {
			return "[]"
		}
		var elements []string
		for _, v := range j.Array {
			elements = append(elements, v.String())
		}
		return "[" + strings.Join(elements, ",") + "]"
	case TypeString:
		return fmt.Sprintf("(type=str val=%q)", j.Str)
	case TypeNumber:
		return fmt.Sprintf("(type=num val=%d)", j.Number)
	case TypeBoolean:
		return fmt.Sprintf("(type=bool val=%v)", j.Boolean)
	case TypeNull:
		return "(type=null val=null)"
	default:
		return "unknown"
	}
}

func Object(m map[string]JSON) JSON {
	return JSON{
		Type:   TypeObject,
		Object: m,
	}
}

func Array(arr []JSON) JSON {
	return JSON{
		Type:  TypeArray,
		Array: arr,
	}
}

func Str(s string) JSON {
	return JSON{
		Type: TypeString,
		Str:  s,
	}
}

func Number(i int) JSON {
	return JSON{
		Type:   TypeNumber,
		Number: i,
	}
}

func Boolean(b bool) JSON {
	return JSON{
		Type:    TypeBoolean,
		Boolean: b,
	}
}

func Null() JSON {
	return JSON{
		Type: TypeNull,
	}
}

// object = "{" (key ":" value,)* "}"
func JSONParser() parser.Parser[JSON] {
	keyValParser := parser.Many0(
		parser.Pair(
			parser.Left(
				QuoteStringParser(),
				parser.Literal(":"),
			),
			parser.Left(
				ValueParser(),
				parser.Opt(parser.Literal(",")),
			),
		),
	)
	objectParser := parser.Left(
		parser.Right(
			parser.Literal("{"),
			keyValParser,
		),
		parser.Literal("}"),
	)
	return parser.Map(objectParser, func(entries []types.Tuple[string, JSON]) JSON {
		m := map[string]JSON{}
		for _, entry := range entries {
			m[entry.Fst()] = entry.Snd()
		}
		return Object(m)
	})
}

// value = array | string | number | "true" | "false" | "null"
func ValueParser() parser.Parser[JSON] {
	return parser.OneOf(
		StrParser(),
		NumberParser(),
		BooleanParser(),
		NullParser(),
	)
}

// string = ".*"
func StrParser() parser.Parser[JSON] {
	return parser.Map(QuoteStringParser(), func(s string) JSON {
		return Str(s)
	})
}

// number = [0-9]*
func NumberParser() parser.Parser[JSON] {
	return parser.Map(parser.Digit(), func(val int) JSON {
		return Number(val)
	})
}

// boolean = "true" | "false"
func BooleanParser() parser.Parser[JSON] {
	return parser.Map(parser.Either(parser.Literal("true"), parser.Literal("false")), func(s string) JSON {
		return Boolean(s == "true")
	})
}

// null = "null"
func NullParser() parser.Parser[JSON] {
	return parser.Map(parser.Literal("null"), func(_ string) JSON {
		return Null()
	})
}

// quote_string = ".*"
func QuoteStringParser() parser.Parser[string] {
	return parser.Map(parser.Right(
		parser.Literal("\""),
		parser.Left(
			parser.Many0(
				parser.Pred(parser.AnyChar(), func(c rune) bool { return c != '"' }),
			),
			parser.Literal("\""),
		),
	),
		func(chars []rune) string {
			return string(chars)
		},
	)
}
