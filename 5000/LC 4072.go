const inf = math.MinInt / 2
func maxAlternatingSum(nums []int) int64 {
    n := len(nums)
    f := make([][2]int, n)
    g := make([][2]int, n)
    for i := range n {
        f[i] = [2]int{inf, inf}
        g[i] = [2]int{inf, inf}
    }
    f[0][0] = nums[0]

    ans := f[0][0]
    for i := 1; i < n; i++ {
        f[i][0] = max(0, f[i - 1][1]) + nums[i]
        f[i][1] = f[i - 1][0] - nums[i]
        g[i][0] = g[i - 1][1] + nums[i]
        g[i][1] = g[i - 1][0] - nums[i]
        if i >= 2 {
            g[i][0] = max(g[i][0], f[i - 2][1] + nums[i])
            g[i][1] = max(g[i][1], f[i - 2][0] - nums[i])
        }
        ans = max(ans, f[i][0], f[i][1], g[i][0], g[i][1])
    }

    return int64(ans)
}