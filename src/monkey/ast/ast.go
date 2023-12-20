package ast

import (
	"bytes"
	"writeingo/src/monkey/token"
)

type Node interface {
	TokenLiteral() string
	String() string
}

// Expressions produce values:
//
//	add(5, 5)
//	5
type Expression interface {
	Node
	expressionNode()
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string {
	return i.Value
}

// Program will serve as the root node of every AST our parser produces.
// Every valid Monkey program is a series of statements.
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString((s.String()))
	}

	return out.String()
}
