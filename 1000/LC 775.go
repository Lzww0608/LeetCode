
func isIdealPermutation(nums []int) bool {
    n := len(nums)
    minVal := nums[n-1]
    for i := n - 3; i >= 0; i-- {
        if nums[i] > minVal {
            return false
        }
        minVal = min(minVal, nums[i+1])
    }
    return true

}
