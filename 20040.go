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

// ============= sized uf ============
type SizedUF struct {
	parent []int
	rank   []int
	size   []int
}

func NewSizedUF(cnt int) *SizedUF {
	uf := &SizedUF{
		parent: make([]int, cnt),
		rank:   make([]int, cnt),
		size:   make([]int, cnt),
	}
	for elem := 0; elem < cnt; elem++ {
		uf.parent[elem] = elem
		uf.rank[elem] = 0
		uf.size[elem] = 1
	}
	return uf
}

func (uf *SizedUF) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}
func (uf *SizedUF) Union(x, y int) error {
	rootX := uf.Find(x)
	rootY := uf.Find(y)
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
	return fmt.Errorf("elements %d and %d are already in the same set", x, y)
}
func (uf *SizedUF) Roots() []int {
	roots := make([]int, 0)
	for id := 0; id < len(uf.parent); id++ {
		parent := uf.Find(id)
		if parent == id {
			roots = append(roots, id)
		}
	}
	return roots
}
func (uf *SizedUF) GetRootCount() int {
	roots := uf.Roots()
	return len(roots)
}
func (uf *SizedUF) GetSize(x int) int {
	root := uf.Find(x)
	if root < 0 || root >= len(uf.size) {
		return 0
	}
	return uf.size[root]
}

// ============= generic uf ============
type UnionFind interface {
	Find(x int) int
	Union(x, y int) error
	Roots() []int
	GetRootCount() int
	GetSize(x int) int
}

type GenericUF[T comparable] struct {
	id2val map[int]T
	val2id map[T]int
	internal UnionFind
}
func NewGenericUF[T comparable](elements []T, impl UnionFind) *GenericUF[T] {
	uf := &GenericUF[T]{
		id2val: make(map[int]T, len(elements)),
		val2id: make(map[T]int, len(elements)),
		internal: impl,
	}
	for idx, elem := range elements {
		uf.val2id[elem] = idx
		uf.id2val[idx] = elem
	}
	return uf
}
func (uf *GenericUF[T]) Find(x T) (T, error) {
	id, exists := uf.val2id[x]
	if !exists {
		return x, fmt.Errorf("element not found: %v", x)
	}
	rootId := uf.internal.Find(id)
	return uf.id2val[rootId], nil
}
func (uf *GenericUF[T]) Union(x, y T) error {
	idX, existsX := uf.val2id[x]
	idY, existsY := uf.val2id[y]
	if !existsX || !existsY {
		return fmt.Errorf("one or both elements not found: %v, %v", x, y)
	}
	err := uf.internal.Union(idX, idY)
	if err != nil {
		return err
	}
	return nil
}
func (uf *GenericUF[T]) Roots() []T {
	roots := uf.internal.Roots()
	result := make([]T, len(roots))
	for i, root := range roots {
		result[i] = uf.id2val[root]
	}
	return result
}
func (uf *GenericUF[T]) GetRootCount() int {
	roots := uf.internal.Roots()
	return len(roots)
}
func (uf *GenericUF[T]) GetSize(x T) (int, error) {
	id, exists := uf.val2id[x]
	if !exists {
		return 0, fmt.Errorf("element not found: %v", x)
	}
	size := uf.internal.GetSize(id)
	return size, nil
}

func main() {
	defer writer.Flush()

	N, M := scanInt(), scanInt()
	arr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = i
	}
	impl := NewSizedUF(N)
	uf := NewGenericUF(arr, impl)

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