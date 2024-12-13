func getFinalState(nums []int, k int, multiplier int) []int {
    for i := 0; i < k; i++ {
        mn := slices.Min(nums)
        for j := 0; j < len(nums); j++ {
            if nums[j] == mn {
                nums[j] *= multiplier
                break
            }
        }
    }

    return nums
}