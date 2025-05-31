package main

import (
	"bufio"
	"fmt"
	"os"
)

// ============ io =============

var (
	reader = bufio.NewReaderSize(os.Stdin, 1<<20)
	writer = bufio.NewWriter(os.Stdout)
)

func scanString() string {
	var b byte
	for {
		b, _ = reader.ReadByte()
		if b > ' ' {
			break
		}
	}
	buf := make([]byte, 0, 20)
	for {
		buf = append(buf, b)
		b, _ = reader.ReadByte()
		if b <= ' ' {
			break
		}
	}
	return string(buf)
}
func scanInt() int {
	var n int
	sign := 1
	b, _ := reader.ReadByte()
	for (b < '0' || b > '9') && b != '-' {
		b, _ = reader.ReadByte()
	}
	if b == '-' {
		sign = -1
		b, _ = reader.ReadByte()
	}
	for '0' <= b && b <= '9' {
		n = n*10 + int(b-'0')
		b, _ = reader.ReadByte()
	}
	return n * sign
}

func print(a ...any) { fmt.Fprint(writer, a...) }
func println(a ...any) { fmt.Fprintln(writer, a...) }

// ============ Union-Find ============
type GenericUnionFind[T comparable] struct {
	val2id map[T] int
	id2val []T
	parent []int
	rank   []int
	size   []int
}

func NewGenericUF[T comparable](elements []T) *GenericUnionFind[T] {
	uf := &GenericUnionFind[T]{
		val2id: make(map[T]int, len(elements)),
		id2val: make([]T, len(elements)),
		parent: make([]int, len(elements)),
		rank:   make([]int, len(elements)),
		size:   make([]int, len(elements)),
	}
	for idx, elem := range elements {
		uf.val2id[elem] = idx
		uf.id2val[idx] = elem
		uf.parent[idx] = idx
		uf.rank[idx] = 0
		uf.size[idx] = 1
	}
	return uf
}
func (uf *GenericUnionFind[T]) FindRoot(id int) int {
	if uf.parent[id] != id {
		uf.parent[id] = uf.FindRoot(uf.parent[id])
	}
	return uf.parent[id]
}

func (uf *GenericUnionFind[T]) Find(x T) (T, error) {
	id, exists := uf.val2id[x]
	if !exists {
		return x, fmt.Errorf("element not found: %v", x)
	}
	rootId := uf.FindRoot(id)
	return uf.id2val[rootId], nil
}

func (uf *GenericUnionFind[T]) Union(x, y T) error {
	idX, existsX := uf.val2id[x]
	idY, existsY := uf.val2id[y]
	if !existsX || !existsY {
		return fmt.Errorf("one or both elements not found: %v, %v", x, y)
	}

	rootX := uf.FindRoot(idX)
	rootY := uf.FindRoot(idY)

	if rootX != rootY {
		if uf.rank[rootX] < uf.rank[rootY] {
			uf.parent[rootX] = rootY
			uf.size[rootY] += uf.size[rootX]
		} else if uf.rank[rootX] > uf.rank[rootY] {
			uf.parent[rootY] = rootX
			uf.size[rootX] += uf.size[rootY]
		} else {
			uf.parent[rootY] = rootX
			uf.rank[rootX]++
			uf.size[rootX] += uf.size[rootY]
		}
		return nil
	}
	return fmt.Errorf("elements are already in the same set: %v, %v", x, y)
}

func (uf *GenericUnionFind[T]) Roots() []T {
	roots := make([]T, 0)
	for id := 0; id < len(uf.parent); id++ {
		parent := uf.FindRoot(id)
		if parent == id {
			roots = append(roots, uf.id2val[id])
		}
	}
	return roots
}

func (uf *GenericUnionFind[T]) GetRootCount() int {
	roots := uf.Roots()
	return len(roots)
}

func (uf *GenericUnionFind[T]) GetSize(x T) (int, error) {
	id, exists := uf.val2id[x]
	if !exists {
		return 0, fmt.Errorf("element not found: %v", x)
	}
	rootId := uf.FindRoot(id)
	return uf.size[rootId], nil
}
func (uf *GenericUnionFind[T]) Add(x T) {
	if _, exists := uf.val2id[x]; exists {
		return // Element already exists
	}
	id := len(uf.id2val)
	uf.val2id[x] = id
	uf.id2val = append(uf.id2val, x)
	uf.parent = append(uf.parent, id)
	uf.rank = append(uf.rank, 0)
	uf.size = append(uf.size, 1)
}

func main() {
	defer writer.Flush()

	N, M := scanInt(), scanInt()
	arr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = i
	}
	uf := NewGenericUF(arr)

	for cnt := 1; cnt <= M; cnt++ {
		a, b := scanInt(), scanInt()
		err := uf.Union(a, b)
		if err != nil {
			println(cnt)
			return
		}
	}
	println(0)
}