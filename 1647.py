import sys; inputs = sys.stdin.read().split()
it = iter(map(int, inputs))

class UnionFind:
    def __init__(self, n):
        self.parent = list(range(n))
        self.rank = [0] * n

    def find(self, u):
        if self.parent[u] != u:
            self.parent[u] = self.find(self.parent[u])
        return self.parent[u]

    def union(self, u, v):
        root_u = self.find(u)
        root_v = self.find(v)

        if root_u != root_v:
            if self.rank[root_u] > self.rank[root_v]:
                self.parent[root_v] = root_u
            elif self.rank[root_u] < self.rank[root_v]:
                self.parent[root_u] = root_v
            else:
                self.parent[root_v] = root_u
                self.rank[root_u] += 1
    
class Kruskal:
    def __init__(self, v:int):
        self.v_cnt = v
        self.edges = []

    def add_edge(self, u, v, w):
        self.edges.append((w, u, v))

    def get_mst(self):
        uf = UnionFind(self.v_cnt)
        self.edges.sort()
        mst_weight = 0
        max_weight = 0
        for w, u, v in self.edges:
            if uf.find(u) != uf.find(v):
                uf.union(u, v)
                mst_weight += w
                max_weight = max(max_weight, w)
        return mst_weight, max_weight
    

N, M = next(it), next(it)
mst = Kruskal(N)
for _ in range(M):
    u, v, w = next(it), next(it), next(it)
    mst.add_edge(u - 1, v - 1, w)

mst_weight, max_weight = mst.get_mst()
print(mst_weight - max_weight)