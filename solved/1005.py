import sys;
from collections import deque

sys.setrecursionlimit(10**6)
inputs=sys.stdin.read().split()
it = iter(map(int, inputs))

class Build:
    id_: int
    consume: int
    parents: list[int]
    def __init__(self, consume):
        self.consume = consume
        self.parents = []

def solve():
    N, K = next(it), next(it)
    builds:list[Build] = [0]
    for idx in range(1, N+1):
        consume = next(it)
        builds.append(Build(consume))
        builds[idx].id_ = idx
    for _ in range(K):
        X, Y = next(it), next(it)
        builds[Y].parents.append(X)

    # dfsing reversely
    W = next(it)

    memo = [-1] * (N+1)

    def dfs(idx):
        if memo[idx] != -1:
            return memo[idx]
        
        if not builds[idx].parents:
            memo[idx] = builds[idx].consume
        else:
            memo[idx] = builds[idx].consume + max(dfs(p) for p in builds[idx].parents)
        return memo[idx]
    
    print(dfs(W))


T = next(it)
for _ in range(T):
    solve()