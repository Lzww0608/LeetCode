func countPalindromicSubsequences(s string) int {
    const MOD int = 1e9 + 7
    n := len(s)
    f := make([][]int, n)
    for i := range f {
        f[i] = make([]int, n)
        f[i][i] = 1
    }

    for d := 1; d < n; d++ {
        for i := 0; i < n - d; i++ {
            j := i + d
            if s[i] != s[j] {
                f[i][j] = f[i+1][j] + f[i][j-1] - f[i+1][j-1]
            } else {
                l, r := i + 1, j - 1
                for l <= r && s[l] != s[i] {
                    l++
                }
                for l <= r && s[r] != s[i] {
                    r--
                }

                if l > r {
                    f[i][j] = f[i+1][j-1] * 2 + 2
                } else if l == r {
                    f[i][j] = f[i+1][j-1] * 2 + 1 
                } else if l < r {
                    f[i][j] = f[i+1][j-1] * 2 - f[l+1][r-1]
                }
            }
            if f[i][j] < 0 {
                f[i][j] += MOD
            }
            f[i][j] %= MOD
        }
    }

    return f[0][n-1]
}