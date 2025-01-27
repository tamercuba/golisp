package object

import (
	"fmt"
	"strings"
)

type ObjectType string

type Object interface {
	Type() ObjectType
	Inspect() string
}

const (
	INT_TYPE     ObjectType = "INT_TYPE"
	FLOAT_TYPE              = "FLOAT_TYPE"
	STRING_TYPE             = "STRING_TYPE"
	BOOLEAN_TYPE            = "BOOLEAN_TYPE"
	NIL_TYPE                = "NIL_TYPE"
	LIST_TYPE               = "LIST_TYPE"
)

type Integer struct {
	Value int32
}

func (i *Integer) Type() ObjectType {
	return INT_TYPE
}

func (i Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

type Float struct {
	Value float64
}

func (f *Float) Type() ObjectType {
	return FLOAT_TYPE
}

func (f Float) Inspect() string {
	return fmt.Sprintf("%f", f.Value)
}

type String struct {
	Value string
}

func (s *String) Type() ObjectType {
	return STRING_TYPE
}

func (s String) Inspect() string {
	return s.Value
}

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType {
	return BOOLEAN_TYPE
}

func (b Boolean) Inspect() string {
	if b.Value {
		return "true"
	} else {
		return "false"
	}
}

type Nil struct{}

func (n *Nil) Type() ObjectType {
	return NIL_TYPE
}

func (n Nil) Inspect() string {
	return "nil"
}

type List struct {
	Content []Object
}

func (l *List) Type() ObjectType {
	return LIST_TYPE
}

func (l List) Inspect() string {
	sl := []string{}

	for _, e := range l.Content {
		sl = append(sl, e.Inspect())
	}

	return fmt.Sprintf("(%s)", strings.Join(sl, " "))
}
