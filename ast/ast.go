package ast

import (
	"bytes"
)

// Node is the base node interface. All nodes in the AST should implement this interface.
type Node interface {
	TokenLiteral() string
	String() string
}

// Program will serve as the root node of every AST our parser produces.
// Every valid Monkey program is a series of statements.
type Program struct {
	Statements []Statement
}

// TokenLiteral returns the token literal of the first statement in the program.
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// String returns a string representation of the program, calling String on all the statements in the program.
func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString((s.String()))
	}

	return out.String()
}
