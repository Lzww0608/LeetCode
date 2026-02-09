func finalElement(nums []int) int {
    n := len(nums)
    return max(nums[0], nums[n - 1])
}