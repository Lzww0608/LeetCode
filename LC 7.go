func reverse(x int) int {
    ans, t := 0, 0
    for x != 0 {
        if ans < math.MinInt32 / 10 || ans > math.MaxInt32 / 10 {
            return 0
        }
        t, x = x % 10, x / 10
        ans = ans * 10 + t
    }
    return ans
}