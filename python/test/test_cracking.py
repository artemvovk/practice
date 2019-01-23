"""Testing Cracking stuff"""
import random
import pytest
from cracking import oop, dynamic

def generate_array(size):
    if not size:
        return []
    arr = [0] * size
    for index in range(size):
        arr[index] = random.randint(0, size+1)
    return arr

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

@pytest.mark.parametrize("test_input,expected", [
    (10, 274),
    (3, 4),
    (0, 0),
    (20, 121415),
    (50, 10562230626642)
])
def test_triple_step(test_input, expected):
    assert dynamic.triple_step(test_input) == expected

@pytest.mark.parametrize("test_input,expected", [
    ([1, 2, 3, 4, 4], 4),
    ([1, 2, 3, 4, 5], -1),
    ([0, 2, 4], 0),
    ([1, 2, 3], -1)
])
def test_magic_index(test_input, expected):
    assert dynamic.magic_index(test_input) == expected

@pytest.mark.parametrize("test_input,expected", [
    (3, [2, 1, 0]),
    (2, [1, 0]),
    (4, [3, 2, 1, 0]),
    (5, [4, 3, 2, 1, 0])
])
def test_hanoi_towers(test_input, expected):
    first_tower = dynamic.Tower()
    for i in reversed(range(test_input)):
        first_tower.push(i)
    towers = {
        1: first_tower,
        2: dynamic.Tower(0),
        3: dynamic.Tower(0)
    }
    dynamic.print_towers(towers)

    assert dynamic.hanoi_towers(towers, 1, 3).get(3).stack == expected
