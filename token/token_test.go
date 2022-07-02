package token

import (
	"testing"
)

func TestIdentifierLookup(t *testing.T) {
	tests := []struct {
		identifier        string
		expectedTokenType TokenType
	}{
		{"ADD", ADD},
		{"MOV", MOVE},
		{"JMP", JUMP},
		{"SPL", SPLIT},
	}

	for i, test := range tests {
		if test.expectedTokenType != LookupIdentifier(test.identifier) {
			t.Fatalf("[%v] - expected token type=%v. got=%v", i, test.expectedTokenType, LookupIdentifier(test.identifier))
		}
	}
}
