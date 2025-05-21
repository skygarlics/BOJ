package algos

const (
	RED   = true
	BLACK = false
)

type RBnode struct {
	val    int64
	cnt    int
	color  bool
	rt     *RBnode
	lt     *RBnode
	parent *RBnode
}
type RBTree struct {
	root *RBnode
	size int
}

func NewRBTree() *RBTree {
	return &RBTree{root: nil, size: 0}
}
func (t *RBTree) leftRotate(x *RBnode) *RBnode {
	y := x.rt

	x.rt = y.lt
	if y.lt != nil {
		y.lt.parent = x
	}

	y.parent = x.parent
	if x.parent == nil {
		t.root = y
	} else if x == x.parent.lt {
		x.parent.lt = y
	} else {
		x.parent.rt = y
	}
	y.lt = x
	x.parent = y
	return y
}
func (t *RBTree) rightRotate(x *RBnode) *RBnode {
	y := x.lt

	x.lt = y.rt
	if y.rt != nil {
		y.rt.parent = x
	}

	y.parent = x.parent
	if x.parent == nil {
		t.root = y
	} else if x == x.parent.rt {
		x.parent.rt = y
	} else {
		x.parent.lt = y
	}

	y.rt = x
	x.parent = y
	return y
}
func (t *RBTree) insert(val int64) {
	node := &RBnode{val: val, color: RED, cnt: 1}
	var y *RBnode = nil
	x := t.root

	for x != nil {
		y = x
		if node.val < x.val {
			x = x.lt
		} else if node.val > x.val {
			x = x.rt
		} else {
			x.cnt++
			return
		}
	}
	node.parent = y

	if y == nil {
		t.root = node
	} else if node.val < y.val {
		y.lt = node
	} else {
		y.rt = node
	}

	t.size++
	t.insertFixup(node)
}
func (t *RBTree) insertFixup(node *RBnode) {
	for node.parent != nil && node.parent.color == RED {
		if node.parent == node.parent.parent.lt {
			y := node.parent.parent.rt
			if y != nil && y.color == RED {
				node.parent.color = BLACK
				y.color = BLACK
				node.parent.parent.color = RED
				node = node.parent.parent
			} else {
				if node == node.parent.rt {
					node = node.parent
					t.leftRotate(node)
				}
				node.parent.color = BLACK
				node.parent.parent.color = RED
				t.rightRotate(node.parent.parent)
			}
		} else {
			y := node.parent.parent.lt
			if y != nil && y.color == RED {
				node.parent.color = BLACK
				y.color = BLACK
				node.parent.parent.color = RED
				node = node.parent.parent
			} else {
				if node == node.parent.lt {
					node = node.parent
					t.rightRotate(node)
				}
				node.parent.color = BLACK
				node.parent.parent.color = RED
				t.leftRotate(node.parent.parent)
			}
		}
	}
	t.root.color = BLACK
}
func (t *RBTree) min() *RBnode {
	node := t.root
	for node.lt != nil {
		node = node.lt
	}
	return node
}
func (t *RBTree) max() *RBnode {
	node := t.root
	for node.rt != nil {
		node = node.rt
	}
	return node
}
func (t *RBTree) deleteVal(val int64) {
	node := t.search(val)
	if node == nil {
		return
	}
	t.decNode(node)
}
func (t *RBTree) decNode(node *RBnode) {
	// Decrease the count of the node and delete it if the count is 0
	if node == nil {
		return
	}
	if node.cnt > 1 {
		node.cnt--
		return
	}
	t.deleteNode(node)
}

func (t *RBTree) deleteNode(node *RBnode) {
	if node == nil {
		return
	}
	t.size--
	var y *RBnode = node
	var yOriginalColor bool = y.color

	var x *RBnode
	if node.lt == nil {
		x = node.rt
		t.transplant(node, node.rt)
	} else if node.rt == nil {
		x = node.lt
		t.transplant(node, node.lt)
	} else {
		y = t.min()
		yOriginalColor = y.color
		x = y.rt
		if y.parent == node {
			x.parent = y
		} else {
			t.transplant(y, y.rt)
			y.rt = node.rt
			y.rt.parent = y
		}
		t.transplant(node, y)
		y.lt = node.lt
		y.lt.parent = y
		y.color = node.color
	}
	if yOriginalColor == BLACK {
		t.deleteFixup(x)
	}
}
func (t *RBTree) transplant(u, v *RBnode) {
	if u.parent == nil {
		t.root = v
	} else if u == u.parent.lt {
		u.parent.lt = v
	} else {
		u.parent.rt = v
	}
	if v != nil {
		v.parent = u.parent
	}
}
func (t *RBTree) deleteFixup(x *RBnode) {
	var parent *RBnode
	if x != nil {
		parent = x.parent
	}
	for x != t.root && (x == nil || x.color == BLACK) {
		if parent != nil && x == parent.lt {
			w := parent.rt
			if w != nil && w.color == RED {
				w.color = BLACK
				parent.color = RED
				t.leftRotate(parent)
				w = parent.rt
			}
			if (w == nil || (w.lt == nil || w.lt.color == BLACK)) && (w == nil || (w.rt == nil || w.rt.color == BLACK)) {
				if w != nil {
					w.color = RED
				}
				x = parent
				if x != nil {
					parent = x.parent
				}
			} else {
				if w != nil && (w.rt == nil || w.rt.color == BLACK) {
					if w.lt != nil {
						w.lt.color = BLACK
					}
					w.color = RED
					t.rightRotate(w)
					w = parent.rt
				}
				if w != nil {
					w.color = parent.color
				}
				parent.color = BLACK
				if w != nil && w.rt != nil {
					w.rt.color = BLACK
				}
				t.leftRotate(parent)
				x = t.root
				break
			}
		} else {
			w := parent.lt
			if w != nil && w.color == RED {
				w.color = BLACK
				parent.color = RED
				t.rightRotate(parent)
				w = parent.lt
			}
			if (w == nil || (w.lt == nil || w.lt.color == BLACK)) && (w == nil || (w.rt == nil || w.rt.color == BLACK)) {
				if w != nil {
					w.color = RED
				}
				x = parent
				if x != nil {
					parent = x.parent
				}
			} else {
				if w != nil && (w.lt == nil || w.lt.color == BLACK) {
					if w.rt != nil {
						w.rt.color = BLACK
					}
					w.color = RED
					t.leftRotate(w)
					w = parent.lt
				}
				if w != nil {
					w.color = parent.color
				}
				parent.color = BLACK
				if w != nil && w.lt != nil {
					w.lt.color = BLACK
				}
				t.rightRotate(parent)
				x = t.root
				break
			}
		}
	}
	if x != nil {
		x.color = BLACK
	}
}
func (t *RBTree) search(val int64) *RBnode {
	node := t.root
	for node != nil {
		if val < node.val {
			node = node.lt
		} else if val > node.val {
			node = node.rt
		} else {
			return node
		}
	}
	return nil
}
func (t *RBTree) len() int {
	return t.size
}
