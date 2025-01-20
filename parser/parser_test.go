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

	assert.IsType(t, ast.ListExpression{}, statement)
	assert.Equal(t, 3, len(statement.Elements))

	initElement := int32(1)
	for _, el := range statement.Elements {
		assert.Equal(t, initElement, el.GetValue())
		initElement += 1
	}
}

func TestParseNestedListsOfNumbers(t *testing.T) {
	input := `(1 1.1 (2 3))`
	lexer := lexer.NewLexer(input)
	result, err := ParseProgram(lexer)

	assert.Nil(t, err)
	assert.IsType(t, result, &ast.Program{})
	assert.Equal(t, 1, len(result.ListStatements))
}

func TestParseSumOfIntegers(t *testing.T) {
	input := `(+ 1 2)`
	lexer := lexer.NewLexer(input)
	result, err := ParseProgram(lexer)

	assert.Nil(t, err)
	assert.IsType(t, &ast.Program{}, result)
	assert.Equal(t, 1, len(result.ListStatements))

	statement := result.ListStatements[0]

	assert.Equal(t, "(", statement.TokenLiteral())
	assert.Equal(t, 1, len(statement.Elements))

	callExpr := statement.Elements[0]

	assert.IsType(t, &ast.CallExpression{}, callExpr)
	assert.Equal(t, "+", callExpr.TokenLiteral())
	assert.Equal(t, "(+ 1 2)", callExpr.String())
}

func TestTwoStatements(t *testing.T) {
	input := `(1 2 3)
(+ 1 2)`
	lexer := lexer.NewLexer(input)
	result, err := ParseProgram(lexer)

	assert.Nil(t, err)
	assert.IsType(t, &ast.Program{}, result)
	assert.Equal(t, 2, len(result.ListStatements))

	assert.Equal(t, "(", result.ListStatements[0].TokenLiteral())
	assert.Equal(t, 0, result.ListStatements[0].Token.Pos.Col)
	assert.Equal(t, 0, result.ListStatements[0].Token.Pos.Ch)
	assert.Equal(t, 3, len(result.ListStatements[0].Elements))
	assert.Equal(t, int32(1), result.ListStatements[0].Elements[0].GetValue())
	assert.Equal(t, int32(2), result.ListStatements[0].Elements[1].GetValue())
	assert.Equal(t, int32(3), result.ListStatements[0].Elements[2].GetValue())

	assert.Equal(t, "(", result.ListStatements[1].TokenLiteral())
	assert.Equal(t, 1, result.ListStatements[1].Token.Pos.Col)
	assert.Equal(t, 0, result.ListStatements[1].Token.Pos.Ch)
	assert.Equal(t, 1, len(result.ListStatements[1].Elements))
	assert.Equal(t, "(+ 1 2)", result.ListStatements[1].Elements[0].String())
}

func TestNestedFunctionCalls(t *testing.T) {
	input := `(+ 1 (+ 2 3))`
	lexer := lexer.NewLexer(input)
	result, err := ParseProgram(lexer)

	assert.Nil(t, err)
	assert.IsType(t, &ast.Program{}, result)
	assert.Equal(t, 1, len(result.ListStatements))

	listElemen := result.ListStatements[0]
	callExpr := listElemen.Elements[0]

	assert.IsType(t, &ast.CallExpression{}, callExpr)
	// TODO: (+ 2 3) should be a single Node instead of a list node with
	//       a single Node
	assert.Equal(t, "(+ 1 ((+ 2 3)))", callExpr.String())
}
