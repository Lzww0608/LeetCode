func minOperations(nums []int, p int) int {
    n := len(nums)
    i, j, k := 0, 0, n
    for i < k {
        if nums[i] > p {
            k--
            nums[i], nums[k] = nums[k], nums[i]
        } else if nums[i] < p {
            nums[i], nums[j] = nums[j], nums[i]
            i++
            j++
        } else {
            i++
        }
    }

    return j
}