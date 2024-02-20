func isMonotonic(nums []int) bool {
    a, b := true, true
    for i := 1; i < len(nums); i++ {
        if nums[i] < nums[i-1] {
            a = false
        } else if nums[i] > nums[i-1] {
            b = false
        }
    }
    return a || b
}