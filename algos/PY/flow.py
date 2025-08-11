from collections import deque
class Dinic:
    """ Dinic's Algorithm for Maximum Flow """

    def __init__(self, N):
        self.N = N
        self.graph = [[] for _ in range(N)]
        self.cap = dict()

    def add_edge(self, u, v, cap):
        """ Add edge with capacity, add reverse edge """
        self.graph[u].append(v)
        self.graph[v].append(u)
        self.cap[(u, v)] = self.cap.get((u, v), 0) + cap
        self.cap.setdefault((v, u), 0)  # ensure reverse edge exists

    def bfs(self, s, t, level):
        queue = deque([s])
        level[s] = 0
        while queue:
            v = queue.popleft()
            for u in self.graph[v]:
                if self.cap[(v, u)] > 0 and level[u] < 0:
                    level[u] = level[v] + 1
                    queue.append(u)
        return level[t] >= 0

    def dfs(self, v, t, upTo, level, iter):
        if v == t:
            return upTo
        for i in range(iter[v], len(self.graph[v])):
            u = self.graph[v][i]
            if self.cap[(v, u)] > 0 and level[v] < level[u]:
                d = self.dfs(u, t, min(upTo, self.cap[(v, u)]), level, iter)
                if d > 0:
                    self.cap[(v, u)] -= d
                    self.cap[(u, v)] += d
                    return d
            iter[v] += 1
        return 0

    def max_flow(self, s, t):
        flow = 0
        INF = float('inf')
        level = [-1] * self.N
        while self.bfs(s, t, level):
            iter = [0] * self.N
            f = self.dfs(s, t, INF, level, iter)
            while f > 0:
                flow += f
                f = self.dfs(s, t, INF, level, iter)
            level = [-1] * self.N  # reset level graph for next round
        return flow
    

from collections import deque
import heapq
class MCMF_BF:
    def __init__(self, N, cutoff, pathfunc=None):
        self.N = N
        self.graph = [[] for _ in range(N)]
        self.cutoff = cutoff
        self.cost = {}
        self.cap = {}
        if pathfunc == 'dijkstra':
            self.h = [0] * self.N
            self.f = self.dijkstra
        else:
            self.f = self.spfa

    def add_edge(self, u, v, capacity, cost):
        self.graph[u].append(v)
        self.graph[v].append(u)
        self.cap[(u, v)] = self.cap.get((u, v), 0) + capacity
        self.cap.setdefault((v, u), 0)
        self.cost[(u, v)] = cost
        self.cost[(v, u)] = -cost
    
    def spfa(self, s, t):
        dist = [float('inf')] * self.N
        parent = [-1] * self.N
        in_queue = [False] * self.N
        
        dist[s] = 0
        queue = deque([s])
        in_queue[s] = True
        
        while queue:
            u = queue.popleft()
            in_queue[u] = False
            
            for v in self.graph[u]:
                if self.cap[(u, v)] > 0 and dist[u] + self.cost[(u, v)] < dist[v]:
                    dist[v] = dist[u] + self.cost[(u, v)]
                    parent[v] = u
                    if not in_queue[v]:
                        queue.append(v)
                        in_queue[v] = True
        
        return dist[t] != float('inf'), parent
    
    def dijkstra(self, s, t):
        """Dijkstra with Johnson's potential function for MCMF"""
        dist = [float('inf')] * self.N
        parent = [-1] * self.N
        
        dist[s] = 0
        pq = [(0, s)]
        visited = [False] * self.N
        
        while pq:
            d, u = heapq.heappop(pq)
            
            if visited[u]:
                continue
            visited[u] = True
            
            if u == t:
                break
                
            for v in self.graph[u]:
                if self.cap[(u, v)] > 0 and not visited[v]:
                    reduced_cost = self.cost[(u, v)] + self.h[u] - self.h[v]
                    new_dist = dist[u] + reduced_cost
                    
                    if new_dist < dist[v]:
                        dist[v] = new_dist
                        parent[v] = u
                        heapq.heappush(pq, (new_dist, v))
        
        if dist[t] != float('inf'):
            for i in range(self.N):
                if visited[i]:
                    self.h[i] += dist[i]
        
        return dist[t] != float('inf'), parent

    def mcmf(self, s, t):
        max_flow = 0
        min_cost = 0
        
        while True:
            has_path, parent = self.f(s, t)
            if not has_path:
                break

            flow = float('inf')
            curr = t
            while curr != s:
                prev = parent[curr]
                flow = min(flow, self.cap[(prev, curr)])
                curr = prev

            curr = t
            while curr != s:
                prev = parent[curr]
                self.cap[(prev, curr)] -= flow
                self.cap[(curr, prev)] += flow
                min_cost += flow * self.cost[(prev, curr)]
                curr = prev
            max_flow += flow

            if max_flow >= self.cutoff:
                break
        return max_flow, min_cost
    

