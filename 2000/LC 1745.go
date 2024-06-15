func checkPartitioning(s string) bool {
    n := len(s)
    f := make([][]bool, n + 1)
    for i := range f {
        f[i] = make([]bool, n + 1)
        f[i][i] = true
    }

    check := func(i, j int) {
        for i >= 0 && i < n && j >= 0 && j < n {
            if s[i] == s[j] {
                f[i][j] = true
                i--
                j++
            } else {
                break
            }
        }
    }

    for i := range s {
        check(i, i)
        check(i, i + 1)
    }

    for i := range s {
        for j := n - 1; j > i; j-- {
            if f[0][i] && f[j][n-1] && f[i+1][j-1] {
                return true
            }
        }
    }

    return false
}
