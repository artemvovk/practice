""" HackerRank Python Practice """
if __name__ == '__main__':
    print("Run me from unit tests")

def loops():
    end = int(input())
    out = []
    if end < 0:
        return 0
    for i in range(0, end):
        out.append(i**2)
    return out

def is_leap(year):
    leap = False
    if not year % 400:
        leap = True
    if not year % 4 and year % 100:
        leap = True
    return leap
