func findColumnWidth(grid [][]int) []int {
    n := len(grid[0])
    ans := make([]int, n)
    for i := 0; i < n; i++ {
        mx, mn := 0, 0
        for j := range grid {
            mn = min(mn, grid[j][i])
            mx = max(mx, grid[j][i])
        }
        ans[i] = max(len(strconv.Itoa(mx)), len(strconv.Itoa(mn)))
    }

    return ans
}