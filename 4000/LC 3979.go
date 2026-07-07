func maxValidPairSum(nums []int, k int) (ans int) {
    n := len(nums)
    suf := make([]int, n - k + 1)
    for i := n - 1; i >= k; i-- {
        j := i - k
        suf[j] = max(suf[j + 1], nums[i])
    }

    for i := range n - k {
        ans = max(ans, nums[i] + suf[i])
    }

    return 
}