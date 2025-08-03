from collections import deque
class Hopcroft:
    """ Bipartite Graph Class """
    def __init__(self, lt_cnt, rt_cnt):
        self.lt_cnt = lt_cnt
        self.rt_cnt = rt_cnt
        self.lt_edges = [[] for _ in range(lt_cnt)]

    def add_edge(self, u, v):
        """ Add edge from u in left to v in right """
        self.lt_edges[u].append(v)

    def hopcroft_karp(self):
        """ Hopcroft-Karp Algorithm """
        pair_u = [-1] * self.lt_cnt
        pair_v = [-1] * self.rt_cnt
        dist = [0] * self.lt_cnt

        def bfs():
            queue = deque()
            # Add unmatched left v into queue
            for u in range(self.lt_cnt):
                if pair_u[u] == -1:
                    dist[u] = 0
                    queue.append(u)
                else:
                    dist[u] = float('inf')
            # Set distance by BFS
            found = False
            while queue:
                u = queue.popleft()
                for v in self.lt_edges[u]:
                    if pair_v[v] == -1:
                        found = True
                    elif dist[pair_v[v]] == float('inf'):
                        dist[pair_v[v]] = dist[u] + 1
                        queue.append(pair_v[v])
            return found

        def dfs(u):
            for v in self.lt_edges[u]:
                # If v is not matched or finds an augmenting path
                if pair_v[v] == -1 or (dist[pair_v[v]] == dist[u] + 1 and dfs(pair_v[v])):
                    # In case of augment path, previous pair is updated by dfs(pair_v[v])
                    pair_u[u] = v
                    pair_v[v] = u
                    return True
            dist[u] = float('inf')
            return False

        matching_size = 0
        while bfs():
            for u in range(self.lt_cnt):
                if pair_u[u] == -1 and dfs(u):
                    matching_size += 1
            if matching_size == self.lt_cnt:
                return matching_size

        return matching_size