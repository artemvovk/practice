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
