import sys;inputs = sys.stdin.read().split()
it = iter(map(int, inputs))

N, M = next(it), next(it)

memories = [0] * (N+1)
for idx in range(1, N+1):
    memories[idx] = next(it)
costs = [0] * (N+1)
for idx in range(1, N+1):
    costs[idx] = next(it)

Cmax = sum(costs)
dp = [-1] * (Cmax+1)
dp[0] = 0

for idx in range(1, N+1):
    cost = costs[idx]
    mem = memories[idx]

    for j in range(Cmax, cost-1, -1):
        if dp[j-cost] != -1:
            dp[j] = max(dp[j], dp[j-cost] + mem)

answer = None
for idx in range(Cmax+1):
    if dp[idx] >= M:
        answer = idx
        break

print(answer)
