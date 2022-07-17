package vm

import (
	"fmt"
	"math/rand"
)

type VirtualMachine struct {
	Memory    []Cell
	APointers *ProcessQueue
	BPointers *ProcessQueue
	AProgram  string
	BProgram  string
	Winner    string
	Turn      int
}

type Cell struct {
	Operation OpCode
	AField    Field
	BField    Field
}

type Field struct {
	Mode  AddressingMode
	Value int
}

type AddressingMode int64

const (
	IMMEDIATE AddressingMode = iota
	DIRECT
	INDIRECT
)

type OpCode int64

const (
	DAT OpCode = iota
	MOV
	ADD
	SUB
	MUL
	DIV
	MOD
	JMP
	JMZ
	JMN
	DJN
	SPL
	CMP
	SEQ
	SNE
	SLT
	LDP
	STP
	NOP
)

func NewVM(size int, prog1, prog2 string) *VirtualMachine {
	mem := make([]Cell, 0)

	// ensure evenly divisible VM size
	if size%2 != 0 {
		size += 1
	}

	for i := 0; i < size; i++ {
		cell := Cell{
			Operation: DAT,
			AField: Field{
				Mode:  DIRECT,
				Value: 0,
			},
		}
		mem = append(mem, cell)
	}

	// add user programs into memory
	bytecode1, _ := compile(prog1)
	bytecode2, _ := compile(prog2)

	pos1 := rand.Intn(size - len(bytecode1))
	pos2 := rand.Intn(size - len(bytecode2))

	for inRange(pos1, pos1+len(bytecode1), pos2) {
		pos2 = rand.Intn(size)
	}

	mem = programVM(mem, prog1, pos1)
	mem = programVM(mem, prog2, pos2)

	return &VirtualMachine{
		Memory:    mem,
		APointers: NewProcessQueue(0, size),
		BPointers: NewProcessQueue(size/2, size),
		AProgram:  prog1,
		BProgram:  prog2,
		Winner:    "",
		Turn:      0,
	}
}

func programVM(mem []Cell, prog string, idx int) []Cell {
	bytecode, err := compile(prog)
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(bytecode); i++ {
		mem[idx+i] = bytecode[i]
	}
	return mem
}

func inRange(start, end, val int) bool {
	return val >= start && val <= end
}

func (vm *VirtualMachine) Tick() bool {
	// check for end of game conditions
	if vm.APointers.IsEmpty() && vm.BPointers.IsEmpty() || vm.Turn >= 1000 {
		vm.Winner = "Draw"
		return false
	}

	if vm.APointers.IsEmpty() {
		vm.Winner = "BPlayer"
		return false
	}
	if vm.BPointers.IsEmpty() {
		vm.Winner = "APlayer"
		return false
	}

	// move the game forward 1 turn
	vm.execute(vm.APointers.GetNextProcessPtr(), vm.APointers)
	vm.APointers.TickProcess()
	vm.execute(vm.BPointers.GetNextProcessPtr(), vm.BPointers)
	vm.BPointers.TickProcess()
	vm.Turn += 1

	return true
}

func (vm *VirtualMachine) Display() {
	for i := 0; i < len(vm.Memory); i++ {
		fmt.Printf("[%v %v%v, %v%v] ",
			vm.Memory[i].Operation,
			vm.Memory[i].AField.Mode,
			vm.Memory[i].AField.Value,
			vm.Memory[i].BField.Mode,
			vm.Memory[i].BField.Value,
		)

		if (i+1)%4 == 0 {
			fmt.Println()
		}
	}
}
