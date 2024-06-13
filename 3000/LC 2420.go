func goodIndices(nums []int, k int) (ans []int) {
    n := len(nums)
    suf := make([]int, n)


    suf[n-1] = 1
    for i := n - 2; i >= 0; i-- {
        if nums[i] <= nums[i+1] {
            suf[i] = suf[i+1] + 1
        } else {
            suf[i] = 1
        }
    }

    for i, pre := 1, 1; i < n - k; i++ {
        if pre >= k && suf[i+1] >= k {
            ans = append(ans, i)
        }

        if nums[i] <= nums[i-1] {
            pre++
        } else {
            pre = 1
        }
    }
    return
}
