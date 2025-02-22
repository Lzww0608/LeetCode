func maximumTotalCost(nums []int) int64 {
    n := len(nums)
    f := make([][2]int, n + 1)
    f[1][0], f[1][1] = nums[0], math.MinInt
    for i := 1; i < n; i++ {
        x := nums[i]
        f[i+1][0] = max(f[i][0], f[i][1]) + x
        f[i+1][1] = f[i][0] - x
    }

    return int64(max(f[n][0], f[n][1]))
}