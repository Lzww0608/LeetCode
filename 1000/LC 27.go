func removeElement(nums []int, val int) int {
    n := len(nums)
    i, j := 0, 0
    for j < n {
        if nums[j] != val {
            nums[i] = nums[j]
            i++
        }
        j++
    }
    return i
}
