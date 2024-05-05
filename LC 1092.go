func shortestCommonSupersequence(s string, t string) string {
    n, m := len(s), len(t)
    f := make([][]int, n + 1)
    for i := range f {
        f[i] = make([]int, m + 1)
        f[i][0] = i
    }
    for j := 1; j < m; j++ {
        f[0][j] = j
    }

    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if s[i] == t[j] {
                f[i+1][j+1] = f[i][j] + 1
            } else {
                f[i+1][j+1] = min(f[i][j+1], f[i+1][j]) + 1
            }
        }
    }

    ans := []byte{}
    i, j := n - 1, m - 1 
    for i >= 0 && j >= 0 {
        if s[i] == t[j] {
            ans = append(ans, s[i])
            i--
            j--
        } else if f[i+1][j+1] == f[i+1][j] + 1 {
            ans = append(ans, t[j])
            j--
        } else {
            ans = append(ans, s[i])
            i--
        }
    }

    slices.Reverse(ans)

    return t[:j+1] + s[:i+1] + string(ans)
}