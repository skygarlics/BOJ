package algos

import (
	"math/bits"
	"sort"
)

type Edge[V number, W number] struct {
	From   V
	To     V
	Weight W
}

type MST[T number, W number] interface {
	// AddEdge adds an edge to the MST algorithm.
	AddEdge(from, to T, weight W)

	// GetMST returns the edges in the minimum spanning tree.
	GetMST() []Edge[T, W]
}

type Kruskal[V number, W number] struct {
	edges []Edge[V, W]
	uf    *UnionFind[V]
}
func (k *Kruskal[V, W]) AddEdge(from, to V, weight W) {
	k.edges = append(k.edges, Edge[V, W]{From: from, To: to, Weight: weight})
}
func (k *Kruskal[V, W]) GetMST() []Edge[V, W] {
	if len(k.edges) == 0 {
		return nil	}

	// Sort edges by weight
	sort.Slice(k.edges, func(i, j int) bool {
		return k.edges[i].Weight < k.edges[j].Weight
	})
	
	mst := []Edge[V, W]{}
	for _, edge := range k.edges {
		if k.uf.Find(edge.From) != k.uf.Find(edge.To) {
			k.uf.Union(edge.From, edge.To)
			mst = append(mst, edge)
		}
	}
	return mst
}
func NewKruskal[V number, W number](vertices []V) *Kruskal[V, W] {
	uf := NewUnionFind[V](vertices)
	return &Kruskal[V, W]{
		edges: []Edge[V, W]{},
		uf:    uf,
	}
}

type Prim[V number, W number] struct {
	adjacency map[V]map[V]W // adjacency list with weights
	visited   map[V]bool
	minHeap   *Heap[Edge[V, W]]
}
func (p *Prim[V, W]) AddEdge(from, to V, weight W) {
	if p.adjacency == nil {
		p.adjacency = make(map[V]map[V]W)
	}
	if p.adjacency[from] == nil {
		p.adjacency[from] = make(map[V]W)
	}
	if p.adjacency[to] == nil {
		p.adjacency[to] = make(map[V]W)
	}
	p.adjacency[from][to] = weight
	p.adjacency[to][from] = weight
}
func (p *Prim[V, W]) GetMST() []Edge[V, W] {
	if len(p.adjacency) == 0 {
		return nil
	}

	p.visited = make(map[V]bool)

	var startVertex V
	for v := range p.adjacency {
		startVertex = v
		break
	}

	p.visited[startVertex] = true
	for to, weight := range p.adjacency[startVertex] {
		p.minHeap.PushVal(Edge[V, W]{From: startVertex, To: to, Weight: weight})
	}

	mst := []Edge[V, W]{}
	for p.minHeap.Len() > 0 {
		edge := p.minHeap.PopVal()
		if p.visited[edge.To] {
			continue
		}
		mst = append(mst, edge)
		p.visited[edge.To] = true

		for to, weight := range p.adjacency[edge.To] {
			if !p.visited[to] {
				p.minHeap.PushVal(Edge[V, W]{From: edge.To, To: to, Weight: weight})
			}
		}
	}
	return mst
}
func NewPrim[V number, W number]() *Prim[V, W] {
	mh := &Heap[Edge[V, W]]{
		less: func(a, b Edge[V, W]) bool {
			return a.Weight < b.Weight
		},
	}

	return &Prim[V, W]{
		adjacency: make(map[V]map[V]W),
		visited:   make(map[V]bool),
		minHeap:   mh,
	}
}

// NewMST creates a new MST instance based on v, e counts
func NewMST[V number, W number](v_cnt, e_cnt uint) MST[V, W] {
	v_log_floor := bits.Len(v_cnt) - 1
	if e_cnt > v_cnt*uint(v_log_floor) {
		// dense graph
		return NewPrim[V, W]()
	} else {
		// sparse graph
		vertices := make([]V, 0, v_cnt)
		for i := uint(0); i < v_cnt; i++ {
			var vertex V
			vertices = append(vertices, vertex)
		}
		return NewKruskal[V, W](vertices)
	}
}