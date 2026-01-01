func maximumScore(nums []int) int64 {
    n := len(nums)
    suf := make([]int, n)
    suf[n - 1] = nums[n - 1]
    for i := n - 2; i >= 0; i-- {
        suf[i] = min(nums[i], suf[i + 1])
    }

    pre := 0
    ans := math.MinInt
    for i := range n - 1 {
        pre += nums[i]
        ans = max(ans, pre - suf[i + 1])
    }

    return int64(ans)
}