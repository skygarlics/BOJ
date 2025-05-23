import sys
inputs = iter(map(int, sys.stdin.read().split()))

def solve():
    W, L, H = (next(inputs)+2 for _ in range(3))

    PLANE = W * L
    SIZE = PLANE * H

    box = bytearray(SIZE)
    que = []
    raw = 0

    append = que.append # lookup overhead

    nz = 0
    for z in range(1, H-1):
        nz += PLANE
        ny = nz
        for y in range(1, L-1):
            ny += W
            idx = ny
            for x in range(1, W-1):
                idx += 1
                state = next(inputs)
                sb = state == 0
                box[idx] = sb
                if sb:
                    raw += 1
                elif state > 0 :
                    append(idx)
                    
    if raw == 0: return 0

    dirs = (-1, 1, -W, W, -PLANE, PLANE)

    day = -1
    append(0)
    for idx in que:
        if idx == 0:
            day += 1
            if que[-1] == 0:break
            append(0)
            continue
        for d in dirs:
                nidx = idx + d
                if box[nidx]:
                    box[nidx] = 0; raw -= 1; append(nidx)

    if raw == 0:
        return day
    return -1

print(solve())