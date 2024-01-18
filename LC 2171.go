func minimumRemoval(beans []int) int64 {
    sort.Ints(beans)
    n := len(beans)
    pre := make([]int64, n+1)
    for i, x := range beans {
        pre[i+1] = pre[i] + int64(x)
    }
    var ans int64 = pre[n]
    for i := 0; i < n; i++ {
        ans = min(ans, pre[i] + pre[n] - pre[i] - int64(beans[i]) * int64(n - i))
    }
    return ans
}