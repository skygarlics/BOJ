import sys;inputs=sys.stdin.read().split()
it = iter(map(int,inputs))

T = next(it)

def Dice(arr):
    idx = -1
    def ret(arr):
        nonlocal idx
        while True:
            idx = (idx + 1) % 6
            yield arr[idx]
    return ret(arr)

for _ in range(T):
    N = next(it)
    arr = [next(it) for _ in range(6)]

    dice = Dice(arr)

    idx = 0
    cnt = 0
    while idx < N:
        rolled = next(dice)
        if rolled + idx > N:
            continue
        idx += rolled
        cnt += idx
    print(cnt)