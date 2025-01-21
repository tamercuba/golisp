package ast

import (
	"fmt"
	lx "github.com/tamercuba/golisp/lexer"
	"strings"
)

type ListNode struct {
	LNode Node
	Next  *ListNode
}

type ListExpression struct {
	token lx.Token
	Head  *ListNode
	Size  int
}

func NewListExpression(token lx.Token) *ListExpression {
	return &ListExpression{
		token: token,
		Head:  nil,
		Size:  0,
	}
}

func (le *ListExpression) GetToken() lx.Token {
	return le.token
}

func (le *ListExpression) String() string {
	if le.Head == nil {
		return ""
	}

	currentNode := le.Head
	values := []string{}
	for {
		if currentNode == nil {
			return fmt.Sprintf("(%s)", strings.Join(values, " "))
		}

		values = append(values, currentNode.LNode.String())
		currentNode = currentNode.Next
	}
}

func (le *ListExpression) Append(node Node) {
	newNode := &ListNode{LNode: node, Next: nil}
	if le.Head == nil {
		le.Head = newNode
		le.Size += 1
		return
	}
	currentNode := le.Head
	for {
		if currentNode.Next == nil {
			currentNode.Next = newNode
			le.Size++
			return
		}

		currentNode = currentNode.Next
	}
}

func (le *ListExpression) GetValue() any {
	result := make([]string, le.Size-1)
	currentNode := le.Head

	for {
		if currentNode == nil {
			return result
		}

		result = append(result, currentNode.LNode.String())
		currentNode = currentNode.Next
	}
}
