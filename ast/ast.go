package ast

import "github.com/0xedb/lang/token"

// Node is an AST node
type Node interface {
	TokenLiteral() string
}

// Statement is an AST statement
type Statement interface {
	Node
	statementNode()
}

// Expression is an AST expression
type Expression interface {
	Node
	expressionNode()
}

// Program is an AST program representation
type Program struct {
	Statements []Statement
}

// TokenLiteral returns program token literal
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}

	return ""
}

// LetStatement is a let statment node
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

// TokenLiteral returns let token literal
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// Identifier is an AST identifier
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode() {}

// TokenLiteral is an identifier token literal
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
