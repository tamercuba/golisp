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
	case *ast.OperationNode:
		assert.Equal(t, "+", l.Name.String())
		assert.Equal(t, 2, len(l.Params))
		assert.IsType(t, &ast.IntLiteral{}, l.Params[0])
		assert.IsType(t, &ast.Symbol{}, l.Params[1])
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
	assert.IsType(t, &ast.OperationNode{}, result.ListStatements[1])
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
	case *ast.OperationNode:
		assert.Equal(t, 2, len(l.Params))
		assert.IsType(t, &ast.IntLiteral{}, l.Params[0])
		switch p1 := l.Params[1].(type) {
		case *ast.OperationNode:
			assert.Equal(t, 2, len(p1.Params))
			assert.IsType(t, &ast.IntLiteral{}, p1.Params[0])
			assert.IsType(t, &ast.IntLiteral{}, p1.Params[1])
		default:
			assert.Fail(t, fmt.Sprintf("Invalid type: %+v", l))
		}
	default:
		assert.Fail(t, fmt.Sprintf("Invalid type: %+v", l))
	}
}

func TestVarDifinitionNodeNode(t *testing.T) {
	input := `(let x 10)`
	l := lexer.NewLexer(input)
	r, err := ParseProgram(l)

	assert.Equal(t, 1, len(r.ListStatements))
	s := r.ListStatements[0]

	assert.Nil(t, err)
	assert.IsType(t, &ast.VarDifinitionNode{}, s)
	assert.Equal(t, "let", s.GetToken().Literal)

	switch v := s.(type) {
	case *ast.VarDifinitionNode:
		assert.Equal(t, "x", v.Name.String())
		assert.Equal(t, int32(10), v.Value.GetValue())
		assert.Equal(t, "(let x 10)", fmt.Sprintf("%v", v))
	default:
		assert.Fail(t, fmt.Sprintf("Invalid type: %+v", v))
	}
}

func TestVarDifinitionNodeWithNestedList(t *testing.T) {
	input := `(let x (+ 1 2))`
	l := lexer.NewLexer(input)
	r, err := ParseProgram(l)

	assert.Equal(t, 1, len(r.ListStatements))
	s := r.ListStatements[0]

	assert.Nil(t, err)
	assert.IsType(t, &ast.VarDifinitionNode{}, s)
	assert.Equal(t, "let", s.GetToken().Literal)

	switch v := s.(type) {
	case *ast.VarDifinitionNode:
		assert.Equal(t, "x", v.Name.String())
		assert.IsType(t, &ast.OperationNode{}, v.Value)
		assert.Equal(t, "(let x (+ 1 2))", fmt.Sprintf("%v", v))
	default:
		assert.Fail(t, fmt.Sprintf("Invalid type: %+v", v))
	}
}

func TestLetWithoutValue(t *testing.T) {
	input := `(let x)`
	l := lexer.NewLexer(input)
	r, err := ParseProgram(l)

	assert.Error(t, err)
	assert.Nil(t, r)
}

func TestLetWithStringValue(t *testing.T) {
	input := `(let x "a")`
	l := lexer.NewLexer(input)
	r, err := ParseProgram(l)

	assert.Equal(t, 1, len(r.ListStatements))
	s := r.ListStatements[0]

	assert.Nil(t, err)
	assert.IsType(t, &ast.VarDifinitionNode{}, s)
	assert.Equal(t, "let", s.GetToken().Literal)

	switch v := s.(type) {
	case *ast.VarDifinitionNode:
		assert.Equal(t, "x", v.Name.String())
		assert.Equal(t, "a", v.Value.GetValue())
		assert.Equal(t, "(let x 'a')", fmt.Sprintf("%v", v))
	default:
		assert.Fail(t, fmt.Sprintf("Invalid type: %+v", v))
	}
}

func TestVarDifinitionNodeWithInvalidName(t *testing.T) {
	input := `(let @ 1)`
	l := lexer.NewLexer(input)
	r, err := ParseProgram(l)

	assert.Error(t, err)
	assert.Nil(t, r)
}

