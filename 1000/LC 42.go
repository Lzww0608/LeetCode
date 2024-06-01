func trap(height []int) int {
    n := len(height)
    pre, suf := make([]int, n), make([]int, n)
    pre[0], suf[n-1] = height[0], height[n-1]

    for i := 1; i < n; i++ {
        pre[i] = max(pre[i-1], height[i])
    }

    for i := n - 2; i >= 0; i-- {
        suf[i] = max(suf[i+1], height[i])
    }

    ans := 0
    for i := range height {
        ans += min(pre[i], suf[i]) - height[i]
    }
    return ans
}
