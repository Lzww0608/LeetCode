func maxKilledEnemies(grid [][]byte) (ans int) {
    m, n := len(grid), len(grid[0])
    a := make([][]int, m)

    pre := 0
    for i := 0; i < m; i++ {
        a[i] = make([]int, n)

        pre = 0
        for j := 0; j < n; j++ {
            a[i][j] += pre
            if grid[i][j] == 'E' {
                pre++
            } else if grid[i][j] == 'W' {
                pre = 0
            } 
        }

        pre = 0
        for j := n - 1; j >= 0; j-- {
            a[i][j] += pre
            if grid[i][j] == 'E' {
                pre++
            } else if grid[i][j] == 'W' {
                pre = 0
            } 
        }
    }

    for j := 0; j < n; j++ {
        pre = 0
        for i := 0; i < m; i++ {
            a[i][j] += pre
            if grid[i][j] == 'E' {
                pre++
            } else if grid[i][j] == 'W' {
                pre = 0
            } 
        }

        pre = 0
        for i := m - 1; i >= 0; i-- {
            a[i][j] += pre
            if grid[i][j] == 'E' {
                pre++
            } else if grid[i][j] == 'W' {
                pre = 0
            } 
        }
    }

    

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == '0' {
                ans = max(ans, a[i][j])
            }
        }
    }

    return
}