import sys
inputs = sys.stdin.read().split()
it = iter(map(int, inputs))

N = next(it)
houses = []
for _ in range(N):
    houses.append((next(it), next(it), next(it)))

INF = 10**9

ans = INF
R,G,B = 0,0,0

for first in range(3):
    if first == 0:
        R,G,B = houses[0][0], INF, INF
    elif first == 1:
        R,G,B = INF, houses[0][1], INF
    else:
        R,G,B = INF, INF, houses[0][2]

    for house in houses[1:]:
        R,G,B = min(G, B) + house[0], min(R, B) + house[1],min(R, G) + house[2]

    cand = INF
    if first == 0:
        cand = min(G, B)
    elif first == 1:
        cand = min(R, B)
    else:
        cand = min(R, G)
    ans = min(ans, cand)

print(ans)