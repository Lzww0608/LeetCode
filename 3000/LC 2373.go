func largestLocal(grid [][]int) [][]int {
    n := len(grid)
    ans := make([][]int, n - 2)
    for i := range ans {
        ans[i] = make([]int, n - 2)
        for j := 0; j < n - 2; j++ {
            mx := 0
            for p := i; p <= i + 2; p++ {
                for q := j; q <= j + 2; q++ {
                    mx = max(mx, grid[p][q])
                } 
            }

            ans[i][j] = mx
        }
    }

    return ans
}