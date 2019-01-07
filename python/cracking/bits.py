"""Bit Manipulation chapter"""
from math import log

def get_bit(number, index):
    return (number & (1 << index)) != 0

def set_bit(number, val, index):
    mask = ~(1 << index)
    return (number & mask) | (val << index)

def clear_most_significant(number, index):
    mask = (1 << index) - 1
    return number & mask

def clear_least_significant(number, index):
    mask = -1 << (index + 1)
    return number & mask

def insert(number, insertee, start, end):
    if end-start < int(log(insertee, 2)):
        return None
    for index in range(0, int(log(insertee, 2))+1):
        bit = get_bit(insertee, index)
        number = set_bit(number, bit, index+start)
    return number

def flip_to_win(number):
    previous = 0
    current = 0
    running_max = 1

    while number:
        bit = number & 1
        number = number >> 1

        if bit == 1:
            current += bit
        else:
            previous = current
            current = 0

        running_max = max(running_max, previous + current + 1)
    return running_max
