package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLiteral(t *testing.T) {
	// arrange
	p := Literal("test")
	input := "test_desu"
	expected := Result("test", "_desu")
	// act
	actual, err := p.Parse(input)
	// assert
	assert.Equal(t, expected, actual)
	assert.NoError(t, err)
}

func TestLiteralFail(t *testing.T) {
	p := Literal("test")
	input := "different input"
	expected := Empty[string]()
	// act
	actual, err := p.Parse(input)
	// assert
	assert.Equal(t, expected, actual)
	assert.Error(t, err)
}

func TestDigit(t *testing.T) {
	// arrange
	p := Digit()
	input := "100x"
	expected := Result(100, "x")
	// act
	actual, err := p.Parse(input)
	// assert
	assert.Equal(t, expected, actual)
	assert.NoError(t, err)
}

func TestDigitFail(t *testing.T) {
	// arrange
	p := Digit()
	input := "x100"
	expected := Empty[int]()
	// act
	actual, err := p.Parse(input)
	// assert
	assert.Equal(t, expected, actual)
	assert.Error(t, err)
}
