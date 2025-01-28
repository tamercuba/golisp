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
	x := e.Env.Get("x")

	assert.Equal(t, object.NIL_TYPE, r.Type())

	assert.IsType(t, &ast.IntLiteral{}, x)
	assert.Equal(t, int32(10), x.GetValue())
	assert.Nil(t, e.Env.outerScope)
}

func TestEvalSumOfIntegers(t *testing.T) {
	input := `(+ 1 2)`
	l := lx.NewLexer(input)
	p, _ := pr.ParseProgram(l)

	e := NewEvaluator()
	r := e.EvalProgram(p)

	assert.Equal(t, object.INT_TYPE, r.Type())
	assert.Equal(t, "3", r.Inspect())
}

func TestEvalSumOfManyIntegers(t *testing.T) {
	input := `(+ 1 2 3 4 5 6 7 8 9 10)`
	l := lx.NewLexer(input)
	p, _ := pr.ParseProgram(l)

	e := NewEvaluator()
	r := e.EvalProgram(p)

	assert.Equal(t, object.INT_TYPE, r.Type())
	assert.Equal(t, "55", r.Inspect())
}

func TestEvalSumOfIntegersAndFloats(t *testing.T) {
	input := `(+ 1 1 1.5 2)`
	l := lx.NewLexer(input)
	p, _ := pr.ParseProgram(l)

	e := NewEvaluator()
	r := e.EvalProgram(p)

	assert.Equal(t, object.FLOAT_TYPE, r.Type())
	assert.Equal(t, "5.500000", r.Inspect())
}

func TestEvalSubOfManyIntegersAndFloats(t *testing.T) {
	input := `(- 1 2 3.5 4)`
	l := lx.NewLexer(input)
	p, _ := pr.ParseProgram(l)

	e := NewEvaluator()
	r := e.EvalProgram(p)

	assert.Equal(t, object.FLOAT_TYPE, r.Type())
	assert.Equal(t, "-8.500000", r.Inspect())
}

func TestSumOfNegativeNumbers(t *testing.T) {
	input := `(- -1 2 3.5 4)`
	l := lx.NewLexer(input)
	p, _ := pr.ParseProgram(l)

	e := NewEvaluator()
	r := e.EvalProgram(p)

	assert.Equal(t, object.FLOAT_TYPE, r.Type())
	assert.Equal(t, "-10.500000", r.Inspect())
}

func TestLambdaDeclaration(t *testing.T) {
	input := `(lambda (x y) (+ x y))`
	l := lx.NewLexer(input)
	p, _ := pr.ParseProgram(l)

	e := NewEvaluator()
	r := e.EvalProgram(p)

	assert.IsType(t, &object.Nil{}, r)
}

func TestLambdaDeclarationAndCall(t *testing.T) {
	input := `((lambda (x y) (+ x y)) 5 6)`
	l := lx.NewLexer(input)
	p, _ := pr.ParseProgram(l)

	e := NewEvaluator()
	r := e.EvalProgram(p)

	assert.Equal(t, object.INT_TYPE, r.Type())
	assert.Equal(t, "11", r.Inspect())
}

func TestLambdaCallWithDefine(t *testing.T) {
	input := `
  (define x (lambda (y z) (* z y)))
  (x 5 10)
`
	l := lx.NewLexer(input)
	p, _ := pr.ParseProgram(l)

	e := NewEvaluator()
	r := e.EvalProgram(p)

	assert.Equal(t, object.INT_TYPE, r.Type())
	assert.Equal(t, "50", r.Inspect())
}
