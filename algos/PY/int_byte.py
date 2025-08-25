int.bit_length
int.bit_count

def ctz(x: int) -> int:
    return (x & -x).bit_length() - 1 if x != 0 else 0