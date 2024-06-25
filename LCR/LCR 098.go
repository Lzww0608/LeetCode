func uniquePaths(m int, n int) int {
    f := make([]int, n + 1)
    f[1] = 1
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            f[j+1] += f[j] 
        }
    }

    return f[n]
}
