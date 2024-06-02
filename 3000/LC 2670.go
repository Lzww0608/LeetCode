func distinctDifferenceArray(nums []int) []int {
    m1, m2 := map[int]int{}, map[int]int{}
    n := len(nums)
    pre, suf := make([]int, n), make([]int, n+1)
    pre[0], suf[n-1] = 1, 1
    m1[nums[0]], m2[nums[n-1]] = 1, 1
    for i := 1; i < n; i++ {
        if _, ok := m1[nums[i]]; !ok {
            m1[nums[i]] = 1
            pre[i] = pre[i-1] + 1
        } else {
            pre[i] = pre[i-1]
        }
    }
    
    for i := n - 2; i >= 0; i-- {
        if _, ok := m2[nums[i]]; !ok {
            m2[nums[i]] = 1
            suf[i] = suf[i+1] + 1
        } else {
            suf[i] = suf[i+1]
        }
    }

    ans := make([]int, n)
    for i := range ans {
        ans[i] = pre[i] - suf[i+1]
    }
    return ans
}