func TestListWithBooleanValues(t *testing.T) {
	input := `(true false)`
	lexer := lexer.NewLexer(input)
	result, err := ParseProgram(lexer)

	assert.Nil(t, err)
	assert.IsType(t, &ast.Program{}, result)
	assert.Equal(t, 1, len(result.ListStatements))

	s := result.ListStatements[0]

	switch l := s.(type) {
	case *ast.ListExpression:
		assert.Equal(t, 2, l.Size)
		assert.Equal(t, true, l.Head.LNode.GetValue())
		assert.Equal(t, false, l.Head.Next.LNode.GetValue())
	default:
		assert.Fail(t, fmt.Sprintf("Invalid type: %+v", l))
	}
}

func TestLetWithBooleanValue(t *testing.T) {
	input := `(let x true)`
	l := lexer.NewLexer(input)
	r, err := ParseProgram(l)

	assert.Equal(t, 1, len(r.ListStatements))
	s := r.ListStatements[0]

	assert.Nil(t, err)
	assert.IsType(t, &ast.VarDifinitionNode{}, s)
	assert.Equal(t, "let", s.GetToken().Literal)

	switch v := s.(type) {
	case *ast.VarDifinitionNode:
		assert.Equal(t, "x", v.Name.String())
		assert.Equal(t, true, v.Value.GetValue())
		assert.Equal(t, "(let x true)", fmt.Sprintf("%v", v))
		assert.Equal(t, ast.LET, v.DefinitionType)
	default:
		assert.Fail(t, fmt.Sprintf("Invalid type: %+v", v))
	}
}

func TestDefineWithBooleanValue(t *testing.T) {
	input := `(define x true)`
	l := lexer.NewLexer(input)
	r, err := ParseProgram(l)

	assert.Equal(t, 1, len(r.ListStatements))
	s := r.ListStatements[0]

	assert.Nil(t, err)
	assert.IsType(t, &ast.VarDifinitionNode{}, s)
	assert.Equal(t, "define", s.GetToken().Literal)

	switch v := s.(type) {
	case *ast.VarDifinitionNode:
		assert.Equal(t, "x", v.Name.String())
		assert.Equal(t, true, v.Value.GetValue())
		assert.Equal(t, "(define x true)", fmt.Sprintf("%v", v))
		assert.Equal(t, ast.DEFINE, v.DefinitionType)
	default:
		assert.Fail(t, fmt.Sprintf("Invalid type: %+v", v))
	}
}

func TestDefineLambda(t *testing.T) {
	input := `
    (define func
      (lambda (x) (+ x x))
    )
  `
	l := lexer.NewLexer(input)
	r, err := ParseProgram(l)

	assert.Nil(t, err)

	ll := r.ListStatements[0]
	assert.IsType(t, &ast.VarDifinitionNode{}, ll)
	assert.Equal(t, "define", ll.GetToken().Literal)
	assert.IsType(t, &ast.LambdaNode{}, ll.GetValue())

	switch ln := ll.GetValue().(type) {
	case *ast.LambdaNode:
		assert.Equal(t, 1, len(ln.Args))
		assert.Equal(t, "x", ln.Args[0].GetValue())
		switch lb := ln.Body.(type) {
		case *ast.OperationNode:
			assert.Equal(t, "+", lb.Name.String())
			assert.Equal(t, 2, len(lb.Params))
			assert.Equal(t, "x", lb.Params[0].String())
			assert.Equal(t, "x", lb.Params[1].String())
		default:
			assert.Fail(t, fmt.Sprintf("Invalid type: %+v", lb))
		}
	default:
		assert.Fail(t, fmt.Sprintf("Invalid type: %+v", ln))
	}
}

func TestIntAloneInput(t *testing.T) {
	input := `1`
	l := lexer.NewLexer(input)
	r, err := ParseProgram(l)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(r.ListStatements))

	s := r.ListStatements[0]

	assert.Equal(t, "1", s.String())
	assert.Equal(t, int32(1), s.GetValue())
	assert.IsType(t, &ast.IntLiteral{}, s)
}

