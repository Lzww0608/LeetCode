func maxSumOfThreeSubarrays(nums []int, k int) (ans []int) {
    n := len(nums)
    sum := make([]int, n)
    pre := make([]int, n)
    suf := make([]int, n)
    for i := 0; i < k; i++ {
        sum[0] += nums[i]
    }
    for i := k; i < n; i++ {
        sum[i-k+1] = sum[i-k] + nums[i] - nums[i-k]
    }

    pre[0] = 0
    for i := 1; i < n - k + 1; i++ {
        if sum[i] <= sum[pre[i-1]] {
            pre[i] = pre[i-1]
        } else {
            pre[i] = i
        }
    }

    suf[n-1] = n - 1
    for i := n - 2; i > k; i-- {
        if sum[i] >= sum[suf[i+1]] {
            suf[i] = i
        } else {
            suf[i] = suf[i+1]
        }
    }

    s := 0
    for i := k; i < n - k; i++ {
        if s < sum[i] + sum[pre[i-k]] + sum[suf[i+k]] {
            s = sum[i] + sum[pre[i-k]] + sum[suf[i+k]]
            ans = []int{pre[i-k], i, suf[i+k]}
        }
    }

    return
}
