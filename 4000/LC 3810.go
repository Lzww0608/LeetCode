func minOperations(nums []int, target []int) int {
    vis := make(map[int]bool)
    for i := range nums {
        if nums[i] != target[i] {
            vis[nums[i]] = true
        }
    }

    return len(vis)
}