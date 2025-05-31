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
        mst_edges = []
        for w, u, v in self.edges:
            if uf.find(u) != uf.find(v):
                uf.union(u, v)
                mst_weight += w
                mst_edges.append((u, v, w))
        return mst_weight, mst_edges
    
class Prim:
    def __init__(self, v:int):
        self.v_cnt = v
        self.graph = [[] for _ in range(v)]

    def add_edge(self, u, v, w):
        self.graph[u].append((w, v))
        self.graph[v].append((w, u))

    def get_mst(self):
        import heapq
        mst_weight = 0
        mst_edges = []
        visited = [False] * self.v_cnt
        min_heap = []
        visited[0] = True
        for w, v in self.graph[0]:
            heapq.heappush(min_heap, (w, v))

        while min_heap:
            w, v = heapq.heappop(min_heap)
            if not visited[v]:
                visited[v] = True
                mst_weight += w
                mst_edges.append((0, v, w))
                for next_w, next_v in self.graph[v]:
                    if not visited[next_v]:
                        heapq.heappush(min_heap, (next_w, next_v))
        return mst_weight, mst_edges

def mst_algo(V, E):
    floor_log_V = V.bit_length() - 1
    if E > (V * floor_log_V): # dense
        return Prim(V)
    else: # sparse
        return Kruskal(V)