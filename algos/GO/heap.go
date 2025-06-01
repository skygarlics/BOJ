package algos

import (
	"container/heap"
)


type Heap[T any] struct {
	data []T
	less func(a, b T) bool
}
func (h *Heap[T]) Len() int           { return len(h.data) }
func (h *Heap[T]) Less(i, j int) bool { return h.less(h.data[i], h.data[j]) }
func (h *Heap[T]) Swap(i, j int)      { h.data[i], h.data[j] = h.data[j], h.data[i] }
func (h *Heap[T]) Push(x any) {
	h.data = append(h.data, x.(T))
}
func (h *Heap[T]) PushVal(v T) { heap.Push(h, v) }
func (h *Heap[T]) Pop() any {
	last := len(h.data) - 1
	v := h.data[last]
	h.data = h.data[:last]
	return v
}
func (h *Heap[T]) PopVal() T { return heap.Pop(h).(T) }
func (h *Heap[T]) Peek() T { 
	return h.data[0]
}
func (h *Heap[T]) IsEmpty() bool { return len(h.data) == 0 }
func (h *Heap[T]) Clear() { h.data = h.data[:0] }
func NewMinHeap[T number]() *Heap[T] {
	h := &Heap[T]{less: func(a, b T) bool { return a < b }}
	heap.Init(h)
	return h
}

func NewMaxHeap[T number]() *Heap[T] {
	h := &Heap[T]{less: func(a, b T) bool { return a > b }}
	heap.Init(h)
	return h
}


type heapElement[T number] struct {
	value T
	id	  uint
}
type MinMaxHeap[T number] struct {
	minHeap *Heap[heapElement[T]]
	maxHeap *Heap[heapElement[T]]
	deleted []bool
	cnt     uint
	id      uint
}
func NewMinMaxHeap[T number](max_cnt int) *MinMaxHeap[T] {
	minh := &Heap[heapElement[T]]{
		less: func(a, b heapElement[T]) bool {
			if a.value == b.value {
				return a.id < b.id
			}
			return a.value < b.value
		},
		data: make([]heapElement[T], 0),
	}
	maxh := &Heap[heapElement[T]]{
		less: func(a, b heapElement[T]) bool {
			if a.value == b.value {
				return a.id < b.id
			}
			return a.value > b.value
		},
		data: make([]heapElement[T], 0),
	}
	h := &MinMaxHeap[T]{
		minHeap: minh,
		maxHeap: maxh,
		deleted: make([]bool, max_cnt+1),
		id:      0,
		cnt:     0,
	}
	return h
}
func (h *MinMaxHeap[T]) Push(v T) {
	elem := heapElement[T]{value: v, id: h.id}
	h.minHeap.PushVal(elem)
	h.maxHeap.PushVal(elem)
	h.deleted[h.id] = false
	h.id++
	h.cnt++
}
func (h *MinMaxHeap[T]) PopMin() T {
	for {
		if h.minHeap.Len() == 0 {
			return 0
		}
		elem := h.minHeap.PopVal()
		if !h.deleted[elem.id] {
			h.deleted[elem.id] = true
			h.cnt--
			if h.cnt == 0 {
				// reset heap
				h.maxHeap.Clear()
				h.id = 0
			}
			return elem.value
		}
	}
}
func (h *MinMaxHeap[T]) PeekMin() T {
	for {
		elem := h.minHeap.Peek()
		if !h.deleted[elem.id] {
			return elem.value
		}
		h.minHeap.PopVal() // remove deleted element
	}
}
func (h *MinMaxHeap[T]) PopMax() T {
	for {
		if h.maxHeap.Len() == 0 {
			return 0
		}
		elem := h.maxHeap.PopVal()
		if !h.deleted[elem.id] {
			h.deleted[elem.id] = true
			h.cnt--
			if h.cnt == 0 {
				h.minHeap.Clear()
				h.id = 0
			}
			return elem.value
		}
	}
}
func (h *MinMaxHeap[T]) PeekMax() T {
	for {
		elem := h.maxHeap.Peek()
		if !h.deleted[elem.id] {
			return elem.value
		}
		h.maxHeap.PopVal()
	}
}
func (h *MinMaxHeap[T]) IsEmpty() bool {
	return h.cnt == 0
}