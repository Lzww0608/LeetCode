func maximumProduct(nums []int, m int) int64 {
    n := len(nums)
    type pair struct {
        mx, mn int 
    }
    suf := make([]pair, n + 1)
    suf[n] = pair{mx: math.MinInt32, mn: math.MaxInt32}
    for i := n - 1; i >= 0; i-- {
        x := nums[i]
        suf[i] = suf[i + 1]
        suf[i].mx = max(suf[i].mx, x)
        suf[i].mn = min(suf[i].mn, x)
    }

    ans := math.MinInt
    for i := 0; i + m <= n; i++ {
        ans = max(ans, nums[i] * suf[i + m - 1].mx, nums[i] * suf[i + m - 1].mn)
    }

    return int64(ans)
}