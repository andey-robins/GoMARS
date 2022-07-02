package ast

// Included in this File:
// DataStatement
// MoveStatement
// SaveStatement
// LoadStatement
// NopStatement

import (
	"bytes"

	"github.com/andey-robins/GoMARS/token"
)

type DataStatement struct {
	Token  token.Token
	AField Expression
	BField Expression
}

func (ds *DataStatement) statementNode()       {}
func (ds *DataStatement) TokenLiteral() string { return ds.Token.Literal }
func (ds *DataStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ds.TokenLiteral() + " ")
	out.WriteString(ds.AField.String() + ", ")
	out.WriteString(ds.BField.String())

	return out.String()
}

type MoveStatement struct {
	Token  token.Token
	AField Expression
	BField Expression
}

func (ms *MoveStatement) statementNode()       {}
func (ms *MoveStatement) TokenLiteral() string { return ms.Token.Literal }
func (ms *MoveStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ms.TokenLiteral() + " ")
	out.WriteString(ms.AField.String() + ", ")
	out.WriteString(ms.BField.String())

	return out.String()
}

type SaveStatement struct {
	Token  token.Token
	AField Expression
	BField Expression
}

func (ss *SaveStatement) statementNode()       {}
func (ss *SaveStatement) TokenLiteral() string { return ss.Token.Literal }
func (ss *SaveStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ss.TokenLiteral() + " ")
	out.WriteString(ss.AField.String() + ", ")
	out.WriteString(ss.BField.String())

	return out.String()
}

type LoadStatement struct {
	Token  token.Token
	AField Expression
	BField Expression
}

func (ls *LoadStatement) statementNode()       {}
func (ls *LoadStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LoadStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.AField.String() + ", ")
	out.WriteString(ls.BField.String())

	return out.String()
}

type NopStatement struct {
	Token  token.Token
	AField Expression
	BField Expression
}

func (ns *NopStatement) statementNode()       {}
func (ns *NopStatement) TokenLiteral() string { return ns.Token.Literal }
func (ns *NopStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ns.TokenLiteral())

	return out.String()
}
