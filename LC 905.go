func sortArrayByParity(nums []int) []int {
    n := len(nums)
    l, r := 0, n - 1
    for i := 0; i <= r; i++ {
        if nums[i] & 1 == 0 {
            nums[l] = nums[i]
            l++
        } else {
            nums[i], nums[r] = nums[r], nums[i]
            r--
            i--
        }
    }
    return nums
}