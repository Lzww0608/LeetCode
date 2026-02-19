func rob(nums []int, colors []int) int64 {
    n := len(nums)
    f := make([][2]int, n)
    f[0][1] = nums[0]

    for i := 1; i < n; i++ {
        f[i][0] = max(f[i - 1][0], f[i - 1][1])
        if colors[i] == colors[i - 1] {
            f[i][1] = max(f[i - 1][1], f[i - 1][0] + nums[i]) 
        } else {
            f[i][1] = max(f[i - 1][1], f[i - 1][0]) + nums[i]
        }
    }

    return int64(max(f[n - 1][0], f[n - 1][1]))
}