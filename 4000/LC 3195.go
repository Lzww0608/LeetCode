func minimumArea(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    up, down := -1, -1
    for i := 0; i < m && up == -1; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                up = i 
                break
            }
        }
    }
    for i := m - 1; i >= 0 && down == -1; i-- {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                down = i 
                break
            }
        }
    }

    l, r := -1, -1
    for j := 0; j < n && l == -1; j++ {
        for i := 0; i < m; i++ {
            if grid[i][j] == 1 {
                l = j 
                break
            }
        } 
    }
    for j := n - 1; j >= 0 && r == -1; j-- {
        for i := 0; i < m; i++ {
            if grid[i][j] == 1 {
                r = j 
                break
            }
        }
    }
    fmt.Println(up, down, l, r)
    return (r - l + 1) * (down - up + 1)
}