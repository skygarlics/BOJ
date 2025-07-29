def catalan(n):
    # Returns the nth Catalan number
    if n == 0:
        return 1
    if n == 1:
        return 1
    c = 1
    for i in range(2, n + 1):
        c = c * (2 * (2 * i - 1)) // (i + 1)
    return c