func canPartitionGrid(grid [][]int) bool {
    sum := 0
    for i := range len(grid) {
        for j := range len(grid[0]) {
            sum += grid[i][j]
        }
    }

    check := func(a [][]int) bool {
        m, n := len(a), len(a[0])
        f := func() bool {
            vis := make(map[int]bool)
            vis[0] = true 
            cur := 0
            for i := range m {
                for j := range n {
                    x := a[i][j]
                    cur += x 
                    if i > 0 || j == 0 || j == n - 1 {
                        vis[x] = true
                    }
                }
                if n == 1 {
                    k := cur * 2 - sum
                    if k == 0 || k == a[0][0] || k == a[i][0] {
                        return true
                    }
                } else if vis[cur * 2 - sum] {
                    return true
                }

                if i == 0 {
                    for j := range n {
                        vis[a[i][j]] = true
                    }
                }
            }

            return false
        }
        
        if f() {
            return true 
        }
        for i := 0; i + i < m; i++ {
            a[i], a[m - i - 1] = a[m - i - 1], a[i]
        }
        return f()
    }

    return check(grid) || check(rotate(grid))
}

func rotate(a [][]int) [][]int {
    m, n := len(a), len(a[0])
    b := make([][]int, n)
    for i := range n {
        b[i] = make([]int, m)
    }

    for i := range m {
        for j := range n {
            b[j][m - 1 - i] = a[i][j]
        }
    }

    return b
}