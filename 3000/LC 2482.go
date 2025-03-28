func onesMinusZeros(grid [][]int) [][]int {
    m, n := len(grid), len(grid[0])
    ans := make([][]int, m)
    for i := range ans {
        ans[i] = make([]int, n)
    }

    for i := 0; i < m; i++ {
        cnt := 0
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                cnt++
            } else {
                cnt--
            }
        }

        for j := 0; j < n; j++ {
            ans[i][j] = cnt
        }
    }

    for j := 0; j < n; j++ {
        cnt := 0
        for i := 0; i < m; i++ {
            if grid[i][j] == 1 {
                cnt++
            } else {
                cnt--
            }
        } 
        for i := 0; i < m; i++ {
            ans[i][j] += cnt
        }
    }

    return ans
}