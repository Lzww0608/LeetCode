func minOperations(nums []int) int {
    n := len(nums)
    f := []int{}
    for i := n - 1; i >= 0; i-- {
        if len(f) == 0 || nums[i] >= f[len(f)-1] {
            f = append(f, nums[i])
        } else {
            pos := sort.SearchInts(f, nums[i] + 1)
            f[pos] = nums[i]
        }
    }

    return len(f)
}