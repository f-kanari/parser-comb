package main

import (
	"fmt"
	"os"
	"parser/calc"
	"parser/json"
	"strings"
	"unicode"
)

const (
	CALC = "calc"
	JSON = "json"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "error: invalid args len")
		return
	}
	parserType := os.Args[1]
	input := os.Args[2]

	if parserType == CALC {
		p := calc.CalcParser()
		ret, err := p.Parse(input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to parse: %v\n", err)
			return
		}
		fmt.Printf("%s evaluated into %d\n", ret.Parsed, ret.Parsed.Eval())
		return
	}

	if parserType == JSON {
		input = removeWhitespace(input)
		p := json.JSONParser()
		ret, err := p.Parse(input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: failed to parse %v\n", err)
			return
		}
		fmt.Printf("json=%s", ret.Parsed)
		return
	}
}

func removeWhitespace(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1 // -1 を返すと、その文字は削除される
		}
		return r
	}, str)
}
