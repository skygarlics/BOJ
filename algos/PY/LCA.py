class LCA:
    def __init__(self, n, edges):
        self.n = n
        self.max_depth_log = (n - 1).bit_length()
        
        self.adj = [[] for _ in range(n + 1)]
        self.depth = [-1] * (n + 1)
        self.partsum = [0] * (n + 1)
        
        self.up = [[-1] * (n + 1) for _ in range(self.max_depth_log)] 

        for u, v, dist in edges:
            self.adj[u].append((v, dist))
            self.adj[v].append((u, dist))

        self.dfs(1, 0, 0)

        for i in range(1, self.max_depth_log):
            for u in range(1, n + 1):
                if self.up[i - 1][u] != -1:
                    self.up[i][u] = self.up[i - 1][self.up[i - 1][u]]

    def dfs(self, u, p, current_dist):
        self.depth[u] = self.depth[p] + 1 if p != 0 else 0
        self.partsum[u] = current_dist #
        self.up[0][u] = p

        for neigh, dist in self.adj[u]:
            if neigh != p:
                self.dfs(neigh, u, current_dist + dist)

    def lca(self, u, v):

        if self.depth[u] < self.depth[v]:
            u, v = v, u

        diff = self.depth[u] - self.depth[v]
        for i in range(self.max_depth_log):
            if (diff >> i) & 1:
                u = self.up[i][u]

        if u == v:
            return u

        for i in range(self.max_depth_log - 1, -1, -1):
            if self.up[i][u] != -1 and self.up[i][u] != self.up[i][v]:
                u = self.up[i][u]
                v = self.up[i][v]
        return self.up[0][u]

    def distance(self, u, v):
        ancestor = self.lca(u, v)
        return self.partsum[u] + self.partsum[v] - 2 * self.partsum[ancestor]