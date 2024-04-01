func minimumLevels(a []int) int {
    n := len(a)
    pre := make([]int, n + 1)
    for i := range a {
        if a[i] == 0 {
            pre[i+1] = pre[i] - 1
        } else {
            pre[i+1] = pre[i] + 1
        }
    }
    
    for i := 0; i < n - 1; i++ {
        if pre[i+1] > pre[n] - pre[i+1] {
            return i + 1
        }
    }
    return -1
}