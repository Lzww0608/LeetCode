func massage(nums []int) int {
    n := len(nums)
    if n == 0 {
        return 0
    }
    f := make([]int, n + 1)
    f[1] = nums[0]
    for i := 1; i < n; i++ {
        f[i+1] = max(f[i], f[i-1] + nums[i])
    }

    return f[n]
}