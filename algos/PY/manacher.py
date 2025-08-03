def manacher(s):
    processed = '^#' + '#'.join(s) + '#$'
    n = len(processed)
    p = [0] * n
    center = right = 0
    for i in range(1, n - 1):
        if i < right:
            mirror = 2 * center - i
            p[i] = min(right - i, p[mirror])
        while processed[i + p[i] + 1] == processed[i - p[i] - 1]:
            p[i] += 1
        if i + p[i] > right:
            center, right = i, i + p[i]
    return p

def count_palindromes(s):
    p = manacher(s)
    count = 0
    for radius in p:
        count += (radius + 1) // 2
    return count