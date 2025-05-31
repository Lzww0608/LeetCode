const MOD int = 1_000_000_007
func numberOfCombinations(s string) (ans int) {
    if s[0] == '0' {
        return
    }

    n := len(s)
    lcp := make([][]int, n + 1)
    for i := range lcp {
        lcp[i] = make([]int, n + 1)
    }

    for i := n - 1; i >= 0; i-- {
        for j := n - 1; j >= 0; j-- {
            if s[i] == s[j] {
                lcp[i][j] = lcp[i + 1][j + 1] + 1 
            }
        }
    }

    less := func(a, b, c int) bool {
        x := lcp[a][b]
        return x >= c - b || s[a + x] < s[b + x]
    }

    f := make([][]int, n)
    for i := range f {
        f[i] = make([]int, n)
    }
    for j := 0; j < n; j++ {
        f[0][j] = 1
    }

    for i := 1; i < n; i++ {
        if s[i] == '0' {
            continue
        }
        for j, k, sum := i, i - 1, 0; j < n; j++ {
            f[i][j] = sum 
            if k < 0 {
                continue
            }
            if s[k] > '0' && less(k, i, j + 1) {
                f[i][j] = (f[i][j] + f[k][i-1]) % MOD
            }
            sum = (sum + f[k][i-1]) % MOD
            k--
        }
    }

    for _, v := range f {
        ans = (ans + v[n-1]) % MOD
    }

    return
}