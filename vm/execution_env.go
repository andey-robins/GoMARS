package vm

func (vm *VirtualMachine) execute(idx int, processQueue *ProcessQueue) {
	currCell := vm.Memory[idx]
	switch currCell.Operation {
	case DAT:
		processQueue.KillProcess()
	case MOV:
		sourceIdx := vm.resolveFieldToIndex(idx, currCell.AField)
		targetIdx := vm.resolveFieldToIndex(idx, currCell.BField)
		vm.Memory[targetIdx] = vm.Memory[sourceIdx]
	case ADD:
		addedVal := vm.resolveFieldToIndex(idx, currCell.AField)
		targetIdx := vm.resolveFieldToIndex(idx, currCell.BField)
		vm.Memory[targetIdx].BField.Value += addedVal
	case SUB:
		subtractedVal := vm.resolveFieldToIndex(idx, currCell.AField)
		targetIdx := vm.resolveFieldToIndex(idx, currCell.BField)
		vm.Memory[targetIdx].BField.Value -= subtractedVal
	case MUL:
		multipliedVal := vm.resolveFieldToIndex(idx, currCell.AField)
		targetIdx := vm.resolveFieldToIndex(idx, currCell.BField)
		vm.Memory[targetIdx].BField.Value *= multipliedVal
	case DIV:
		divisor := vm.resolveFieldToIndex(idx, currCell.AField)
		targetIdx := vm.resolveFieldToIndex(idx, currCell.BField)
		vm.Memory[targetIdx].BField.Value /= divisor
	case MOD:
		modVal := vm.resolveFieldToIndex(idx, currCell.AField)
		targetIdx := vm.resolveFieldToIndex(idx, currCell.BField)
		vm.Memory[targetIdx].BField.Value %= modVal
	case JMP:
		offset := vm.resolveFieldToIndex(idx, currCell.AField)
		processQueue.Jump(offset)
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
		return field.Value % len(vm.Memory)
	case DIRECT:
		return (idx + field.Value) % len(vm.Memory)
	case INDIRECT:
		return (idx + vm.Memory[idx+field.Value].BField.Value) % len(vm.Memory)
	default:
		// this should never occur and also it will force a crash if we somehow reach it
		return -1
	}
}
