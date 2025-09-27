func minMaxGame(nums []int) int {
    for len(nums) > 1 {
        n := len(nums)
        cur := make([]int, n / 2)
        for i := range cur {
            if i & 1 == 0 {
                cur[i] = min(nums[i * 2], nums[i * 2 + 1])
            } else {
                cur[i] = max(nums[i * 2], nums[i * 2 + 1])
            }
        }
        nums = cur
    }

    return nums[0]
}