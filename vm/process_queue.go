package vm

type ProcessQueue struct {
	pointers      []int
	activePointer int
}

func NewProcessQueue(ptr int) *ProcessQueue {
	ptrs := make([]int, 1)
	ptrs[0] = ptr
	return &ProcessQueue{
		pointers:      ptrs,
		activePointer: 0,
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

func (pq *ProcessQueue) TickProcess() {
	pq.pointers[pq.activePointer] += 1
}

func (pq *ProcessQueue) IsEmpty() bool {
	return len(pq.pointers) == 0
}
