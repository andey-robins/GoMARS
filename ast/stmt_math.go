package ast

// Included in this file:
// AddStatement
// SubtractStatement
// MultiplyStatement
// DivideStatement
// ModulusStatement

import (
	"bytes"

	"github.com/andey-robins/GoMARS/token"
)

type AddStatement struct {
	Token  token.Token
	AField Expression
	BField Expression
}

func (as *AddStatement) statementNode()       {}
func (as *AddStatement) TokenLiteral() string { return as.Token.Literal }
func (as *AddStatement) String() string {
	var out bytes.Buffer

	out.WriteString(as.TokenLiteral() + " ")
	out.WriteString(as.AField.String() + ", ")
	out.WriteString(as.BField.String())

	return out.String()
}

type SubtractStatement struct {
	Token  token.Token
	AField Expression
	BField Expression
}

func (ss *SubtractStatement) statementNode()       {}
func (ss *SubtractStatement) TokenLiteral() string { return ss.Token.Literal }
func (ss *SubtractStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ss.TokenLiteral() + " ")
	out.WriteString(ss.AField.String() + ", ")
	out.WriteString(ss.BField.String())

	return out.String()
}

type MultiplyStatement struct {
	Token  token.Token
	AField Expression
	BField Expression
}

func (ms *MultiplyStatement) statementNode()       {}
func (ms *MultiplyStatement) TokenLiteral() string { return ms.Token.Literal }
func (ms *MultiplyStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ms.TokenLiteral() + " ")
	out.WriteString(ms.AField.String() + ", ")
	out.WriteString(ms.BField.String())

	return out.String()
}

type DivideStatement struct {
	Token  token.Token
	AField Expression
	BField Expression
}

func (ds *DivideStatement) statementNode()       {}
func (ds *DivideStatement) TokenLiteral() string { return ds.Token.Literal }
func (ds *DivideStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ds.TokenLiteral() + " ")
	out.WriteString(ds.AField.String() + ", ")
	out.WriteString(ds.BField.String())

	return out.String()
}

type ModStatement struct {
	Token  token.Token
	AField Expression
	BField Expression
}

func (ms *ModStatement) statementNode()       {}
func (ms *ModStatement) TokenLiteral() string { return ms.Token.Literal }
func (ms *ModStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ms.TokenLiteral() + " ")
	out.WriteString(ms.AField.String() + ", ")
	out.WriteString(ms.BField.String())

	return out.String()
}
