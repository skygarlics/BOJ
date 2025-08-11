def merge_inv_cnt(arr):
    if len(arr) <= 1:
        return arr, 0
    
    mid = len(arr) // 2
    left, left_inv = merge_inv_cnt(arr[:mid])
    right, right_inv = merge_inv_cnt(arr[mid:])
    
    merged = []
    i = j = 0
    inv_count = left_inv + right_inv
    
    while i < len(left) and j < len(right):
        if left[i] <= right[j]:
            merged.append(left[i])
            i += 1
        else:
            merged.append(right[j])
            inv_count += len(left) - i
            j += 1
    
    merged.extend(left[i:])
    merged.extend(right[j:])
    
    return merged, inv_count

class FenwickTree:
    def __init__(self, size):
        self.size = size
        self.tree = [0] * (size + 1)
    
    def update(self, idx):
        while idx <= self.size:
            self.tree[idx] += 1
            idx += idx & (-idx)
    
    def query(self, idx):
        result = 0
        while idx > 0:
            result += self.tree[idx]
            idx -= idx & (-idx)
        return result

def fenwick_inv_cnt(positions):
    n = len(positions)
    if n <= 1:
        return 0
    
    sorted_pos = sorted(set(positions))
    coord_map = {pos: i+1 for i, pos in enumerate(sorted_pos)}
    compressed = [coord_map[pos] for pos in positions]
    
    fenwick = FenwickTree(len(sorted_pos))
    inv_count = 0

    for i in range(n-1, -1, -1):
        if compressed[i] > 1:
            inv_count += fenwick.query(compressed[i] - 1)
        fenwick.update(compressed[i])
    
    return inv_count