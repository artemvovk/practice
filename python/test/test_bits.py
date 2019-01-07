"""Test cracking.bits"""
import unittest

import random
from cracking import bits

class TestBits(unittest.TestCase):

    def test_get_bit(self):
        number = random.randint(1, 100)
        offset = random.randint(1, 5)
        print("Bit of {0:b}\nat {1} is {2:b}"
              .format(number, offset, bits.get_bit(number, offset)))
        self.assertTrue(number > 0)

    def test_set_bit(self):
        number = random.randint(32, 100)
        val = random.randint(0, 1)
        offset = random.randint(1, 5)
        print("Bit of {0:b}\nat {1} to {2}\nresult {3:b}"
              .format(number, offset, val, bits.set_bit(number, val, offset)))
        self.assertTrue(number > 0)

    def test_clear_ms(self):
        number = random.randint(32, 100)
        offset = random.randint(2, 3)
        print("Cleared bits above {0} of {1:b}\nresult {2:b}"
              .format(offset, number, bits.clear_most_significant(number, offset)))
        self.assertTrue(number > 0)

    def test_clear_ls(self):
        number = random.randint(32, 100)
        offset = random.randint(2, 3)
        print("Cleared bits below {0} of {1:b}\nresult {2:b}"
              .format(offset, number, bits.clear_least_significant(number, offset)))
        self.assertTrue(number > 0)

    def test_insert(self):
        number = random.randint(32, 1024)
        insertee = random.randint(8, 16)
        start = 2
        end = 10
        print("Inserting {0:b} into {1:b} for {2} - {3}:\n\t {4:b}"
              .format(number, insertee, start, end, bits.insert(number,
                                                                insertee,
                                                                start, end)))
        self.assertTrue(number > 0)

    def test_ftw(self):
        number = random.randint(256, 2048)
        print("Number {0:b} can have {1} ones"
              .format(number, bits.flip_to_win(number)))
        self.assertTrue(number > 0)
