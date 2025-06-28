func findLength(a []int, b []int) (ans int) {
    n, m := len(a), len(b)
    f := make([][]int, n + 1)
    for i := range f {
        f[i] = make([]int, m + 1)
    }

    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if a[i] == b[j] {
                f[i+1][j+1] = f[i][j] + 1
                ans = max(ans, f[i+1][j+1])
            } 
        }
    }

    return 
}