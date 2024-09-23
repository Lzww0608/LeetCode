func maxScoreSightseeingPair(values []int) int {
    n := len(values)
    tmp := make([]int, n)
    for i, x := range values {
        tmp[i] = x - i
        values[i] += i
    }

    ans, cur := math.MinInt32, math.MinInt32
    for i := 1; i < n; i++ {
        cur = max(cur, values[i-1])
        ans = max(ans, cur + tmp[i])
    }

    return ans
}