package ast

import (
	"bytes"

	"github.com/andey-robins/GoMARS/token"
)

// Included in this file:
// CompareStatement
// SkipStatement
// SkipNeqStatement
// SkipLowerStatement

type CompareStatement struct {
	Token  token.Token
	AField Expression
	BField Expression
}

func (cs *CompareStatement) statementNode()       {}
func (cs *CompareStatement) TokenLiteral() string { return cs.Token.Literal }
func (cs *CompareStatement) String() string {
	var out bytes.Buffer

	out.WriteString(cs.TokenLiteral() + " ")
	out.WriteString(cs.AField.String() + ", ")
	out.WriteString(cs.BField.String())

	return out.String()
}

type SkipStatement struct {
	Token  token.Token
	AField Expression
	BField Expression
}

func (ss *SkipStatement) statementNode()       {}
func (ss *SkipStatement) TokenLiteral() string { return ss.Token.Literal }
func (ss *SkipStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ss.TokenLiteral() + " ")
	out.WriteString(ss.AField.String() + ", ")
	out.WriteString(ss.BField.String())

	return out.String()
}

type SkipNeqStatement struct {
	Token  token.Token
	AField Expression
	BField Expression
}

func (sns *SkipNeqStatement) statementNode()       {}
func (sns *SkipNeqStatement) TokenLiteral() string { return sns.Token.Literal }
func (sns *SkipNeqStatement) String() string {
	var out bytes.Buffer

	out.WriteString(sns.TokenLiteral() + " ")
	out.WriteString(sns.AField.String() + ", ")
	out.WriteString(sns.BField.String())

	return out.String()
}

type SkipLowerStatement struct {
	Token  token.Token
	AField Expression
	BField Expression
}

func (sls *SkipLowerStatement) statementNode()       {}
func (sls *SkipLowerStatement) TokenLiteral() string { return sls.Token.Literal }
func (sls *SkipLowerStatement) String() string {
	var out bytes.Buffer

	out.WriteString(sls.TokenLiteral() + " ")
	out.WriteString(sls.AField.String() + ", ")
	out.WriteString(sls.BField.String())

	return out.String()
}
