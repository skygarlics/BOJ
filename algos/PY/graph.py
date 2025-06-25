
# CSR based graph
from array import array
class Graph:  
    def __init__(self, N):
        self.headers = array('I', [1 << 30]) * N
        self.next_edge = array('I')
        self.edge = array('I')

    def connect(self, u, v):
        ptr = self.headers[u]
        self.headers[u] = len(self.next_edge)
        self.next_edge.append(ptr)
        self.edge.append(v)

    def neighbors(self, u):
        ptr = self.headers[u]
        while ptr < len(self.next_edge) and self.next_edge[ptr] != 1 << 30:
            yield self.edge[ptr]
            ptr = self.next_edge[ptr]


# flat array based graph
from array import array
INF = 1 << 30
class Graph:
    def __init__(self, N: int, M: int):
        self.N = N
        self.M = M
        self.start = array('I', [0] * (N + 1))
        self.edges = array('I', [0] * (2 * M))

        self.tmp = [[]for _ in range(N)]

    def add(self, u: int, v: int):
        self.tmp[u].append(v)
        self.tmp[v].append(u)

    def build(self):
        idx = 0
        for u in range(self.N):
            self.start[u] = idx
            for v in self.tmp[u]:
                self.edges[idx] = v
                idx += 1
        self.start[self.N] = idx

        for u in range(self.N):
            offset = self.start[u]
            for idx, val in enumerate(self.tmp[u]):
                self.edges[offset + idx] = val

        del self.tmp
    
    def neighbors(self, u: int):
        start = self.start[u]
        end = self.start[u + 1]
        for i in range(start, end):
            yield self.edges[i]
