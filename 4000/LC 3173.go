func orArray(nums []int) []int {
    for i := 0; i < len(nums) - 1; i++ {
        nums[i] |= nums[i + 1]
    }

    return nums[: len(nums) - 1]
}