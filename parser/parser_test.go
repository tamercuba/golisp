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
	assert.IsType(t, result, &ast.Program{})
	assert.Equal(t, len(result.ListStatements), 1)

	statement := result.ListStatements[0]

	assert.IsType(t, statement, ast.ListExpression{})
	assert.Equal(t, len(statement.Elements), 3)

	initElement := int32(1)
	for _, el := range statement.Elements {
		assert.Equal(t, el.GetValue(), initElement)
		initElement += 1
	}
}

func TestParseSumOfIntegers(t *testing.T) {
	input := `(+ 1 2)`
	lexer := lexer.NewLexer(input)
	result, err := ParseProgram(lexer)

	assert.Nil(t, err)
	assert.IsType(t, result, &ast.Program{})
	assert.Equal(t, len(result.ListStatements), 1)

	statement := result.ListStatements[0]

	assert.Equal(t, statement.TokenLiteral(), "(")
	assert.Equal(t, len(statement.Elements), 1)

	callExpr := statement.Elements[0]

	assert.IsType(t, callExpr, &ast.CallExpression{})
	assert.Equal(t, callExpr.TokenLiteral(), "+")
	assert.Equal(t, callExpr.String(), "(+ 1 2)")
}
