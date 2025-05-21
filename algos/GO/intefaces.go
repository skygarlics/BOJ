package algos

type signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type integer interface {
	signed | unsigned
}

type float interface {
	~float32 | ~float64
}

type number interface {
	integer | float
}
type Heap[T any] struct {
	data []T
	less func(a, b T) bool
}