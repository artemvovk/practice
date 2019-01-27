"""Hard Stuff"""
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
