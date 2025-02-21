func maxScore(a []int, b []int) int64 {
    n := len(b)
    f := make([][4]int, n)
    for i := range f {
        for j := range f[i] {
            f[i][j] = math.MinInt / 2 
        }
    }
    f[0][0] = a[0] * b[0]
    for i := 1; i < n; i++ {
        f[i][0] = max(f[i-1][0], a[0] * b[i])
        for j := 1; j < 4; j++ {
            f[i][j] = max(f[i-1][j], f[i-1][j-1] + a[j] * b[i]) 
        }
    }

    return int64(f[n-1][3])
}