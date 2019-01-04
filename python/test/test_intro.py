"""Test hackerrank.intro"""
import unittest
from unittest import mock

import random
from hackerrank import intro

class TestIntro(unittest.TestCase):

    @mock.patch('hackerrank.intro.input', create=True)
    def test_loops(self, mocked_input):
        inp = random.randint(1, 100)
        mocked_input.side_effect = [inp]

        out = intro.loops()
        print("\nSquares of {} are:\n\t {}".format(inp, out))
        self.assertFalse(len(out) == 0)

    def test_leap_year(self):
        leap_years = [2000, 2400, 4, 16, 0]
        for year in leap_years:
            print("Year {} is {}".format(year, intro.is_leap(year)))
            self.assertTrue(intro.is_leap(year))

        non_leap_years = [2100, 55, 17, 13, 1234]
        for year in non_leap_years:
            print("Year {} is {}".format(year, intro.is_leap(year)))
            self.assertFalse(intro.is_leap(year))
