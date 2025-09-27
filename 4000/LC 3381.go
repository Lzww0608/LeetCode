func maxSubarraySum(nums []int, k int) int64 {
    f := make([]int, k)
    for i := range f {
        f[i] = math.MaxInt / 2
    }

    pre, ans := 0, math.MinInt
    f[0] = 0
    for i, x := range nums {
        pre += x 
        ans = max(ans, pre - f[(i + 1) % k])
        f[(i + 1) % k] = min(f[(i + 1) % k], pre)
    }

    return int64(ans)
}