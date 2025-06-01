def ccw(a, b, c):
    """Check if the points a, b, c are in counter-clockwise order."""
    return (b[0] - a[0]) * (c[1] - a[1]) > (b[1] - a[1]) * (c[0] - a[0])

def intersect(a, b, c, d):
    """Check if the line segments ab and cd intersect."""
    return ccw(a, b, c) != ccw(a, b, d) and ccw(c, d, a) != ccw(c, d, b)