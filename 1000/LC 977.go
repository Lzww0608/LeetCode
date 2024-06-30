func sortedSquares(nums []int) []int {
    i, j := 0, len(nums) - 1
    ans := make([]int, len(nums))
    for k := len(nums) - 1; k >= 0; k-- {
        if nums[i] * nums[i] >= nums[j] * nums[j] {
            ans[k] = nums[i] * nums[i]
            i++
        } else {
            ans[k] = nums[j] * nums[j]
            j--
        }
    }
    return ans
}
