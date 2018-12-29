"""Test hackerrank.intro"""
import unittest
from unittest import mock

from hackerrank import intro

class TestIntro(unittest.TestCase):

    @mock.patch('hackerrank.intro.input', create=True)
    def test_loops(self, mocked_input):
        mocked_input.side_effect = ['5']
        intro.loops()
        self.assertFalse(5 == 10)

    def test_leap_year(self):
        leap_years = [2000, 2400, 4, 16, 0]
        for year in leap_years:
            print("Year {} is {}".format(year, intro.is_leap(year)))
            self.assertTrue(intro.is_leap(year))
        non_leap_years = [2100, 55, 17, 13, 1234]
        for year in non_leap_years:
            print("Year {} is {}".format(year, intro.is_leap(year)))
            self.assertFalse(intro.is_leap(year))
