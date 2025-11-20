func climbStairs(n int, costs []int) int {
    f := make([]int, n + 1)

    for i := 1; i <= n; i++ {
        f[i] = math.MaxInt32
        for j := i - 1; j >= max(0, i - 3); j-- {
            f[i] = min(f[i], f[j] + costs[i - 1] + (i - j) * (i - j))
        }
    }

    return f[n]
}