"""Hard Stuff"""
import math
from . import bits
def no_plus_add(num1, num2):
    nstr1 = str(num1)
    nstr2 = str(num2)
    mlen = max(len(nstr1), len(nstr2))
    nstr1 = nstr1[::-1].ljust(mlen, "0")
    nstr2 = nstr2[::-1].ljust(mlen, "0")

    carry_over = "0"
    out = ""
    addition_matrix = [
        ["0", "1", "2", "3", "4", "5", "6", "7", "8", "9"],
        ["1", "2", "3", "4", "5", "6", "7", "8", "9", "10"],
        ["2", "3", "4", "5", "6", "7", "8", "9", "10", "11"],
        ["3", "4", "5", "6", "7", "8", "9", "10", "11", "12"],
        ["4", "5", "6", "7", "8", "9", "10", "11", "12", "13"],
        ["5", "6", "7", "8", "9", "10", "11", "12", "13", "14"],
        ["6", "7", "8", "9", "10", "11", "12", "13", "14", "15"],
        ["7", "8", "9", "10", "11", "12", "13", "14", "15", "16"],
        ["8", "9", "10", "11", "12", "13", "14", "15", "16", "17"],
        ["9", "10", "11", "12", "13", "14", "15", "16", "17", "18"]
    ]
    for idx in range(0, mlen):
        char1 = nstr1[idx]
        char2 = nstr2[idx]
        temp_carry = "0"
        nsum = addition_matrix[int(char2)][int(char1)]
        if len(nsum) > 1:
            temp_carry = nsum[0]
            nsum = nsum[1]
        # Semi-redundant carry over check
        if carry_over != "0":
            nsum = addition_matrix[int(nsum)][int(carry_over)]
            if len(nsum) > 1:
                temp_carry = nsum[0]
                nsum = nsum[1]
        out += nsum
        carry_over = temp_carry
    out += carry_over
    out = out[::-1].lstrip("0")
    if not out:
        out = "0"
    return out

def missing_int_by_bit(arr):
    size = len(arr)
    missing = None
    for idx in range(size):
        if not idx%2:
            zero = bits.get_bit(arr[idx], 0)
            if zero != 0:
                missing = idx-1
                zero = bits.get_bit(arr[idx-1], 0)
                if zero != 0:
                    missing = idx
                print("Missing odd {}".format(bin(missing-1)))
        if idx%2 and not missing:
            one = bits.get_bit(arr[idx], int(math.log(max(idx-1, 1), 2)))
            if one != 1:
                missing = idx - 1
        if missing:
            return missing
    return missing

def letters_and_numbers(arr):
    print("".join(arr))
    letters = 0
    numbers = 0
    previous = ''
    index_map = {}
    for idx, char in enumerate(arr):
        if char.isalpha():
            if not previous.isalpha():
                letters = 0
            letters += 1
            if letters == numbers:
                start_idx = idx + 1 - letters - numbers
                index_map.update({start_idx: letters})
                numbers = 0
        else:
            if previous.isalpha():
                numbers = 0
            numbers += 1
            if letters == numbers:
                start_idx = idx + 1 - letters - numbers
                index_map.update({start_idx: letters})
                letters = 0
        previous = char
    max_idx = 0
    max_val = 0
    for key, val in index_map.items():
        if val > max_val:
            max_idx = key
            max_val = val
    print("Max length {}".format(index_map.get(max_idx)))
    return arr[max_idx:(max_idx + index_map.get(max_idx) + index_map.get(max_idx))]
