func minimumSum(nums []int) int {
    n := len(nums)
    suf := make([]int, n + 1)
    suf[n-1] = nums[n-1]
    for i := n - 2; i >= 0; i-- {
        suf[i] = min(nums[i], suf[i+1])
    }

    pre, ans := nums[0], math.MaxInt32
    for i := 1; i < n - 1; i++ {
        if pre < nums[i] && suf[i+1] < nums[i] {
            ans = min(ans, pre + nums[i] + suf[i+1])
        }
        pre = min(pre, nums[i])
    }

    if ans == math.MaxInt32 {
        return -1
    }
    return ans
}