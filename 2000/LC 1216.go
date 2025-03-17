func isValidPalindrome(s string, k int) bool {
    n := len(s)
    f := make([][]int, n + 1)
    for i := range f {
        f[i] = make([]int, n + 1)
    }

    for i := n - 1; i >= 0; i-- {
        f[i][i] = 1
        for j := i + 1; j < n; j++ {
            if s[i] == s[j] {
                f[i][j] = f[i+1][j-1] + 2
            } else {
                f[i][j] = max(f[i+1][j], f[i][j-1])
            }
        }
    }

    return f[0][n-1] >= n - k
}