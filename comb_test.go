package parser

import (
	"parser/types"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPair(t *testing.T) {
	// arrange
	p := Pair(Digit(), Literal("+"))
	input := "1+1"
	expected := Result(types.NewTuple(1, "+"), "1")
	// act
	actual, err := p.Parse(input)
	// assert
	assert.Equal(t, expected, actual)
	assert.NoError(t, err)
}

func TestMap(t *testing.T) {
	p := Map(Digit(), func(v int) string { return strconv.Itoa(v) })
	input := "100+1"
	expected := Result("100", "+1")
	// act
	actual, err := p.Parse(input)
	// assert
	assert.Equal(t, expected, actual)
	assert.NoError(t, err)
}

func TestLeft(t *testing.T) {
	p := Left(Digit(), Literal(","))
	input := "10,1"
	expected := Result(10, "1")
	// act
	actual, err := p.Parse(input)
	// assert
	assert.Equal(t, expected, actual)
	assert.NoError(t, err)
}

func TestRight(t *testing.T) {
	p := Right(Literal(":"), Digit())
	input := ":100,"
	expected := Result(100, ",")
	// act
	actual, err := p.Parse(input)
	// assert
	assert.Equal(t, expected, actual)
	assert.NoError(t, err)
}

func TestMany0(t *testing.T) {
	// arrange
	tmp := Right(Literal(","), Digit())
	p := Many0(tmp)
	tests := []struct {
		input    string
		expected ParseResult[[]int]
	}{
		{",1,2,3,", Result([]int{1, 2, 3}, ",")},
		{"1,2,3", Result([]int{}, "1,2,3")},
	}
	for _, tt := range tests {
		// act
		actual, err := p.Parse(tt.input)
		// assert
		assert.Equal(t, tt.expected, actual)
		assert.NoError(t, err)
	}
}

func TestMany1(t *testing.T) {
	// arrange
	tmp := Right(Literal(","), Digit())
	p := Many0(tmp)
	input := ",1,2,3,"
	expected := Result([]int{1, 2, 3}, ",")
	// act
	actual, err := p.Parse(input)
	// assert
	assert.Equal(t, expected, actual)
	assert.NoError(t, err)
}

func TestMany1Fail(t *testing.T) {
	// arrange
	tmp := Right(Literal(","), Digit())
	p := Many1(tmp)
	input := "1,2,3,"
	expected := Empty[[]int]()
	// act
	actual, err := p.Parse(input)
	// assert
	assert.Equal(t, expected, actual)
	assert.Error(t, err)
}
