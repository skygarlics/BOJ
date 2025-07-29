def ccw(a, b, c):
    """Check if the points a, b, c are in counter-clockwise order."""
    return (b[0] - a[0]) * (c[1] - a[1]) > (b[1] - a[1]) * (c[0] - a[0])