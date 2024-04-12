func canJump(nums []int) bool {
    n := len(nums)
    maxPos, end := 0, 0
    
    for i := 0; i < n - 1; i++ {
        maxPos = max(nums[i] + i, maxPos)
        if i == end {
            end = maxPos
        }
    }

    return end >= n - 1
}