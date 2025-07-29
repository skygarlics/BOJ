import math
def grahamscan(points):
    """Computes the convex hull of a set of points using Graham's scan algorithm."""
    p0 = min(points, key=lambda p: (p[1], p[0]))
    points = [p for p in points if p != p0]
    
    def dist(a, b):
        return (a[0] - b[0]) ** 2 + (a[1] - b[1]) ** 2

    def ccw(a, b, c):
        return (b[0] - a[0]) * (c[1] - a[1]) - (b[1] - a[1]) * (c[0] - a[0])
    
    def cmp(a, b):
        val = ccw(p0, a, b)
        if val == 0:
            return dist(p0, a) - dist(p0, b)
        return -val
    
    from functools import cmp_to_key
    points.sort(key=cmp_to_key(cmp))

    stack = [p0, points[0]]
    for p in points[1:]:
        while len(stack) > 1:
            if ccw(stack[-2], stack[-1], p) > 0:
                break
            stack.pop()
        stack.append(p)
    return stack
