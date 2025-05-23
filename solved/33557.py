import sys;lines=sys.stdin.read().splitlines()

def mul(a:str, b:str):
    ret = []
    max_len = max(len(a), len(b))
    if len(a) < max_len:
        a = '1' * (max_len - len(a)) + a
    if len(b) < max_len:
        b = '1' * (max_len - len(b)) + b
    for i in range(max_len):
        x, y = int(a[i]), int(b[i])
        ret += [str(x*y)]
    return ''.join(ret)

for line in lines[1:]:
    a, b = line.split()
    mul1 = int(mul(a, b))
    mul2 = int(a) * int(b)
    print(int(mul1 == mul2))