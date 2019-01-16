"""Testing GeeksForGeeks stuff"""
import random
import pytest
from geeksforgeeks import problems

def create_coins(amount):
    if amount > 1 and amount%2:
        amount -= 1
    coins = [0] * amount
    for index in range(amount):
        coins[index] = random.randint(1, 100)
    return coins

def create_maze(length):
    maze = [['O' for i in range(length)] for i in range(length)]
    # pylint: disable=consider-using-enumerate
    for y in range(len(maze)):
        for x in range(len(maze[y])):
            if not random.randint(55, 100)%7:
                maze[y][x] = 'X'
    return maze

@pytest.mark.parametrize("test_input,expected", [
    (10, 0),
    (20, 0)
])
def test_x_shapes(test_input, expected):
    maze = create_maze(test_input)
    print("\n")
    for row in maze:
        print(row)
    shapes = problems.count_x_shapes(maze)
    print(shapes)
    assert shapes > expected

@pytest.mark.parametrize("test_input,expected", [
    (10, 0),
    (20, 0)
])
def test_coin_game(test_input, expected):
    coins = create_coins(test_input)
    print(coins)
    assert problems.coin_game(coins) > expected
