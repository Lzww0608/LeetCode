func peakValleyDifference(nums []int) (ans int) {
    n := len(nums)
    for i := 1; i < n - 1; i++ {
        if nums[i] > nums[i - 1] && nums[i] > nums[i + 1] {
            ans += nums[i]
        } else if nums[i] < nums[i - 1] && nums[i] < nums[i + 1] {
            ans -= nums[i]
        }
    }

    return
}