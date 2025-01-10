func applyOperations(nums []int) []int {
    n := len(nums)
    for i := 0; i < n - 1; i++ {
        if nums[i] == nums[i+1] {
            nums[i], nums[i+1] = nums[i] * 2, 0
        }
    }

    for l, r := 0, 0; r < n; r++ {
        if nums[r] != 0 && l < n {
            nums[l], nums[r] = nums[r], nums[l]
        }

        if l < n && nums[l] != 0 {
            l++
        }
    }

    return nums
}