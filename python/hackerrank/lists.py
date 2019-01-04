""" HackerRank Python Practice """
import re

if __name__ == '__main__':
    print("Run me from unit tests")

def comprehensions():
    x = int(input())
    y = int(input())
    z = int(input())
    n = int(input())
    cuboid = list([i, j, k]
                  for i in range(0, x+1)
                  for j in range(0, y+1)
                  for k in range(0, z+1) if (i+j+k) != n)
    return cuboid

def runner_up():
    n = int(input())
    arr = list(map(int, input().split()))
    arr.sort(reverse=True)
    k = arr[0]
    for x in arr:
        if x < k or n == 0:
            k = x
            break
        k = x
    return k

def nested(inp):
    grades = []
    length = inp.pop(0)
    for x in range(0, length, 2):
        name = inp[x]
        score = float(inp[x+1])
        grades.append([name, score])
    grades.sort(key=lambda x: x[1])
    worst_grade = grades[0][1]
    secondary = []
    for student in grades:
        if student[1] > worst_grade:
            if not secondary or student[1] == secondary[0][1]:
                secondary += [student]
            if student[1] < secondary[0][1]:
                break
    return sorted(list(map(lambda student: student[0], secondary)))

def decode_matrix(n, j, matrix):
    alphanum = r'[^a-zA-Z0-9]'
    out = ''
    for idx in range(j):
        for row in range(n):
            out += re.sub(alphanum, ' ', matrix[row][idx])
    return re.sub(' +', ' ', out)
