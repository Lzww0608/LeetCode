func getFood(grid [][]byte) int {
    dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
    m, n := len(grid), len(grid[0])
    vis := make([]bool, m * n)
    q := [][]int{}
    for i := 0; i < m && len(q) == 0; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == '*' {
                q = append(q, []int{i, j})
                vis[i * n + j] = true
                break
            }
        }
    }

    d := 1
    for len(q) > 0 {
        for sz := len(q); sz > 0; sz-- {
            cur := q[0]
            q = q[1:]
            i, j := cur[0], cur[1]
            for _, dir := range dirs {
                x, y := i + dir[0], j + dir[1]
                if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] == 'X' || vis[x * n + y] {
                    continue
                }
                vis[x * n + y] = true
                if grid[x][y] == '#' {
                    return d 
                }
                q = append(q, []int{x, y})
            }
        }
        d++
    }

    return -1
}