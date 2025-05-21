def get_primeset(n):
    ret = set()
    # slow prime factorization
    p = 2
    while p * p <= n:
        if n % p == 0:
            ret.add(p)
            while n % p == 0:
                n //= p
        p += 1
    if n > 1:
        ret.add(n)
    return ret

def get_primelist(n):
    ret = []
    # slow prime factorization
    p = 2
    while p * p <= n:
        if n % p == 0:
            ret.append(p)
            while n % p == 0:
                n //= p
        p += 1
    if n > 1:
        ret.append(n)
    return ret

K, L = map(int, input().split())