func TestFloatAloneInput(t *testing.T) {
	input := `1.2`
	l := lexer.NewLexer(input)
	r, err := ParseProgram(l)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(r.ListStatements))

	s := r.ListStatements[0]

	assert.Equal(t, "1.200000f", s.String())
	assert.Equal(t, 1.2, s.GetValue())
	assert.IsType(t, &ast.FloatLiteral{}, s)
}

func TestBooleanAlone(t *testing.T) {
	input := `true`
	l := lexer.NewLexer(input)
	r, err := ParseProgram(l)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(r.ListStatements))

	s := r.ListStatements[0]

	assert.Equal(t, "true", s.String())
	assert.Equal(t, true, s.GetValue())
	assert.IsType(t, &ast.Boolean{}, s)
}

func TestUnbalancedList(t *testing.T) {
	input := `(1 2`
	l := lexer.NewLexer(input)
	r, err := ParseProgram(l)

	expectedMessage := "0:0  (   not closed, expect )."

	assert.Error(t, err)
	assert.Equal(t, expectedMessage, err.Error())
	assert.Nil(t, r)
}

func TestAssignmentWithNilValue(t *testing.T) {
	input := `(let x nil)`
	l := lexer.NewLexer(input)
	r, err := ParseProgram(l)

	assert.Equal(t, 1, len(r.ListStatements))
	s := r.ListStatements[0]

	assert.Nil(t, err)
	assert.IsType(t, &ast.VarDifinitionNode{}, s)
	assert.Equal(t, "let", s.GetToken().Literal)

	switch v := s.(type) {
	case *ast.VarDifinitionNode:
		assert.Equal(t, "x", v.Name.String())
		assert.Nil(t, v.Value.GetValue())
		assert.Equal(t, "(let x nil)", fmt.Sprintf("%v", v))
	default:
		assert.Fail(t, fmt.Sprintf("Invalid type: %+v", v))
	}
}

func TestNil(t *testing.T) {
	input := `nil`
	l := lexer.NewLexer(input)
	r, err := ParseProgram(l)

	assert.Equal(t, 1, len(r.ListStatements))
	s := r.ListStatements[0]

	assert.Nil(t, err)
	assert.IsType(t, &ast.VoidNode{}, s)
	assert.Nil(t, s.GetValue())
	assert.Equal(t, "nil", s.String())
}

func TestListWithNilValue(t *testing.T) {
	input := `(1 nil 2 nil)`
	l := lexer.NewLexer(input)
	r, err := ParseProgram(l)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(r.ListStatements))
	s := r.ListStatements[0]

	assert.IsType(t, &ast.ListExpression{}, s)
}

func TestLambdaDeclaration(t *testing.T) {
	input := `(lambda (x y) (+ x y))`
	l := lexer.NewLexer(input)
	r, err := ParseProgram(l)

	assert.Nil(t, err)
	assert.IsType(t, 1, len(r.ListStatements))

	p := r.ListStatements[0]
	assert.IsType(t, &ast.LambdaNode{}, p)

	switch l := p.(type) {
	case *ast.LambdaNode:
		assert.Equal(t, "x", l.Args[0].GetValue())
		assert.Equal(t, "y", l.Args[1].GetValue())
		switch b := l.Body.(type) {
		case *ast.OperationNode:
			assert.Equal(t, 2, len(b.Params))
			assert.Equal(t, "+", b.Name.String())
		default:
			assert.Fail(t, fmt.Sprintf("Invalid type: %+v", l))
		}
	default:
		assert.Fail(t, fmt.Sprintf("Invalid type: %+v", l))
	}
}

func TestLambdaWithWrongSyntax(t *testing.T) {
	input := `(lambda x (+ x x))`
	l := lexer.NewLexer(input)
	r, err := ParseProgram(l)

	expectedMessage := "0:10  (   isn't a valid function argument, should be a list of symbols"
	assert.Error(t, err)
	assert.Equal(t, expectedMessage, err.Error())
	assert.Nil(t, r)
}
