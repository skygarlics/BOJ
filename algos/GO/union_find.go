package algos

import "fmt"

// ============= lightweight ============
type UF struct {
	parent []int
	rank   []int
}
func NewIntUF(cnt int) *UF {
	uf := &UF{
		parent: make([]int, cnt),
		rank:   make([]int, cnt),
	}
	for elem := 0; elem < cnt; elem++ {
		uf.parent[elem] = elem
		uf.rank[elem] = 0
	}
	return uf
}
func (uf *UF) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}
func (uf *UF) Union(x, y int) error {
	rootX := uf.Find(x)
	rootY := uf.Find(y)
	if rootX != rootY {
		if uf.rank[rootX] < uf.rank[rootY] {
			uf.parent[rootX] = rootY
		} else if uf.rank[rootX] > uf.rank[rootY] {
			uf.parent[rootY] = rootX
		} else {
			uf.parent[rootY] = rootX
			uf.rank[rootX]++
		}
		return nil
	}
	return fmt.Errorf("elements %d and %d are already in the same set", x, y)
}
func (uf *UF) Roots() []int {
	roots := make([]int, 0)
	for id := 0; id < len(uf.parent); id++ {
		parent := uf.Find(id)
		if parent == id {
			roots = append(roots, id)
		}
	}
	return roots
}
func (uf *UF) GetRootCount() int {
	roots := uf.Roots()
	return len(roots)
}
func (uf *UF) GetSize(x int) int {
	root := uf.Find(x)
	size := 0
	for i := range uf.parent {
		if uf.Find(i) == root {
			size++
		}
	}
	return size
}
func (uf *UF) AddElement() int {
	newId := len(uf.parent)
	uf.parent = append(uf.parent, newId)
	uf.rank = append(uf.rank, 0)
	return newId
}

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
func (uf *SizedUF) AddElement() int {
	newId := len(uf.parent)
	uf.parent = append(uf.parent, newId)
	uf.rank = append(uf.rank, 0)
	uf.size = append(uf.size, 1)
	return newId
}

// ============= generic uf ============
type UnionFind interface {
	Find(x int) int
	Union(x, y int) error
	Roots() []int
	GetRootCount() int
	GetSize(x int) int
	AddElement() int
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
func (uf *GenericUF[T]) AddElement(elem T) int {
	newId := uf.internal.AddElement()
	uf.id2val[newId] = elem
	uf.val2id[elem] = newId
	return newId
}