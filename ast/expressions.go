package ast

import (
	"bytes"
	"strconv"

	"github.com/andey-robins/GoMARS/token"
)

type ImmediateExpression struct {
	Token token.Token
	Value IntegerLiteral
}

func (ie *ImmediateExpression) expressionNode()      {}
func (ie *ImmediateExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *ImmediateExpression) String() string {
	var out bytes.Buffer

	out.WriteString("#" + ie.Value.String())
	return out.String()
}

type DirectExpression struct {
	Token token.Token
	Value IntegerLiteral
}

func (de *DirectExpression) expressionNode()      {}
func (de *DirectExpression) TokenLiteral() string { return de.Token.Literal }
func (de *DirectExpression) String() string {
	var out bytes.Buffer

	out.WriteString(de.Value.String())
	return out.String()
}

type IndirectExpression struct {
	Token token.Token
	Value IntegerLiteral
}

func (ie *IndirectExpression) expressionNode()      {}
func (ie *IndirectExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IndirectExpression) String() string {
	var out bytes.Buffer

	out.WriteString("@" + ie.Value.String())

	return out.String()
}

type IntegerLiteral struct {
	Token token.Token // the token.INTEGER token
	Value int64
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string {
	var out bytes.Buffer
	out.WriteString(strconv.Itoa(int(il.Value)))
	return out.String()
}
