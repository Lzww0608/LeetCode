func minOperations(nums []int, k int) int {
    n := len(nums)
    mask := 0
    for i := n - 1; i >= 0; i-- {
        mask |= 1 << (nums[i] - 1)
        if mask & ((1 << k) - 1) == (1 << k) - 1 {
            return n - i
        }
    }

    return -1
}