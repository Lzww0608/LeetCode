func rob(nums []int) int {
    n := len(nums)
    dp := make([]int, n + 2)
    dp[2] = nums[0]
    ans := dp[2]
    for i := 1; i < n; i++ {
        dp[i+2] = max(dp[i-1] + nums[i], dp[i] + nums[i])
        ans = max(ans, dp[i+2])
    }
    return ans
}