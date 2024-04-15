func generateMatrix(n int) [][]int {
    ans := make([][]int, n)
    for i := range ans {
        ans[i] = make([]int, n)
    }

    level := 0
    for k := 1; k <= n * n; {
        i, j := level, level
        for k <= n * n && j < n - level {
…        for k <= n * n && i > level {
            ans[i][j] = k
            k++
            i--
        } 
        level++
    }

    return ans
}