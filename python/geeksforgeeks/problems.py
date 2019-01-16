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
