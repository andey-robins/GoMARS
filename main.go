package main

import (
	"fmt"

	"github.com/andey-robins/GoMARS/vm"
)

func main() {
	imp := `
	MOV 0, 1
	`
	dwarf := `
	ADD #4, 3        ; execution begins here
	MOV 2, @2
	JMP -2
	DAT #0, #0
	`
	core := vm.NewVM(64, dwarf, imp)
	core.Display()
	// Tick only four times to visually check if its working
	for i := 0; i < 64; i++ {
		fmt.Println("---------------")
		core.Tick()
		core.Display()
	}
}
