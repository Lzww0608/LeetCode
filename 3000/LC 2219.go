func maximumSumScore(nums []int) int64 {
    ans := math.MinInt 
    n := len(nums)
    pre := make([]int, n + 1)
    for i, x := range nums {
        pre[i + 1] = pre[i] + x
    }

    suf := 0
    for i := n - 1; i >= 0; i-- {
        suf += nums[i]
        ans = max(ans, suf, pre[i + 1])
    }

    return int64(ans)
}