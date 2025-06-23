func minIncrementOperations(nums []int, k int) int64 {
    n := len(nums)
    f := make([]int, n + 1)
    f[1], f[2] = max(0, k - nums[0]), max(0, k - nums[1])
    
    for i := 2; i < n; i++ {
        f[i+1] = max(0, k - nums[i]) + min(f[i-1], f[i], f[i-2])
    }

    return int64(min(f[n-1], f[n-2], f[n]))
}