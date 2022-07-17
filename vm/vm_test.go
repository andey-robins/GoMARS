package vm

import (
	"math/rand"
	"testing"
)

func TestImpExecution(t *testing.T) {
	imp := `
	MOV 0, 1
	`

	// standardize random seed for deterministic tests
	rand.Seed(1)

	core := NewVM(16, imp, imp)
	for i := 0; i < 7; i++ {
		core.Tick()
	}

	if core.Winner != "Draw" {
		t.Errorf("TestImpExecution failed. got=%v exp=%v", core.Winner, "APlayer")
	}
}

func TestKill(t *testing.T) {
	imp := `
	MOV 0, 1
	`
	suicide := `
	DAT 0, 0
	`

	// standardize random seed for deterministic tests
	rand.Seed(9)
	core := NewVM(16, imp, suicide)
	for core.Tick() {
		// core.Display()
	}
	if core.Winner != "APlayer" {
		t.Errorf("wrong winner. got=%s, exp=%s", core.Winner, "APlayer")
	}
}

func TestDraw(t *testing.T) {
	imp := `
	MOV 0, 1
	`

	// standardize random seed for deterministic tests
	rand.Seed(1)

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

	// standardize random seed for deterministic tests
	rand.Seed(3)
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
