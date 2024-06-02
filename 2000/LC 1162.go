func maxDistance(grid [][]int) (ans int) {
    n := len(grid)

    type pair struct {
        x, y int
    }

    q := []pair{}
    for i := range grid {
        for j, x := range grid[i] {
            if x == 1 {
                q = append(q, pair{i, j})
                grid[i][j] = -1
            }
        }
    }
    if len(q) == 0 || len(q) == n * n {
        return -1
    }

    dirs := [][]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}
    d := -1
    for len(q) > 0 {
        d++
        for k := len(q); k > 0; k-- {
            i, j := q[0].x, q[0].y
            q = q[1:]
            for _, dir := range dirs {
                x, y := i + dir[0], j + dir[1]
                if x < 0 || x >= n || y < 0 || y >= n || grid[x][y] != 0 {
                    continue
                }
                q = append(q, pair{x, y})
                grid[x][y] = -1  // 标记为已访问
            }
        }
    }

    return d
}
