package ast

import (
	"testing"

	"github.com/andey-robins/GoMARS/token"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&MoveStatement{
				Token: token.Token{Type: token.MOVE, Literal: "MOV"},
				AField: &DirectExpression{
					Token: token.Token{Type: token.INTEGER, Literal: "0"},
				},
				BField: &DirectExpression{
					Token: token.Token{Type: token.INTEGER, Literal: "1"},
					Value: IntegerLiteral{
						Token: token.Token{Type: token.INTEGER, Literal: "1"},
						Value: 1,
					},
				},
			},
		},
	}

	if program.String() != "MOV 0, 1" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}

func TestStringTwo(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&MoveStatement{
				Token: token.Token{Type: token.MOVE, Literal: "ADD"},
				AField: &IndirectExpression{
					Token: token.Token{Type: token.INTEGER, Literal: "3"},
					Value: IntegerLiteral{
						Token: token.Token{Type: token.INTEGER, Literal: "3"},
						Value: 3},
				},
				BField: &ImmediateExpression{
					Token: token.Token{Type: token.INTEGER, Literal: "4"},
					Value: IntegerLiteral{
						Token: token.Token{Type: token.INTEGER, Literal: "4"},
						Value: 4},
				},
			},
		},
	}

	if program.String() != "ADD @3, #4" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
