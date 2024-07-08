func maxSubArray(nums []int) int {
    ans := math.MinInt
    cur := 0
    for _, x := range nums {
        cur = max(cur + x, x)
        ans = max(ans, cur)
    }

    return ans
}