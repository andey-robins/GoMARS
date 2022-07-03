package main

import (
	"fmt"

	"github.com/andey-robins/GoMARS/vm"
)

func main() {
	imp := `
	MOV 0, 1
	`
	core := vm.NewVM(16, imp, imp)

	core.Display()

	// Tick only four times to visually check if its working
	for i := 0; i < 4; i++ {
		fmt.Println("---------------")
		core.Tick()
		core.Display()
	}
}
