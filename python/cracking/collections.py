"""Searching and Sorting"""
def sorted_merge(arr1, arr2):
    idx2 = 0
    arr1.extend([None] * len(arr2))
    len1 = len(arr1)
    temp = []
    for idx1 in range(len1):
        current = arr1[idx1]
        number2 = arr2[idx2] if idx2 < len(arr2) else None
        if current is None:
            if temp:
                if number2 is None or temp[0] <= number2:
                    arr1[idx1] = temp.pop(0)
                    continue
            arr1[idx1] = number2
            idx2 += 1
            continue
        if temp:
            if number2 is None or temp[0] < current and temp[0] <= number2:
                arr1[idx1] = temp.pop(0)
                temp.append(current)
                continue
        if number2 is not None and number2 < current:
            arr1[idx1] = number2
            idx2 += 1
            temp.append(current)
    return arr1

class Listy:
    def __init__(self, arr):
        if not arr:
            arr = []
        self._arr = sorted(arr)

    def __repr__(self):
        return "{}".format(self._arr)

    def get(self, index):
        if index >= len(self._arr):
            return -1
        return self._arr[index]

def listy_search(listy, num):
    low_idx = 0
    high_idx = 1
    if listy.get(low_idx) > num:
        return -1
    while listy.get(high_idx) < num and listy.get(high_idx) != -1:
        low_idx = high_idx
        high_idx *= 10
    if listy.get(high_idx) == num:
        return listy.get(high_idx)
    while low_idx < high_idx and listy.get(low_idx) != -1:
        if listy.get(low_idx) == num:
            return listy.get(low_idx)
        low_idx += 1
    if low_idx == high_idx:
        return -1
    return listy.get(low_idx)

# Stream Rank
class StreamRanker:
    def __init__(self):
        self._map = dict()

    def __repr__(self):
        return "{}".format(self._map)

    def track(self, val):
        new_rank = 0
        for num, rank in self._map.items():
            if num < val and new_rank <= rank:
                new_rank = rank + 1
            if num == val:
                rank += 1
                new_rank = rank
            if num > val:
                rank += 1
            self._map.update({num: rank})
        self._map.update({val: new_rank})

    def get_rank(self, num):
        return self._map.get(num, -1)

def get_peak(arr, idx):
    left = arr[idx-1] if idx > 0 else None
    right = arr[idx+1] if idx < len(arr)-1 else None
    if not left:
        if not right:
            return idx
        peak = max(arr[idx], right)
    elif not right:
        peak = max(arr[idx], left)
    else:
        peak = max(arr[idx], max(left, right))
    if peak == left:
        return idx-1
    if peak == right:
        return idx+1
    return idx

def peaks_and_valleys(arr):
    size = len(arr)
    for idx in range(1, size, 2):
        peak_idx = get_peak(arr, idx)
        if idx != peak_idx:
            arr[idx], arr[peak_idx] = arr[peak_idx], arr[idx]
    return arr
