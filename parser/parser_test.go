package parser

import (
	"testing"

	"github.com/andey-robins/GoMARS/ast"
	"github.com/andey-robins/GoMARS/lexer"
)

func TestImpParsing(t *testing.T) {
	input := `
	MOV 0, 1
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 1 statement. got%d", len(program.Statements))
	}

	movStmt, ok := program.Statements[0].(*ast.MoveStatement)
	if !ok {
		t.Fatalf("statement 0 not *ast.MoveStatement. got=%T", program.Statements[0])
	}

	aField, ok := movStmt.AField.(*ast.DirectExpression)
	if !ok {
		t.Fatalf("afield not *ast.DirectExpression. got=%T", movStmt.AField)
	}
	if aField.Value.Value != 0 {
		t.Fatalf("afield value not 0. got=%v, gotType=%T", aField.Value.Value, aField.Value.Value)
	}

	bField, ok := movStmt.BField.(*ast.DirectExpression)
	if !ok {
		t.Fatalf("bfield not *ast.DirectExpression. got=%T", movStmt.BField)
	}
	if bField.Value.Value != 1 {
		t.Fatalf("bfield value not 1. got=%v, gotType=%T", bField.Value.Value, bField.Value.Value)
	}
}

func TestDwarfParsing(t *testing.T) {
	input := `
	ADD #4, 3        ; execution begins here
	MOV 2, @2
	JMP -2
	DAT #0, #0
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 4 {
		t.Fatalf("program.Statements does not contain 4 statements. got%d", len(program.Statements))
	}

	tests := []string{
		"ADD #4, 3",
		"MOV 2, @2",
		"JMP -2",
		"DAT #0, #0",
	}

	for i, expectedString := range tests {
		if expectedString != program.Statements[i].String() {
			t.Errorf("[%v] - expectedString='%v'. got='%v'", i, expectedString, program.Statements[i].String())
		}
	}
}

func TestParsingAllStatements(t *testing.T) {
	input := `
	DAT #1, #2
	MOV 1, 2
	ADD 1, 2
	SUB @1, @2
	MUL 1, 2
	DIV 1, 2
	MOD 1, 2		; here are some comments
	JMP 10
	JMZ 10, 1
	JMN 10, 1
	DJN 10, 10
	SPL 5			; i put them in here to test the parser/lexer
	CMP #4, #5
	SEQ #1, #2
	SNE #1, #2
	SLT #1, #2
	LDP 1, @4		; they shouldn't impact the parsing at all
	STP 1, @5
	NOP
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 19 {
		t.Fatalf("program.Statements does not contain 19 statements. got%d", len(program.Statements))
	}

	tests := []string{
		"DAT #1, #2",
		"MOV 1, 2",
		"ADD 1, 2",
		"SUB @1, @2",
		"MUL 1, 2",
		"DIV 1, 2",
		"MOD 1, 2",
		"JMP 10",
		"JMZ 10, 1",
		"JMN 10, 1",
		"DJN 10, 10",
		"SPL 5",
		"CMP #4, #5",
		"SEQ #1, #2",
		"SNE #1, #2",
		"SLT #1, #2",
		"LDP 1, @4",
		"STP 1, @5",
		"NOP",
	}

	for i, expectedString := range tests {
		if expectedString != program.Statements[i].String() {
			t.Errorf("[%v] - expectedString='%v'. got='%v'", i, expectedString, program.Statements[i].String())
		}
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf("parser had %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}

	t.FailNow()
}
