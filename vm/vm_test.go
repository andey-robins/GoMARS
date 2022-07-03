package vm

import (
	"testing"
)

func TestImpExecution(t *testing.T) {
	imp := `
	MOV 0, 1
	`
	core := NewVM(16, imp, imp)
	for i := 0; i < 7; i++ {
		core.Tick()
	}

	impCell := Cell{
		Operation: MOV,
		AField: Field{
			Mode:  DIRECT,
			Value: 0,
		},
		BField: Field{
			Mode:  DIRECT,
			Value: 1,
		},
	}

	for i, cell := range core.Memory {
		if !(cell.Operation == MOV && testEqualField(impCell.AField, cell.AField) && testEqualField(impCell.BField, cell.BField)) {
			t.Errorf("memory cell %v is wrong.", i)
		}
	}
}

func TestKill(t *testing.T) {
	imp := `
	MOV 0, 1
	`
	suicide := `
	DAT 0, 0
	`

	core := NewVM(16, imp, suicide)
	for core.Tick() {
	}
	if core.Winner != "APlayer" {
		t.Errorf("wrong winner. got=%s, exp=%s", core.Winner, "APlayer")
	}
}

func TestDraw(t *testing.T) {
	imp := `
	MOV 0, 1
	`
	core := NewVM(8, imp, imp)
	for core.Tick() {
	}

	if core.Winner != "Draw" {
		t.Errorf("wrong winner. got=%s, exp=%s", core.Winner, "Draw")
	}

}

func TestRealDuel(t *testing.T) {
	dwarf := `
	ADD #4, 3        ; execution begins here
	MOV 2, @2
	JMP -2
	DAT #0, #0
	`
	imp := `
	MOV 0, 1
	`
	core := NewVM(32, dwarf, imp)
	for core.Tick() {
	}

	if core.Winner != "Draw" {
		t.Errorf("wrong winner. got=%s, exp=%s", core.Winner, "Draw")
	}
}

func testEqualField(f1, f2 Field) bool {
	return f1.Mode == f2.Mode && f1.Value == f2.Value
}
