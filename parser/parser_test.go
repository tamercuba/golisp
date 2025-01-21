package parser

import (
	"fmt"
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
		assert.Fail(t, fmt.Sprintf("Invalid type: %+v", list))
	}
}

func TestParseNestedListsOfNumbers(t *testing.T) {
	input := `(1 1.1 (2 "3"))`
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
			assert.Equal(t, "3", ll.Head.Next.LNode.GetValue())
			assert.Nil(t, ll.Head.Next.Next)
		default:
			assert.Fail(t, "Invalid type")
		}
		assert.Nil(t, l.Head.Next.Next.Next)

	default:
		assert.Fail(t, fmt.Sprintf("Invalid type: %+v", l))
	}
}

func TestParseSumOfIntegers(t *testing.T) {
	input := `(+ 1 x)`
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
		assert.Equal(t, "x", l.Head.Next.Next.LNode.GetValue())
	default:
		assert.Fail(t, fmt.Sprintf("Invalid type: %+v", l))
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
		assert.Fail(t, fmt.Sprintf("Invalid type: %+v", l))
	}
}

func TestLetDeclarationNode(t *testing.T) {
	input := `(let x 10)`
	l := lexer.NewLexer(input)
	r, err := ParseProgram(l)

	assert.Equal(t, 1, len(r.ListStatements))
	s := r.ListStatements[0]

	assert.Nil(t, err)
	assert.IsType(t, &ast.LetDeclaration{}, s)
	assert.Equal(t, "let", s.GetToken().Literal)

	switch v := s.(type) {
	case *ast.LetDeclaration:
		assert.Equal(t, "x", v.Name.String())
		assert.Equal(t, int32(10), v.Value.GetValue())
		assert.Equal(t, "(let x 10)", fmt.Sprintf("%v", v))
	default:
		assert.Fail(t, fmt.Sprintf("Invalid type: %+v", v))
	}
}

func TestLetDeclarationWithNestedList(t *testing.T) {
	input := `(let x (+ 1 2))`
	l := lexer.NewLexer(input)
	r, err := ParseProgram(l)

	assert.Equal(t, 1, len(r.ListStatements))
	s := r.ListStatements[0]

	assert.Nil(t, err)
	assert.IsType(t, &ast.LetDeclaration{}, s)
	assert.Equal(t, "let", s.GetToken().Literal)

	switch v := s.(type) {
	case *ast.LetDeclaration:
		assert.Equal(t, "x", v.Name.String())
		assert.IsType(t, &ast.ListExpression{}, v.Value)
		assert.Equal(t, "(let x (+ 1 2))", fmt.Sprintf("%v", v))
	default:
		assert.Fail(t, fmt.Sprintf("Invalid type: %+v", v))
	}
}

func TestLetWithoutValue(t *testing.T) {
	input := `(let x)`
	l := lexer.NewLexer(input)
	assert.Panics(t, func() {
		ParseProgram(l)
	})
}

func TestLetWithStringValue(t *testing.T) {
	input := `(let x "a")`
	l := lexer.NewLexer(input)
	r, err := ParseProgram(l)

	assert.Equal(t, 1, len(r.ListStatements))
	s := r.ListStatements[0]

	assert.Nil(t, err)
	assert.IsType(t, &ast.LetDeclaration{}, s)
	assert.Equal(t, "let", s.GetToken().Literal)

	switch v := s.(type) {
	case *ast.LetDeclaration:
		assert.Equal(t, "x", v.Name.String())
		assert.Equal(t, "a", v.Value.GetValue())
		assert.Equal(t, "(let x 'a')", fmt.Sprintf("%v", v))
	default:
		assert.Fail(t, fmt.Sprintf("Invalid type: %+v", v))
	}

}

func TestDefunDeclarationNode(t *testing.T) {
	input := `(defun x (y z) (+ y z))`
	l := lexer.NewLexer(input)
	r, err := ParseProgram(l)

	assert.Equal(t, 1, len(r.ListStatements))
	s := r.ListStatements[0]

	assert.Nil(t, err)
	assert.IsType(t, &ast.FunctionDeclaration{}, s)
	assert.Equal(t, "defun", s.GetToken().Literal)

	switch v := s.(type) {
	case *ast.FunctionDeclaration:
		assert.Equal(t, "x", v.Name.String())
		assert.IsType(t, "y", v.Args[0].String())
		assert.IsType(t, "z", v.Args[1].String())
		assert.IsType(t, &ast.ListExpression{}, v.Body)
	default:
		assert.Fail(t, fmt.Sprintf("Invalid type: %+v", v))
	}

}

func TestLetDeclarationWithInvalidName(t *testing.T) {
	input := `(let @ 1)`
	l := lexer.NewLexer(input)
	assert.Panics(t, func() {
		ParseProgram(l)
	})
}
