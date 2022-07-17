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
		if divisor != 0 {
			vm.Memory[targetIdx].BField.Value /= divisor
		}
	case MOD:
		modVal := vm.resolveFieldToIndex(idx, currCell.AField)
		targetIdx := vm.resolveFieldToIndex(idx, currCell.BField)
		vm.Memory[targetIdx].BField.Value %= modVal
	case JMP:
		offset := vm.resolveFieldToIndex(idx, currCell.AField)
		processQueue.Jump(offset)
	case JMZ:
		value := vm.resolveFieldToIndex(idx, currCell.BField)
		offset := vm.resolveFieldToIndex(idx, currCell.AField)
		if value == 0 {
			processQueue.Jump(offset)
		}
	case JMN:
		value := vm.resolveFieldToIndex(idx, currCell.BField)
		offset := vm.resolveFieldToIndex(idx, currCell.AField)
		if value != 0 {
			processQueue.Jump(offset)
		}
	case DJN:
		value := vm.resolveFieldToIndex(idx, currCell.BField) - 1
		vm.Memory[idx].AField.Value -= 1 // decrement the value actually in the cell
		offset := vm.resolveFieldToIndex(idx, currCell.AField)
		if value != 0 {
			processQueue.Jump(offset)
		}
	case SPL:
		targetIdx := vm.resolveFieldToIndex(idx, currCell.AField)
		processQueue.SplitProcessTo(targetIdx)
	case CMP: // equivalent to SEQ
		idxOne := vm.resolveFieldToIndex(idx, currCell.AField)
		idxTwo := vm.resolveFieldToIndex(idx, currCell.BField)
		if vm.Memory[idxOne].Operation == vm.Memory[idxTwo].Operation {
			processQueue.TickProcess()
		}
	case SEQ: // equivalent to CMP
		idxOne := vm.resolveFieldToIndex(idx, currCell.AField)
		idxTwo := vm.resolveFieldToIndex(idx, currCell.BField)
		if vm.Memory[idxOne].Operation == vm.Memory[idxTwo].Operation {
			processQueue.TickProcess()
		}
	case SNE:
		idxOne := vm.resolveFieldToIndex(idx, currCell.AField)
		idxTwo := vm.resolveFieldToIndex(idx, currCell.BField)
		if vm.Memory[idxOne].Operation != vm.Memory[idxTwo].Operation {
			processQueue.TickProcess()
		}
	case SLT:
		idxOne := vm.resolveFieldToIndex(idx, currCell.AField)
		idxTwo := vm.resolveFieldToIndex(idx, currCell.BField)
		if vm.Memory[idxOne].Operation <= vm.Memory[idxTwo].Operation {
			processQueue.TickProcess()
		}
	case LDP:
		processQueue.KillProcess()
	case STP:
		processQueue.KillProcess()
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