class MCMF_BF:
    def __init__(self, N, cutoff, pathfunc=None):
        self.N = N
        self.graph = [[] for _ in range(N)]
        self.cutoff = cutoff
        self.cap = [[0] * N for _ in range(N)]
        self.cost = [[0] * N for _ in range(N)]
        
        self.f = self.spfa

    def add_edge(self, u, v, capacity, cost):
        self.graph[u].append(v)
        self.graph[v].append(u)
        self.cap[u][v] += capacity
        self.cost[u][v] = cost
        self.cost[v][u] = -cost
    
    def spfa(self, s, t):
        dist = [float('inf')] * self.N
        parent = [-1] * self.N
        in_queue = [False] * self.N
        
        dist[s] = 0
        queue = deque([s])
        in_queue[s] = True
        
        while queue:
            u = queue.popleft()
            in_queue[u] = False
            
            for v in self.graph[u]:
                if self.cap[u][v] > 0 and dist[u] + self.cost[u][v] < dist[v]:
                    dist[v] = dist[u] + self.cost[u][v]
                    parent[v] = u
                    if not in_queue[v]:
                        queue.append(v)
                        in_queue[v] = True
        
        return dist[t] != float('inf'), parent, dist
    

    def dfs(self, u, t, pushed, min_cost_dist):
        if u == t:
            return pushed
        
        if self.dfs_visited[u]:
            return 0
        self.dfs_visited[u] = True
        
        total_flow = 0
        for v in self.graph[u]:
            if (self.cap[u][v] > 0 and 
                min_cost_dist[v] == min_cost_dist[u] + self.cost[u][v] and
                pushed > total_flow):
                
                flow = self.dfs(v, t, 
                                min(pushed - total_flow, self.cap[u][v]), 
                                min_cost_dist)
                if flow > 0:
                    self.cap[u][v] -= flow
                    self.cap[v][u] += flow
                    total_flow += flow
        
        self.dfs_visited[u] = False
        return total_flow

    def mcmf(self, s, t):
        max_flow = 0
        min_cost = 0
        while True:
            has_path, parent, dist = self.spfa(s, t)
            if not has_path:
                break
            
            self.dfs_visited = [False] * self.N
            
            flow = self.dfs(s, t, float('inf'), dist)
            if flow == 0:
                break
                
            max_flow += flow
            min_cost += flow * dist[t]
            
            if max_flow >= self.cutoff:
                break
        
        return min_cost, max_flow
    

from collections import deque
import heapq
class BipartiteMCMF:
    def __init__(self, n_employees, n_jobs):
        self.n_emp = n_employees
        self.n_jobs = n_jobs
        self.N = n_employees + n_jobs + 2
        self.graph = [[] for _ in range(self.N)]
        
        self.source = 0
        self.sink = n_employees + n_jobs + 1
        
        self.build()
    
    def build(self):
        for i in range(1, self.n_emp + 1):
            self.add_internal(self.source, i, 1, 0)
    
        for j in range(self.n_emp + 1, self.n_emp + self.n_jobs + 1):
            self.add_internal(j, self.sink, 1, 0)
    
    def add_internal(self, u, v, capacity, cost):
        forward_idx = len(self.graph[v])
        self.graph[u].append([v, capacity, cost, forward_idx])
        backward_idx = len(self.graph[u]) - 1
        self.graph[v].append([u, 0, -cost, backward_idx])
    
    def add_edge(self, emp, job, cost):
        emp_node = emp + 1
        job_node = job + self.n_emp + 1
        self.add_internal(emp_node, job_node, 1, cost)
    
    def _dijkstra(self):
        dist = [float('inf')] * self.N
        parent = [(-1, -1)] * self.N
        
        dist[self.source] = 0
        pq = [(0, self.source)]
        
        while pq:
            d, u = heapq.heappop(pq)
            
            if d > dist[u]:
                continue
            
            if u == self.sink:
                break
            
            for i, (v, capacity, cost, _) in enumerate(self.graph[u]):
                if capacity > 0:
                    new_dist = dist[u] + cost
                    if new_dist < dist[v]:
                        dist[v] = new_dist
                        parent[v] = (u, i)
                        heapq.heappush(pq, (new_dist, v))
        
        return dist[self.sink] != float('inf'), parent
    
    def _spfa(self):
        dist = [float('inf')] * self.N
        parent = [(-1, -1)] * self.N
        in_queue = [False] * self.N
        
        dist[self.source] = 0
        queue = deque([self.source])
        in_queue[self.source] = True
        
        while queue:
            u = queue.popleft()
            in_queue[u] = False
            
            for i, (v, capacity, cost, _) in enumerate(self.graph[u]):
                if capacity > 0:
                    new_dist = dist[u] + cost
                    if new_dist < dist[v]:
                        dist[v] = new_dist
                        parent[v] = (u, i)
                        
                        if not in_queue[v]:
                            queue.append(v)
                            in_queue[v] = True
        
        return dist[self.sink] != float('inf'), parent
    
    def flow(self, parent):
        flow = 1
        path_cost = 0
        
        curr = self.sink
        while curr != self.source:
            prev_node, edge_idx = parent[curr]

            next_node, capacity, cost, reverse_idx = self.graph[prev_node][edge_idx]
            
            path_cost += cost
            
            self.graph[prev_node][edge_idx][1] -= flow
            self.graph[curr][reverse_idx][1] += flow 
            
            curr = prev_node
        
        return path_cost
    
    def mcmf(self, algorithm='spfa'):
        total_matches = 0
        total_cost = 0
        path_finder = self._spfa if algorithm == 'spfa' else self._dijkstra
        
        while True:
            has_path, parent = path_finder()
            if not has_path:
                break
            
            path_cost = self.flow(parent)
            total_matches += 1
            total_cost += path_cost
            if total_matches >= min(self.n_emp, self.n_jobs):
                break
        
        return total_matches, total_cost