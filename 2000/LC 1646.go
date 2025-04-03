func getMaximumGenerated(n int) (ans int) {
    if n < 2 {
        return n
    }

    f := make([]int, n + 1)
    f[1], f[2] = 1, 1
    ans = 1
    for i := 3; i <= n; i++ {
        if i & 1 == 0 {
            f[i] = f[i / 2]
        } else {
            f[i] = f[i / 2] + f[i / 2 + 1]
        }
        ans = max(ans, f[i])
    }

    return
}