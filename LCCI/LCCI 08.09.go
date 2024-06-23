func generateParenthesis(n int) []string {
    f := make([][]string, n + 1)
    f[1] = []string{"()"}
    for i := 2; i <= n; i++ {
        for j := i - 1; j >= 0; j-- {
            for _, p := range f[i-j-1] {
                for _, q := range f[j] {
                    f[i] = append(f[i], p + "(" + q + ")")
                }
            }
        }
    }

    return f[n]
}