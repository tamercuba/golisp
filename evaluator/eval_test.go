package evaluator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tamercuba/golisp/evaluator/object"
	lx "github.com/tamercuba/golisp/lexer"
	pr "github.com/tamercuba/golisp/parser"
	"github.com/tamercuba/golisp/parser/ast"
)

func TestEvalDefineVar(t *testing.T) {
	input := `(define x 10)`
	l := lx.NewLexer(input)
	p, _ := pr.ParseProgram(l)

	e := NewEvaluator()
	r := e.EvalProgram(p)
	x := e.env.Get("x")

	assert.Equal(t, object.NIL_TYPE, r.Type())

	assert.IsType(t, &ast.IntLiteral{}, x)
	assert.Equal(t, int32(10), x.GetValue())
	assert.Nil(t, e.env.outerScope)
}

func TestEvalSumOfIntegers(t *testing.T) {
	input := `(+ 1 2)`
	l := lx.NewLexer(input)
	p, _ := pr.ParseProgram(l)

	e := NewEvaluator()
	r := e.EvalProgram(p)

	assert.Equal(t, object.LIST_TYPE, r.Type())
	assert.Equal(t, "(+ 1 2)", r.Inspect())

}
