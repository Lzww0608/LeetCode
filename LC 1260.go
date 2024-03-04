func shiftGrid(g [][]int, k int) [][]int {
    n, m := len(g), len(g[0])
    ans := make([][]int, n)
    for i := range ans {
        ans[i] = make([]int, m)
    }

    for i := range g {
        for j := range g[i] {
            ans[(i + (j + k) / m) % n][(j + k) % m] = g[i][j]
        }
    }

    return ans
}