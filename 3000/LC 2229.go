func isConsecutive(nums []int) bool {
    cnt := make(map[int]bool)
    mn, mx := nums[0], nums[0]
    for _, x := range nums {
        if cnt[x] {
            return false
        }
        cnt[x] = true
        mn = min(mn, x)
        mx = max(mx, x)
    }

    return mx - mn + 1 == len(nums)
}