func numSquares(n int) int {
    f := make([]int, n + 1)

    for i := 1; i <= n; i++ {
        f[i] = i
        for j := 1; j * j <= i; j++ {
            x := j * j
            f[i] = min(f[i-x] + 1, f[i])
        }
    }

    return f[n]
}