"""https://leetcode.com/problems/valid-number/"""
import re

# Definition for a binary tree node.
class TreeNode: # pylint: disable=too-few-public-methods
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

    def __repr__(self):
        return "N{}".format(self.val)

def is_number(string):
    string = string.strip()
    print("|{}|".format(string))
    pattern = re.compile(r'^[+-]?(\d+(\.\d*)?|\.\d+)(e[-+]?\d+)?$')
    match = bool(pattern.match(string))
    return match

def min_cameras(cams, node):
    if not node:
        return cams, 1
    cams, lnode = min_cameras(cams, node.left)
    cams, rnode = min_cameras(cams, node.right)

    if lnode == 0 or rnode == 0:
        cams += 1
        return cams, 2
    if lnode == 2 or rnode == 2:
        return cams, 1
    return cams, 0
