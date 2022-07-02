package ast

import (
	"bytes"

	"github.com/andey-robins/GoMARS/token"
)

// Included in this file:
// JumpStatement
// JumpZeroStatement
// JumpNonzeroStatement
// DecrementJumpStatement
// SplitStatement

type JumpStatement struct {
	Token  token.Token
	AField Expression
}

func (js *JumpStatement) statementNode()       {}
func (js *JumpStatement) TokenLiteral() string { return js.Token.Literal }
func (js *JumpStatement) String() string {
	var out bytes.Buffer

	out.WriteString(js.TokenLiteral() + " ")
	out.WriteString(js.AField.String())

	return out.String()
}

type JumpZeroStatement struct {
	Token  token.Token
	AField Expression
	BField Expression
}

func (jzs *JumpZeroStatement) statementNode()       {}
func (jzs *JumpZeroStatement) TokenLiteral() string { return jzs.Token.Literal }
func (jzs *JumpZeroStatement) String() string {
	var out bytes.Buffer

	out.WriteString(jzs.TokenLiteral() + " ")
	out.WriteString(jzs.AField.String() + ", ")
	out.WriteString(jzs.BField.String())

	return out.String()
}

type JumpNonzeroStatement struct {
	Token  token.Token
	AField Expression
	BField Expression
}

func (jns *JumpNonzeroStatement) statementNode()       {}
func (jns *JumpNonzeroStatement) TokenLiteral() string { return jns.Token.Literal }
func (jns *JumpNonzeroStatement) String() string {
	var out bytes.Buffer

	out.WriteString(jns.TokenLiteral() + " ")
	out.WriteString(jns.AField.String() + ", ")
	out.WriteString(jns.BField.String())

	return out.String()
}

type DecrementJumpStatement struct {
	Token  token.Token
	AField Expression
	BField Expression
}

func (djs *DecrementJumpStatement) statementNode()       {}
func (djs *DecrementJumpStatement) TokenLiteral() string { return djs.Token.Literal }
func (djs *DecrementJumpStatement) String() string {
	var out bytes.Buffer

	out.WriteString(djs.TokenLiteral() + " ")
	out.WriteString(djs.AField.String() + ", ")
	out.WriteString(djs.BField.String())

	return out.String()
}

type SplitStatement struct {
	Token  token.Token
	AField Expression
	BField Expression
}

func (ss *SplitStatement) statementNode()       {}
func (ss *SplitStatement) TokenLiteral() string { return ss.Token.Literal }
func (ss *SplitStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ss.TokenLiteral() + " ")
	out.WriteString(ss.AField.String() + ", ")
	out.WriteString(ss.BField.String())

	return out.String()
}
