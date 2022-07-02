package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

// token literals
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	NOOP    = "NOP"

	DATA = "DAT"
	MOVE = "MOV"

	ADD      = "ADD"
	SUBTRACT = "SUB"
	MULTIPLY = "MUL"
	DIVIDE   = "DIV"
	MODULUS  = "MOD"

	JUMP           = "JMP"
	JUMP_ZERO      = "JMZ"
	JUMP_NOT_ZERO  = "JMN"
	DECREMENT_JUMP = "DJN"

	SPLIT = "SPL"

	COMPARE        = "CMP"
	SKIP_EQUAL     = "SEQ"
	SKIP_NOT_EQUAL = "SNE"
	SKIP_LOWER     = "SLT"
	LOADP          = "LDP"
	SAVEP          = "STP"

	HASH  = "HASH"
	AT    = "AT"
	MINUS = "MINUS"
	BLING = "BLING"

	// TODO: Extend this later to include the more advanced addressing modes

	INTEGER   = "INT"
	COMMA     = "COMMA"
	SEMICOLON = "SEMICOLON"
	COMMENT   = "COMMENT"
)

var keywords = map[string]TokenType{
	"DAT": DATA,
	"MOV": MOVE,
	"ADD": ADD,
	"SUB": SUBTRACT,
	"MUL": MULTIPLY,
	"DIV": DIVIDE,
	"MOD": MODULUS,
	"JMP": JUMP,
	"JMZ": JUMP_ZERO,
	"JMN": JUMP_NOT_ZERO,
	"DJN": DECREMENT_JUMP,
	"SPL": SPLIT,
	"CMP": COMPARE,
	"SEQ": SKIP_EQUAL,
	"SNE": SKIP_NOT_EQUAL,
	"SLT": SKIP_LOWER,
	"LDP": LOADP,
	"STP": SAVEP,
	"NOP": NOOP,
}

func LookupIdentifier(identifier string) TokenType {
	if tok, ok := keywords[identifier]; ok {
		return tok
	}
	return ILLEGAL
}
