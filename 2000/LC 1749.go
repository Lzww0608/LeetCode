func maxAbsoluteSum(nums []int) int {
    mn, mx := 0, 0
    ans := 0
    for _, x := range nums {
        mn += x
        mx += x
        mn = min(mn, x)
        mx = max(mx, x)
        ans = max(ans, abs(mx), abs(mn))
    }
    return ans
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
