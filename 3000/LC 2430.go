func deleteString(s string) int {
    n := len(s)
    if allEqual(s) {
        return n
    }
    //longest common prefix
    lcp := make([][]int, n + 1)
    lcp[n] = make([]int, n + 1)
    for i := n - 1; i >= 0; i-- {
        lcp[i] = make([]int, n + 1)
        for j := n - 1; j >= i; j-- {
            if s[i] == s[j] {
                lcp[i][j] = lcp[i+1][j+1] + 1
            }
        }
    }

    f := make([]int, n + 1)
    for i := n - 1; i >= 0; i-- {
        for j := 0; i + j <= n; j++ {
            if lcp[i][i+j] >= j {
                f[i] = max(f[i], f[i+j] + 1)
            }
        }
    }

    return f[0]
}

func allEqual(s string) bool {
    for i := range s {
        if s[i] != s[0] {
            return false
        }
    }

    return true
}