func findValidElements(nums []int) (ans []int) {
    n := len(nums)
    suf := make([]int, n + 1)
    for i := n - 1; i >= 0; i-- {
        suf[i] = max(suf[i + 1], nums[i])
    }

    pre := 0
    for i, x := range nums {
        if x > pre || x > suf[i + 1] {
            ans = append(ans, x)
        }
        pre = max(pre, x)
    }

    return 
}