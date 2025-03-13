package main

import (
	"fmt"
	"parser/calc"
	"parser/json"
	"strings"
	"unicode"
)

func main() {
	calcParser := calc.CalcParser()
	ret, err := calcParser.Parse("100/4*2+1")
	if err != nil {
		fmt.Printf("got err=%v\n", err)
	}
	fmt.Printf("parsed=%s\n", ret.Parsed)
	fmt.Printf("evaled=%d\n", ret.Parsed.Eval())

	jsonParser := json.JSONParser()
	input := `
  {
    "key": true,
    "key2": false,
    "key3": "test",
    "key4": 10,
  }
  `
	ret2, err := jsonParser.Parse(removeWhitespace(input))
	if err != nil {
		fmt.Printf("got err=%v\n", err)
	}
	fmt.Printf("parsed=%v\n", ret2.Parsed)
}

func removeWhitespace(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1 // -1 を返すと、その文字は削除される
		}
		return r
	}, str)
}
