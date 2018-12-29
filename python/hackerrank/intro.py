""" This is a docstring"""
import sys

if __name__ == '__main__':
    loops()

def loops():
    end = int(input())
    if end < 0:
        sys.exit()
    for i in range(0, end):
        print(i**2)

def is_leap(year):
    leap = False
    if not year % 400:
        leap = True
    if not year % 4 and year % 100:
        leap = True
    return leap
