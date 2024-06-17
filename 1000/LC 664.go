func strangePrinter(s string) int {
    n := len(s)
    f := make([][]int, n)
    for i := range f {
        f[i] = make([]int, n)
        for j := range f[i] {
            f[i][j] = n
        }
        f[i][i] = 1
    }

    for d := 2; d <= n; d++ {
        for i := 0; i < n; i++ {
            j := i + d - 1
            if j >= n {
                continue
            }
            if s[i] == s[j] {
                f[i][j] = f[i][j-1]
                continue
            }
            for k := i; k < j; k++{
                f[i][j] = min(f[i][j], f[i][k] + f[k+1][j])
            }
        }
    }

    return f[0][n-1]
}