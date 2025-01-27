package evaluator

import (
	"errors"
	"fmt"

	"github.com/tamercuba/golisp/parser/ast"
)

type Envinronment struct {
	vars       map[string]ast.Node
	outerScope *Envinronment
}

func NewEnvironment() *Envinronment {
	return &Envinronment{
		vars:       make(map[string]ast.Node),
		outerScope: nil,
	}
}

func (e *Envinronment) NewScope() {
	ne := NewEnvironment()
	ne.outerScope = e
	*e = *ne
}

func (e *Envinronment) DropScope() {
	if e.outerScope != nil {
		*e = *(e.outerScope)
	}
}

func (e *Envinronment) Bind(name string, value ast.Node) error {
	v, ok := e.vars[name]
	if ok {
		err := fmt.Sprintf("%v already exists (%v)", name, v)
		return errors.New(err)
	}
	e.vars[name] = value
	return nil
}

func (e *Envinronment) BindGlobal(name string, value ast.Node) error {
	if e.outerScope == nil {
		return e.Bind(name, value)
	} else {
		return e.outerScope.Bind(name, value)
	}
}

func (e *Envinronment) Get(name string) ast.Node {
	if e.outerScope != nil {
		return e.outerScope.Get(name)
	} else {
		v, ok := e.vars[name]
		if ok {
			return v
		} else {
			return nil
		}
	}
}
