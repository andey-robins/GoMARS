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
