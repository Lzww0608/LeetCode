func shortestPathBinaryMatrix(grid [][]int) int {
    dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
    n := len(grid)
    if grid[n-1][n-1] != 0 || grid[0][0] != 0 {
        return -1
    }

    q := []int{0}
    grid[0][0] = 1
    d := 1
    for len(q) > 0 {
        for sz := len(q); sz > 0; sz-- {
            cur := q[0]
            q = q[1:]
            i, j := cur / n, cur % n 
            if i == n - 1 && j == n - 1 {
                return d
            }
            for _, dir := range dirs {
                x, y := i + dir[0], j + dir[1]
                if x < 0 || x >= n || y < 0 || y >= n || grid[x][y] != 0 {
                    continue
                }
                grid[x][y] = 1
                q = append(q, x * n + y)
            }
        }
        d++
    }

    return -1
}