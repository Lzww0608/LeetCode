func findMaxAverage(nums []int, k int) float64 {
    sum, ans := 0, 0
    for i := 0; i < k; i++ {
        sum += nums[i]
    }
    ans = sum
    for i := k; i < len(nums); i++ {
        sum = sum - nums[i-k] + nums[i]
        ans = max(ans, sum)
    }
    return float64(ans) / float64(k)
}
