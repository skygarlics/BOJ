package algos

import "container/heap"

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
func (h *Heap[T]) Peek() T { return h.data[0] }

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