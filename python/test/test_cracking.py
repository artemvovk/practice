"""Testing Cracking stuff"""
import random
import pytest
from cracking import oop

@pytest.mark.parametrize("test_input,expected", [
    (10, False),
])
def test_minesweeper(test_input, expected):
    minesmap = oop.Map(test_input, random.randint(0, 2))
    minesmap.click(2, 2)
    minesmap.click(random.randint(0, test_input),
                   random.randint(0, test_input))
    print("\n")
    print(minesmap)
    assert not expected
