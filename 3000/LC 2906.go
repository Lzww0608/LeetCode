const MOD int = 12345
func constructProductMatrix(grid [][]int) [][]int {
    m, n := len(grid), len(grid[0])
    ans := make([][]int, m)
    for i := range ans {
        ans[i] = make([]int, n)
    }

    suf := make([]int, m * n + 1)
    suf[m * n] = 1 
    for i := m * n - 1; i >= 0; i-- {
        x := grid[i / n][i % n]
        suf[i] = suf[i + 1] * x % MOD
    }

    pre := 1
    for i := 0; i < m * n; i++ {
        ans[i / n][i % n] = pre * suf[i + 1] % MOD
        pre = pre * grid[i / n][i % n] % MOD
    }

    return ans
}