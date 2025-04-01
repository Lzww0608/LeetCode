func splitArray(nums []int) bool {
    n := len(nums)
    if n < 7 {
        return false
    }

    pre := make([]int, n + 1)
    for i := 0; i < n; i++ {
        pre[i+1] = pre[i] + nums[i]
    }

    for j := 3; j + 3 < n; j++ {
        vis := make(map[int]bool)
        for i := 1; i + 1 < j; i++ {
            if pre[i] * 2 == pre[j] - nums[i] {
                vis[pre[i]] = true
            }
        }

        for k := j + 2; k + 1 < n; k++ {
            if vis[pre[k] - pre[j+1]] && pre[n] - pre[k+1] == pre[k] - pre[j+1] {
                return true
            }
        }
    }

    return false
}