package vm

func (vm *VirtualMachine) execute(idx int) {
	currCell := vm.Memory[idx]
	switch currCell.Operation {
	case DAT:
		return
	case MOV:
		sourceIdx := vm.resolveFieldToIndex(idx, currCell.AField)
		targetIdx := vm.resolveFieldToIndex(idx, currCell.BField)
		vm.Memory[targetIdx] = vm.Memory[sourceIdx]
	case ADD:
		return
	case SUB:
		return
	case MUL:
		return
	case DIV:
		return
	case MOD:
		return
	case JMP:
		return
	case JMZ:
		return
	case JMN:
		return
	case DJN:
		return
	case SPL:
		return
	case CMP:
		return
	case SEQ:
		return
	case SNE:
		return
	case SLT:
		return
	case LDP:
		return
	case STP:
		return
	case NOP:
		return
	}
}

func (vm *VirtualMachine) resolveFieldToIndex(idx int, field Field) int {
	switch field.Mode {
	case IMMEDIATE:
		return field.Value
	case DIRECT:
		return idx + field.Value
	case INDIRECT:
		return idx + vm.Memory[idx+field.Value].BField.Value
	default:
		// this should never occur and also it will force a crash
		return -1
	}
}
