func maxValue(nums []int) []int {
    n := len(nums)
    pre := make([]int, n + 1)
    for i, x := range nums {
        pre[i + 1] = max(pre[i], x)
    }

    suf := math.MaxInt
    ans := make([]int, n)

    for i := n - 1; i >= 0; i-- {
        if suf >= pre[i + 1] {
            ans[i] = pre[i + 1]
        } else {
            ans[i] = ans[i + 1]
        }
        suf = min(suf, nums[i])
    }

    return ans
}