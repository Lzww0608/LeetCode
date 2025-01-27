func maximumOr(nums []int, k int) int64 {
    ans := 0
    n := len(nums)
    pre := make([]int, n + 1)
    suf := make([]int, n + 1)
    for i, x := range nums {
        y := nums[n-i-1]
        pre[i+1] = pre[i] | x
        suf[n-i-1] = suf[n-i] | y
    }

    for i, x := range nums {
        ans = max(ans, (x << k) | suf[i+1] | pre[i])
    }

    return int64(ans)
}