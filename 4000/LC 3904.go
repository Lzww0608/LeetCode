func firstStableIndex(nums []int, k int) int {
    n := len(nums)
    pre := make([]int, n + 1)
    for i, x := range nums {
        pre[i + 1] = max(pre[i], x)
    }

    ans := n + 1 
    suf := nums[n - 1]
    for i := n - 1; i >= 0; i-- {
        suf = min(suf, nums[i])
        x := pre[i + 1] - suf 
        if x <= k {
            ans = i
        }
    }

    if ans == n + 1 {
        return -1
    }
    return ans
}