package json

import (
	"parser"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoolean(t *testing.T) {
	p := BooleanParser()
	expected := parser.Result(Boolean(true), "")
	actual, err := p.Parse("true")
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestNull(t *testing.T) {
	p := NullParser()
	expected := parser.Result(Null(), "")
	actual, err := p.Parse("null")
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestQuoteString(t *testing.T) {
	p := QuoteStringParser()
	expected := parser.Result("test", "")
	actual, err := p.Parse("\"test\"")
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}
