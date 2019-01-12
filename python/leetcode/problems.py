"""https://leetcode.com/problems/valid-number/"""
import re

def is_number(string):
    string = string.strip()
    print("|{}|".format(string))
    pattern = re.compile(r'^[+-]?(\d+(\.\d*)?|\.\d+)(e[-+]?\d+)?$')
    match = bool(pattern.match(string))
    return match
