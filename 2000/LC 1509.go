func minDifference(nums []int) int {
    n := len(nums)
    if n <= 4 {
        return 0
    }
    sort.Ints(nums)
    ans := math.MaxInt32
    ans = min(nums[n-4] - nums[0], nums[n-1] - nums[3], nums[n-2] - nums[2], nums[n-3] - nums[1])

    return ans
}