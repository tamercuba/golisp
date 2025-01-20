package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tamercuba/golisp/lexer"
	"github.com/tamercuba/golisp/parser/ast"
)

func TestParseSimpleListOfIntegers(t *testing.T) {
	input := `(1 2 3)`
	lexer := lexer.NewLexer(input)
	result, err := ParseProgram(lexer)

	assert.Nil(t, err)
	assert.IsType(t, &ast.Program{}, result)
	assert.Equal(t, 1, len(result.ListStatements))

	statement := result.ListStatements[0]

	assert.IsType(t, &ast.ListExpression{}, statement)

	switch list := statement.(type) {
	case *ast.ListExpression:
		assert.Equal(t, 3, list.Size)
		currValue := 1
		currNode := list.Head
		for {
			assert.NotNil(t, currNode)
			assert.Equal(t, int32(currValue), currNode.LNode.GetValue())

			if currNode.Next == nil {
				break
			} else {
				currNode = currNode.Next
				currValue++
			}
		}
	default:
		assert.Fail(t, "Invalid type")
	}
}

func TestParseNestedListsOfNumbers(t *testing.T) {
	input := `(1 1.1 (2 3))`
	lexer := lexer.NewLexer(input)
	result, err := ParseProgram(lexer)

	assert.Nil(t, err)
	assert.IsType(t, result, &ast.Program{})
	assert.Equal(t, 1, len(result.ListStatements))

	s := result.ListStatements[0]
	switch l := s.(type) {
	case *ast.ListExpression:
		assert.NotNil(t, l.Head)
		assert.Equal(t, int32(1), l.Head.LNode.GetValue())

		assert.NotNil(t, l.Head.Next)
		assert.Equal(t, float64(1.1), l.Head.Next.LNode.GetValue())
		assert.IsType(t, &ast.ListExpression{}, l.Head.Next.Next.LNode)

		nested := l.Head.Next.Next.LNode
		switch ll := nested.(type) {
		case *ast.ListExpression:
			assert.NotNil(t, ll.Head)
			assert.Equal(t, int32(2), ll.Head.LNode.GetValue())
			assert.NotNil(t, ll.Head.Next)
			assert.Equal(t, int32(3), ll.Head.Next.LNode.GetValue())
			assert.Nil(t, ll.Head.Next.Next)
		default:
			assert.Fail(t, "Invalid type")
		}
		assert.Nil(t, l.Head.Next.Next.Next)

	default:
		assert.Fail(t, "Invalid type")
	}
}

func TestParseSumOfIntegers(t *testing.T) {
	input := `(+ 1 2)`
	lexer := lexer.NewLexer(input)
	result, err := ParseProgram(lexer)

	assert.Nil(t, err)
	assert.IsType(t, &ast.Program{}, result)
	assert.Equal(t, 1, len(result.ListStatements))

	s := result.ListStatements[0]
	switch l := s.(type) {
	case *ast.ListExpression:
		assert.Equal(t, 3, l.Size)
		assert.NotNil(t, l.Head)
		assert.Equal(t, "+", l.Head.LNode.GetToken().Literal)
		assert.NotNil(t, l.Head.Next)
		assert.Equal(t, int32(1), l.Head.Next.LNode.GetValue())
		assert.NotNil(t, l.Head.Next.Next)
		assert.Equal(t, int32(2), l.Head.Next.Next.LNode.GetValue())
	default:
		assert.Fail(t, "Invalid type")
	}
}

func TestTwoStatements(t *testing.T) {
	input := `(1 2 3)
(+ 1 2)`
	lexer := lexer.NewLexer(input)
	result, err := ParseProgram(lexer)

	assert.Nil(t, err)
	assert.IsType(t, &ast.Program{}, result)
	assert.Equal(t, 2, len(result.ListStatements))

	assert.IsType(t, &ast.ListExpression{}, result.ListStatements[0])
	assert.IsType(t, &ast.ListExpression{}, result.ListStatements[1])
}

func TestNestedFunctionCalls(t *testing.T) {
	input := `(+ 1 (+ 2 3))`
	lexer := lexer.NewLexer(input)
	result, err := ParseProgram(lexer)

	assert.Nil(t, err)
	assert.IsType(t, &ast.Program{}, result)
	assert.Equal(t, 1, len(result.ListStatements))

	s := result.ListStatements[0]

	switch l := s.(type) {
	case *ast.ListExpression:
		assert.Equal(t, 3, l.Size)
		assert.IsType(t, &ast.ListExpression{}, l.Head.Next.Next.LNode)
	default:
		assert.Fail(t, "Invalid type")
	}
}
