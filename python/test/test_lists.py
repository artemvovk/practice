"""Test hackerrank.lists"""
import unittest
from unittest import mock

import random
import string
from hackerrank import lists

class TestLists(unittest.TestCase):

    @staticmethod
    def random_string(length):
        return ''.join(random.choice(string.ascii_letters) for m in range(0, length))

    @mock.patch('hackerrank.lists.input', create=True)
    def test_comprehensions(self, mocked_input):
        mocked_input.side_effect = [random.randint(1, 5),
                                    random.randint(1, 5),
                                    random.randint(1, 5), 3]
        cuboid = lists.comprehensions()
        print("For sum {} generated cuboid:\n\t {}".format(3, cuboid))
        self.assertTrue(len(cuboid) > 1)

    @mock.patch('hackerrank.lists.input', create=True)
    def test_runner_up(self, mocked_input):
        inp = [random.randint(1, 5),
               random.randint(-100, -1),
               random.randint(1, 9),
               random.randint(1, 100),
               random.randint(1, 9),
               random.randint(1, 10000),
               random.randint(1, 10000)]

        mocked_input.side_effect = [len(inp), ' '.join(map(str, inp))]
        runner = lists.runner_up()
        print("For list {}\n\t runner up is {}".format(inp, runner))
        self.assertTrue(runner != 0)

    def test_nested(self):
        inp = [TestLists.random_string(9),
               random.uniform(1, 100),
               TestLists.random_string(9),
               random.uniform(1, 100),
               TestLists.random_string(9),
               random.uniform(1, 100),
               TestLists.random_string(9),
               random.uniform(1.2, 50.1),
               TestLists.random_string(9),
               random.uniform(1, 9)]
        inp = [len(inp)] + inp
        secondary = lists.nested(inp)
        print(secondary)
        self.assertTrue(len(secondary) > 0)

    def test_matrix_script(self):
        n = 7
        j = 3
        matrix = [
            'Tsi',
            'h%x',
            'i #',
            'sM ',
            '$a ',
            '#t%',
            'ir!',
        ]
        print(lists.decode_matrix(n, j, matrix))
        self.assertTrue(n > j)
