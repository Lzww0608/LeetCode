func minCut(s string) int {
    n := len(s)

    f := make([][]bool, n)
    for i := range f {
        f[i] = make([]bool, n)
        for j := range f[i] {
            f[i][j] = true
        }
    }

    for l := n - 1; l >= 0; l-- {
        for r := l + 1; r < n; r++ {
            if s[l] == s[r] {
                f[l][r] = f[l+1][r-1]
            } else {
                f[l][r] = false
            }
        }
    }

    dp := make([]int, n)
    for i := range dp {
        dp[i] = math.MaxInt32
    }

    for i := range f {
        if f[0][i] {
            dp[i] = 0
        } else {
            for j := 0; j < i; j++ {
                if f[j+1][i] {
                    dp[i] = min(dp[i], dp[j] + 1)
                } 
            }
        }
    }

    return dp[n-1]
}