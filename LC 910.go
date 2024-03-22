func smallestRangeII(nums []int, k int) int {
    sort.Ints(nums)
    n := len(nums)
    ans := nums[n-1] - nums[0]

    for i := range nums[:n-1] {
        ans = min(ans, max(nums[i] + k, nums[n-1] - k) - min(nums[0] + k, nums[i+1] - k))
    }
    
    return ans
}