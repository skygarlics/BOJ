package algos

type UnionFind[T comparable] struct {
	parent map[T]T
	rank   map[T]int
	roots  map[T]struct{}
	size   map[T]int
}

func NewUnionFind[T comparable](elements []T) *UnionFind[T] {
	uf := &UnionFind[T]{
		parent: make(map[T]T),
		rank:   make(map[T]int),
		roots:  make(map[T]struct{}),
		size:   make(map[T]int),
	}
	for _, elem := range elements {
		uf.roots[elem] = struct{}{}
		uf.parent[elem] = elem
		uf.rank[elem] = 0
		uf.size[elem] = 1
	}
	return uf
}
func (uf *UnionFind[T]) Find(x T) T {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}
func (uf *UnionFind[T]) Union(x, y T) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)
	if rootX != rootY {
		if uf.rank[rootX] < uf.rank[rootY] {
			uf.parent[rootX] = rootY
			uf.size[rootY] += uf.size[rootX]
			delete(uf.roots, rootX)
		} else if uf.rank[rootX] > uf.rank[rootY] {
			uf.parent[rootY] = rootX
			uf.size[rootX] += uf.size[rootY]
			delete(uf.roots, rootY)
		} else {
			uf.parent[rootY] = rootX
			uf.size[rootX] += uf.size[rootY]
			delete(uf.roots, rootY)
			uf.rank[rootX]++
		}
	}
}
func (uf *UnionFind[T]) Roots() []T {
	roots := make([]T, 0, len(uf.roots))
	for root := range uf.roots {
		roots = append(roots, root)
	}
	return roots
}
func (uf *UnionFind[T]) GetRootCount() int {
	return len(uf.roots)
}
func (uf *UnionFind[T]) GetSize(x T) int {
	root := uf.Find(x)
	return uf.size[root]
}
func (uf *UnionFind[T]) Add(x T) {
	if _, exists := uf.parent[x]; !exists {
		uf.roots[x] = struct{}{}
		uf.parent[x] = x
		uf.rank[x] = 0
		uf.size[x] = 1
	}
}