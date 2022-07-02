package lexer

import (
	"testing"

	"github.com/andey-robins/GoMARS/token"
)

func TestNextToken(t *testing.T) {
	input := `
		DIV
		MUL
		SLT
		DJN
		JMP
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.DIVIDE, "DIV"},
		{token.MULTIPLY, "MUL"},
		{token.SKIP_LOWER, "SLT"},
		{token.DECREMENT_JUMP, "DJN"},
		{token.JUMP, "JMP"},
	}

	lexer := New(input)

	for i, test := range tests {
		tok := lexer.NextToken()

		if tok.Type != test.expectedType {
			t.Fatalf("tests[%d] = tokenType wrong. expected=%q, got=%q", i, test.expectedType, tok.Type)
		}

		if tok.Literal != test.expectedLiteral {
			t.Fatalf("tests[%d] = tokenLiteral wrong. expected=%q, got=%q", i, test.expectedLiteral, tok.Literal)
		}
	}
}

func TestCommentTokenization(t *testing.T) {
	input := `
	SPL ; execution starts here
	MOV 
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.SPLIT, "SPL"},
		{token.MOVE, "MOV"},
	}

	lexer := New(input)

	for i, test := range tests {
		tok := lexer.NextToken()

		if tok.Type != test.expectedType {
			t.Fatalf("tests[%d] = tokenType wrong. expected=%q, got=%q", i, test.expectedType, tok.Type)
		}

		if tok.Literal != test.expectedLiteral {
			t.Fatalf("tests[%d] = tokenLiteral wrong. expected=%q, got=%q", i, test.expectedLiteral, tok.Literal)
		}
	}
}

func TestWorkingWarrior(t *testing.T) {
	input := `
	SPL	0 		; execution starts here
	MOV 0, 1
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.SPLIT, "SPL"},
		{token.INTEGER, "0"},
		{token.MOVE, "MOV"},
		{token.INTEGER, "0"},
		{token.COMMA, ","},
		{token.INTEGER, "1"},
	}

	lexer := New(input)

	for i, test := range tests {
		tok := lexer.NextToken()

		if tok.Type != test.expectedType {
			t.Fatalf("tests[%d] = tokenType wrong. expected=%q, got=%q", i, test.expectedType, tok.Type)
		}

		if tok.Literal != test.expectedLiteral {
			t.Fatalf("tests[%d] = tokenLiteral wrong. expected=%q, got=%q", i, test.expectedLiteral, tok.Literal)
		}
	}
}

func TestComplexWarrior(t *testing.T) {
	input := `
	ADD #4, 3   	; execution begins here
	MOV 2, @2
	JMP -2			; an extra comment
	DAT #0, #0
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ADD, "ADD"},
		{token.HASH, "#"},
		{token.INTEGER, "4"},
		{token.COMMA, ","},
		{token.INTEGER, "3"},
		{token.MOVE, "MOV"},
		{token.INTEGER, "2"},
		{token.COMMA, ","},
		{token.AT, "@"},
		{token.INTEGER, "2"},
		{token.JUMP, "JMP"},
		{token.MINUS, "-"},
		{token.INTEGER, "2"},
		{token.DATA, "DAT"},
		{token.HASH, "#"},
		{token.INTEGER, "0"},
		{token.COMMA, ","},
		{token.HASH, "#"},
		{token.INTEGER, "0"},
	}

	lexer := New(input)

	for i, test := range tests {
		tok := lexer.NextToken()

		if tok.Type != test.expectedType {
			t.Fatalf("tests[%d] = tokenType wrong. expected=%q, got=%q", i, test.expectedType, tok.Type)
		}

		if tok.Literal != test.expectedLiteral {
			t.Fatalf("tests[%d] = tokenLiteral wrong. expected=%q, got=%q", i, test.expectedLiteral, tok.Literal)
		}
	}
}
