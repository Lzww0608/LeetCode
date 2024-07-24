func minCut(s string) int {
    n := len(s)

    f := make([][]bool, n)
    for i := range f {
        f[i] = make([]bool, n)
        for j := range f[i] {
            f[i][j] = true
        }
    }

    for i := n - 1; i >= 0; i-- {
        for j := i + 1; j < n; j++ {
            if s[i] == s[j] {
                f[i][j] = f[i+1][j-1]
            } else {
                f[i][j] = false
            }
        }
    }

    dp := make([]int, n)
    for i := range dp {
        dp[i] = 0x3f3f3f3f
    }

    for i := range dp {
        if f[0][i] {
            dp[i] = 0
        } else {
            for j := i - 1; j >= 0; j-- {
                if f[j+1][i] {
                    dp[i] = min(dp[i], dp[j] + 1)
                }
            }
        }
    }

    return dp[n-1]
}