func maximumSum(arr []int) int {

    a, b := math.MinInt32 / 2, math.MinInt32 / 2
    ans := math.MinInt32
    for _, x := range arr {
        a, b = max(a, 0) + x, max(b + x, a)
        ans = max(ans, a, b)
    }

    return ans
}