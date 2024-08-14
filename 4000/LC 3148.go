func maxScore(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    ans := math.MinInt32
    f := make([][]int, m + 1)
    f[0] = make([]int, n + 1)
    for i := range f[0] {
        f[0][i] = math.MaxInt32
    }

    for i, row := range grid {
        f[i+1] = make([]int, n + 1)
        f[i+1][0] = math.MaxInt32
        for j, x := range row {
            mn := min(f[i+1][j], f[i][j+1])
            ans = max(ans, x - mn)
            f[i+1][j+1] = min(x, mn)
        }
    }

    return ans
}