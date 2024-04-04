func sortArrayByParityII(nums []int) []int {
    for i, j := 0, 1; j < len(nums); j += 2 {
        for nums[j] & 1 == 0 {
            for nums[i] & 1 == 0 {
                i += 2
            }
            nums[i], nums[j] = nums[j], nums[i]
        }
    }

    return nums
}