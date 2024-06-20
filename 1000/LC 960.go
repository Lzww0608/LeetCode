func minDeletionSize(s []string) (ans int) {
    m, n := len(s), len(s[0])
    f := make([]int, n)

    for i := range f {
        f[i] = 1
        for j := 0; j < i; j++ {
            for k := 0; k < m; k++ {
                if s[k][i] < s[k][j] {
                    break
                } else if k == m - 1 {
                    f[i] = max(f[j] + 1, f[i])
                }
            }
        }
    }

    for _, x := range f {
        ans = max(ans, x)
    }

    return n - ans
}