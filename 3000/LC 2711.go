func differenceOfDistinctValues(grid [][]int) [][]int {
    m, n := len(grid), len(grid[0])
    ans := make([][]int, m)
    for i := range ans {
        ans[i] = make([]int, n)
    }
    st := make(map[int]bool)
    // k = i - j + n 
    for k := 1; k < m + n; k++ {
        clear(st)
        minJ := max(n - k, 0)
        maxJ := min(m + n - 1 - k, n - 1)

        for j := minJ; j <= maxJ; j++ {
            i := k + j - n 
            ans[i][j] = len(st)
            st[grid[i][j]] = true
        }

        clear(st)
        for j := maxJ; j >= minJ; j-- {
            i := k + j - n 
            ans[i][j] = abs(ans[i][j] - len(st))
            st[grid[i][j]] = true 
        }
    }

    return ans
}

func abs(x int) int {
    if x < 0 {
        return -x
    }

    return x
}