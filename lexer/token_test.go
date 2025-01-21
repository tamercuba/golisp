package lexer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPositionStringMethod(t *testing.T) {
	p := NewPos(5, 10)
	expected := "10:5 "

	result := fmt.Sprintf("%v", p)
	assert.Equal(t, expected, result)
}
