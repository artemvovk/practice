"""Testing LeetCode stuff"""
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
