class SegmentTree:
    """ Segment tree for range sum and lazy propagation """
    def __init__(self, n):
        self.n = n
        self.tree = [0] * (n * 4)
        self.lazy = [0] * (n * 4)

    def build(self, arr, node=1, node_start=0, node_end=None):
        if node_end is None:
            node_end = self.n - 1

        if node_start == node_end:
            self.tree[node] = arr[node_start]
            return
        
        mid = (node_start + node_end) // 2
        self.build(arr, node * 2, node_start, mid)
        self.build(arr, node * 2 + 1, mid + 1, node_end)
        self.tree[node] = self.tree[node * 2] + self.tree[node * 2 + 1]

    def push(self, node, node_start, node_end):
        if self.lazy[node] != 0:
            self.tree[node] += self.lazy[node] * (node_end - node_start + 1)
            if node_start != node_end:
                self.lazy[node * 2] += self.lazy[node]
                self.lazy[node * 2 + 1] += self.lazy[node]
            self.lazy[node] = 0

    def update(self, l, r, value, node=1, node_start=0, node_end=None):
        if node_end is None:
            node_end = self.n - 1

        self.push(node, node_start, node_end)

        if r < node_start or l > node_end:
            return
        
        if l <= node_start and node_end <= r:
            self.lazy[node] += value
            self.push(node, node_start, node_end)
            return
        
        mid = (node_start + node_end) // 2
        self.update(l, r, value, node * 2, node_start, mid)
        self.update(l, r, value, node * 2 + 1, mid + 1, node_end)
        self.tree[node] = self.tree[node * 2] + self.tree[node * 2 + 1]

    def query(self, l, r, node=1, node_start=0, node_end=None):
        if node_end is None:
            node_end = self.n - 1

        self.push(node, node_start, node_end)

        if r < node_start or l > node_end:
            return 0
        
        if l <= node_start and node_end <= r:
            return self.tree[node]
        
        mid = (node_start + node_end) // 2
        left_sum = self.query(l, r, node * 2, node_start, mid)
        right_sum = self.query(l, r, node * 2 + 1, mid + 1, node_end)
        return left_sum + right_sum