package ast

import "writeingo/src/monkey/token"

type Node interface {
	TokenLiteral() string
}

// Statement doesn’t produce a value:
//
//	return 5;
//	let x = 5;
type Statement interface {
	Node
	statementNode()
}

// Expressions produce values:
//
//	add(5, 5)
//	5
type Expression interface {
	Node
	expressionNode()
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

type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
