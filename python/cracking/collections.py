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
