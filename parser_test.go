package parser

import (
	"testing"
)

func TestLiteral(t *testing.T) {
	// arrange
	p := Literal("test")
	input := "test_desu"
	expected := Result("test", "_desu")
	// act
	actual, err := p.Parse(input)
	// assert
	if actual != expected {
		t.Errorf("%v expected but got %v", expected, actual)
	}
	if err != nil {
		t.Errorf("error must be nil but got %v", err)
	}
}

func TestDigit(t *testing.T) {
	// arrange
	p := Digit()
	input := "100x"
	expected := Result(100, "x")
	// act
	actual, err := p.Parse(input)
	// assert
	if actual != expected {
		t.Errorf("%v expected but got %v", expected, actual)
	}
	if err != nil {
		t.Errorf("error must be nil but got %v", err)
	}
}
