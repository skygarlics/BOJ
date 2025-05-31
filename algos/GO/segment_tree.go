package algos

type SegmentTreeOperation[T number] interface {
	Combine(a, b T) T
	Identity() T
}

type SumOperation[T number] struct{}

func (SumOperation[T]) Combine(a, b T) T {
	return a + b
}
func (SumOperation[T]) Identity() T {
	return 0
}

type MaxOperation[T number] struct{}

func (MaxOperation[T]) Combine(a, b T) T {
	return max(a, b)
}

type MinOperation[T number] struct{}

func (MinOperation[T]) Combine(a, b T) T {
	return min(a, b)
}

type SegmentTree[T number] struct {
	data    []T
	tree    []T
	lazy    []T
	size    int
	lazySet bool
	op      SegmentTreeOperation[T]
}

func (st *SegmentTree[T]) Build(arr []T) {
	st.size = len(arr)
	st.data = make([]T, st.size)
	copy(st.data, arr)
	st.tree = make([]T, 4*st.size)
	st.lazy = make([]T, 4*st.size)
	st.build(1, 0, st.size-1)
}
func (st *SegmentTree[T]) build(node, l, r int) {
	if l == r {
		st.tree[node] = st.data[l]
		return
	}
	mid := (l + r) / 2
	st.build(node*2, l, mid)
	st.build(node*2+1, mid+1, r)
	st.tree[node] = st.op.Combine(st.tree[node*2], st.tree[node*2+1])
}

func (st *SegmentTree[T]) Update(ql, qr int, val T) {
	st.update(1, 0, st.size-1, ql, qr, val)
}

func (st *SegmentTree[T]) push(node, l, r int) {
	if st.lazy[node] != st.op.Identity() {
		st.tree[node] = st.op.Combine(st.tree[node], st.lazy[node])
		if l != r {
			st.lazy[node*2] = st.op.Combine(st.lazy[node*2], st.lazy[node])
			st.lazy[node*2+1] = st.op.Combine(st.lazy[node*2+1], st.lazy[node])
		}
		st.lazy[node] = st.op.Identity()
	}
}

func (st *SegmentTree[T]) update(node, l, r, ql, qr int, val T) {
	st.push(node, l, r)
	if r < ql || l > qr {
		return
	}
	if ql <= l && r <= qr {
		st.lazy[node] = st.op.Combine(st.lazy[node], val)
		st.push(node, l, r)
		return
	}
	mid := (l + r) / 2
	st.update(node*2, l, mid, ql, qr, val)
	st.update(node*2+1, mid+1, r, ql, qr, val)
	st.tree[node] = st.op.Combine(st.tree[node*2], st.tree[node*2+1])
}

// Range query: get combined value in [ql, qr]
func (st *SegmentTree[T]) Query(ql, qr int) T {
	return st.query(1, 0, st.size-1, ql, qr)
}

func (st *SegmentTree[T]) query(node, l, r, ql, qr int) T {
	st.push(node, l, r)
	if r < ql || l > qr {
		return st.op.Identity()
	}
	if ql <= l && r <= qr {
		return st.tree[node]
	}
	mid := (l + r) / 2
	left := st.query(node*2, l, mid, ql, qr)
	right := st.query(node*2+1, mid+1, r, ql, qr)
	return st.op.Combine(left, right)
}