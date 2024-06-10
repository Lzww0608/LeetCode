func uniquePaths(m int, n int) int {
    f := make([]int, n)
    for i := range f {
        f[i] = 1
    }
    for i := 1; i < m; i++ {
        f[0] = 1
        for j := 1; j < n; j++ {
            f[j] += f[j-1]
        } 
    }

    return f[n-1]
}
