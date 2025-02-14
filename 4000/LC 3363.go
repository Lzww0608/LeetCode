func maxCollectedFruits(fruits [][]int) (ans int) {
    n := len(fruits)
    for i := 0; i < n; i++ {
        ans += fruits[i][i]
    }

    f := make([][]int, n + 1)
    for i := range f {
        f[i] = make([]int, n + 1)
    }
    f[0][n-1] = fruits[n-1][0]
    for j := 1; j < n; j++ {
        for i := n - 1; i >= max(j + 1, n - j - 1); i-- {
            f[j][i] = max(f[j-1][i], f[j-1][i-1], f[j-1][i+1]) + fruits[i][j]
        }
    }

    ans += f[n-2][n-1]

    g := make([][]int, n + 1)
    for i := range g {
        g[i] = make([]int, n + 1)
    }
    g[0][n-1] = fruits[0][n-1]
    for i := 1; i < n; i++ {
        for j := n - 1; j >= max(i + 1, n - i - 1); j-- {
            g[i][j] = max(g[i-1][j], g[i-1][j-1], g[i-1][j+1]) + fruits[i][j]
        }
    }

    ans += g[n-2][n-1]
    return
}