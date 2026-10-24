func lastVisitedCell(grid [][]byte) []int {
    m, n := len(grid), len(grid[0])
    vis := make([][]bool, m)
    for i := range vis {
        vis[i] = make([]bool, n)
    }
    vis[0][0] = true
    i, j := 0, 0

    for {
        v := grid[i][j]
        x, y := i, j 
        if v == 'L' {
            y--
        } else if v == 'R' {
            y++
        } else if v == 'U' {
            x--
        } else {
            x++
        }
        if x < 0 || x >= m || y < 0 || y >= n || vis[x][y] {
            break
        }
        vis[x][y] = true
        i, j = x, y
    }

    return []int{i, j}
}