func partitionDisjoint(nums []int) int {
    n := len(nums)
    suf := make([]int, n + 1)
    suf[n] = math.MaxInt32
    for i := n - 1; i >= 0; i-- {
        suf[i] = min(suf[i + 1], nums[i])
    }

    pre := math.MinInt32
    for i, x := range nums {
        pre = max(pre, x)
        if pre <= suf[i + 1] {
            return i + 1
        }
    }

    return -1
}