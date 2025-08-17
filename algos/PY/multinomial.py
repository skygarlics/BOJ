from math import factorial
def multinomial(counts):
    total = sum(counts)
    res = factorial(total)
    for c in counts:
        res //= factorial(c)
    return res