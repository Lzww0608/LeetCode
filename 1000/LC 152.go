func maxProduct(nums []int) int {
    mx, mn := nums[0], nums[0]
    ans := nums[0]
    for _, x := range nums[1:] {
        a, b := mx * x, mn * x
        mx = max(x, max(a, b))
        mn = min(x, min(a, b))
        ans = max(ans, mx)
    }
    return ans
}
