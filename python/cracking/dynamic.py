"""Cracking Dynamic Programming"""
import curses
import time
def triple_step(n):
    step_sizes = {0: 0,
                  1: 1,
                  2: 2,
                  3: 4}
    if n in step_sizes:
        return step_sizes.get(n)
    subpath = 4
    while subpath <= n:
        new_path = step_sizes.get(subpath - 1) + \
            step_sizes.get(subpath - 2) + \
            step_sizes.get(subpath - 3)
        step_sizes.update({subpath: new_path})
        subpath += 1
    return step_sizes.get(n)

def magic_index(arr):
    for i in range(0, len(arr)): # pylint: disable=consider-using-enumerate
        if i > arr[i]:
            return -1
        if arr[i] == i:
            return i
    return -1

# Hanoi Towers
class Tower:
    def __init__(self, height=0):
        self.stack = [0]*height

    def __repr__(self):
        return "{}".format(self.stack)

    def height(self):
        return len(self.stack)

    def push(self, val):
        if not self.stack:
            self.stack = [val]
            return self
        if val > self.peek():
            return None
        self.stack.append(val)
        return self

    def pop(self):
        if not self.stack:
            return None
        return self.stack.pop()

    def peek(self):
        if not self.stack:
            return -1
        return self.stack[-1]

# Do not use in CI
def print_towers(towers):
    stdscr = curses.initscr()
    stdscr.clear()
    stdscr.addstr(0, 0, '{}'.format(towers.get(1)))
    stdscr.addstr(1, 0, '{}'.format(towers.get(2)))
    stdscr.addstr(2, 0, '{}'.format(towers.get(3)))
    stdscr.refresh()
    time.sleep(0.5)
    curses.endwin()

def hanoi_towers(towers, source_tower, target_tower, to_move=None):
    # Move all by default
    if not to_move:
        to_move = towers.get(source_tower).height()

    # Assume we got another tower we can mess with
    other_tower = 2
    for key in towers.keys():
        if key not in [source_tower, target_tower]:
            other_tower = key

    # Base cases
    if to_move == 1:
        towers.get(target_tower).push(towers.get(source_tower).pop())
    elif to_move == 2:
        towers.get(other_tower).push(towers.get(source_tower).pop())
        towers.get(target_tower).push(towers.get(source_tower).pop())
        towers.get(target_tower).push(towers.get(other_tower).pop())
    elif to_move == 3:
        towers.get(target_tower).push(towers.get(source_tower).pop())
        towers.get(other_tower).push(towers.get(source_tower).pop())
        towers.get(other_tower).push(towers.get(target_tower).pop())
        towers.get(target_tower).push(towers.get(source_tower).pop())
        towers.get(source_tower).push(towers.get(other_tower).pop())
        towers.get(target_tower).push(towers.get(other_tower).pop())
        towers.get(target_tower).push(towers.get(source_tower).pop())
    else:
        towers = hanoi_towers(towers, source_tower, other_tower, to_move-1)
        towers = hanoi_towers(towers, source_tower, target_tower, 1)
        towers = hanoi_towers(towers, other_tower, target_tower)
    return towers
