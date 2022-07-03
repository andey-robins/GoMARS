package vm

type ProcessQueue struct {
	pointers      []int
	activePointer int
	memoryWidth   int
}

func NewProcessQueue(ptr, w int) *ProcessQueue {
	ptrs := make([]int, 1)
	ptrs[0] = ptr
	return &ProcessQueue{
		pointers:      ptrs,
		activePointer: 0,
		memoryWidth:   w,
	}
}

func (pq *ProcessQueue) GetNextProcessPtr() int {
	ptr := pq.pointers[pq.activePointer]
	pq.activePointer = (pq.activePointer + 1) % len(pq.pointers)
	return ptr
}

func (pq *ProcessQueue) SplitProcessTo(newPtr int) {
	pq.pointers = append(pq.pointers, newPtr)
}

func (pq *ProcessQueue) KillProcess() {
	if pq.activePointer == len(pq.pointers)-1 {
		pq.pointers = pq.pointers[:pq.activePointer]
	} else {
		pq.pointers = append(pq.pointers[:pq.activePointer], pq.pointers[pq.activePointer+1:]...)
	}

	pq.activePointer--

	if pq.activePointer < 0 {
		pq.activePointer += len(pq.pointers)
	}

	// delay hack to let us ignore an empty pointer list until it gets properly caught by IsEmpty
	if len(pq.pointers) == 0 {
		pq.activePointer = 0
	}
}

func (pq *ProcessQueue) TickProcess() {
	if len(pq.pointers) != 0 && pq.activePointer < len(pq.pointers) {
		pq.pointers[pq.activePointer] = (pq.pointers[pq.activePointer] + 1) % pq.memoryWidth
	}
}

func (pq *ProcessQueue) IsEmpty() bool {
	return len(pq.pointers) == 0
}

func (pq *ProcessQueue) Jump(newPtr int) {
	pq.pointers[pq.activePointer] = newPtr
}
