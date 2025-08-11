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

def miller_rabin(n):
    """ Miller-Rabin primality test, ranged in long long int """
    # Check n is in range [2, 2^64-1]
    if n < 2:
        raise ValueError("n must be >= 2")
    if n > 2**64 - 1:
        raise ValueError("n must be <= 2^64-1")
    
    # Handle small primes
    if n in (2, 3):
        return True
    if n % 2 == 0:
        return False
    
    # Write n-1 as d * 2^r
    d = n - 1
    r = 0
    while d % 2 == 0:
        d //= 2
        r += 1

    # Test with bases
    for a in [2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37]:
        if a >= n:
            break
        x = pow(a, d, n)
        if x == 1 or x == n - 1:
            continue
        for _ in range(r - 1):
            x = pow(x, 2, n)
            if x == n - 1:
                break
        else:
            return False
    return True

from random import randint
from math import gcd
def pollard_rho(n, prime_test=miller_rabin):
    """ Pollard's Rho algorithm for integer factorization """
    if prime_test(n):
        return n
    if n == 1:
        return 1
    if n % 2 == 0:
        return 2
    
    x = y = randint(2, n - 1)
    c = randint(1, n - 1)
    d = 1

    while d == 1:
        x = (x**2+c) % n
        y = (y**2+c) % n
        y = (y**2+c) % n
        d = gcd(abs(x - y), n)
        if d == n: return pollard_rho(n, prime_test) # failure, retry with new n

    if prime_test(d): return d
    return pollard_rho(d, prime_test)

    