"""https://practice.geeksforgeeks.org/"""
def coin_game(coins):
    variations = [[0 for i in range(len(coins))] for i in range(len(coins))]
    for leftover in range(len(coins)):
        for j in range(leftover, len(coins)):
            i = j - leftover
            var1 = 0
            if (i + 2) <= j:
                var1 = variations[i + 2][j]
            var2 = 0
            if (i + 1) <= (j - 1):
                var2 = variations[i + 1][j - 1]
            var3 = 0
            if i <= (j - 2):
                var3 = variations[i][j-2]
            variations[i][j] = max(coins[i] + min(var1, var2),
                                   coins[j] + min(var2, var3))
    print("\t{}".format(variations[0][len(coins)-1]))
    return variations[0][len(coins)-1]

def count_x_shapes(maze):
    def dimensions(maze):
        return len(maze), len(maze[0])
    def traverse(maze, visited, x, y):
        height, width = dimensions(maze)
        if x < 0 or x > width-1:
            return
        if y < 0 or y > height-1:
            return
        if maze[y][x] == 'O' or visited[y][x]:
            return
        visited[y][x] = True
        traverse(maze, visited, x+1, y)
        traverse(maze, visited, x, y+1)
        traverse(maze, visited, x-1, y)
        traverse(maze, visited, x, y-1)
    height, width = dimensions(maze)
    shapes = 0
    visited = [[False for i in range(width)] for i in range(height)]
    for y in range(height):
        for x in range(width):
            if maze[y][x] == 'X' and not visited[y][x]:
                traverse(maze, visited, x, y)
                shapes += 1
    return shapes

def interpolation_search(arr, number):
    arr = hoare_quicksort(arr)
    idx = 0
    low = 0
    high = len(arr) -1
    while low <= high and arr[low] <= number <= arr[high]:
        idx = low + int((
                        (float(high - low) /
                         (arr[high] - arr[low])) *
                        (number - arr[low]))
                        )
        if arr[idx] == number:
            return idx
        if arr[idx] < number:
            low = idx + 1
        else:
            high = idx - 1
    return -1

def hoare_quicksort(arr):
    def partition(arr, low, high):
        pivot = arr[int((low + high) / 2)]
        low -= 1
        high += 1
        while True:
            low += 1
            while low < len(arr) and arr[low] < pivot:
                low += 1

            high -= 1
            while high > 0 and arr[high] > pivot:
                high -= 1
            if low >= high:
                return high
            arr[low], arr[high] = arr[high], arr[low]

    length = len(arr)
    stack = [0] * length
    top = 0
    stack[top] = arr[0]
    top += 1
    stack[top] = length-1

    while top > 0:
        high = stack[top]
        top -= 1
        low = stack[top]
        top -= 1
        pivot = partition(arr, low, high)

        if pivot - 1 > low:
            top += 1
            stack[top] = low
            top += 1
            stack[top] = pivot - 1
        if pivot + 1 < high:
            top += 1
            stack[top] = pivot + 1
            top += 1
            stack[top] = high
    return arr
