func hasPath(maze [][]int, start []int, destination []int) bool {
    dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
    m, n := len(maze), len(maze[0])

    q := []int{start[0] * n + start[1]}
    vis := make([]bool, m * n)
    for len(q) > 0 {
        i, j := q[0] / n, q[0] % n
        q = q[1:]
        if i == destination[0] && j == destination[1] {
            return true
        }

        for _, dir := range dirs {
            x, y := i + dir[0], j + dir[1]
            for x >= 0 && x < m && y >= 0 && y < n && maze[x][y] == 0 {
                x, y = x + dir[0], y + dir[1]
            }
            x, y = x - dir[0], y - dir[1]
            if x >= 0 && x < m && y >= 0 && y < n && maze[x][y] == 0 && !vis[x * n + y] {
                vis[x * n + y] = true
                q = append(q, x * n + y)
            }
        }
    }

    return false
}