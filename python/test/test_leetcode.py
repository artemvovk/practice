"""Testing LeetCode stuff"""
import random
import pytest
from leetcode import problems

@pytest.mark.parametrize("test_input,expected", [
    ("0", True),
    (" 0.1", True),
    ("abc", False),
    ("1 a", False),
    ("2e10", True),
    (" -90e3", True),
    (" 1e", False),
    ("e3", False),
    (" 6e-1", True),
    (" 99e2.5", False),
    ("53.5e93", True),
    (" --6", False),
    ("-+3", False),
    ("95a54e53", False),
])
def test_is_number(test_input, expected):
    assert problems.is_number(test_input) == expected

def create_tree(height):
    root = problems.TreeNode(height)
    current_node = root
    leftovers = []
    trunk = height
    for count in range(trunk):
        if count%2:
            current_node.left = problems.TreeNode(random.randint(0, 30))
            current_node.right = problems.TreeNode(random.randint(0, 30))
            leftovers.append(current_node.right)
            current_node = current_node.left
        else:
            current_node.right = problems.TreeNode(random.randint(0, 30))
            current_node.left = problems.TreeNode(random.randint(0, 30))
            leftovers.append(current_node.left)
            current_node = current_node.right
    for node in leftovers:
        node.right = problems.TreeNode(random.randint(0, 30))
        node.left = problems.TreeNode(random.randint(0, 30))
    print("Total made: {}".format(trunk + len(leftovers) + 2*len(leftovers)))
    return root

@pytest.mark.parametrize("test_input,expected", [
    (10, 0),
    (20, 0),
    (300, 0)
])
def test_min_camera_cover(test_input, expected):
    root = create_tree(test_input)
    cams, node = problems.min_cameras(0, root)
    if cams == 0:
        cams += 1
    print("{} cameras needed with root {}".format(cams, node))
    assert cams > expected
