func maximumTotalSum(a []int) int64 {
    ans := 0
    sort.Ints(a)
    n := len(a)
    cur := a[n-1]
    for i := n - 1; i >= 0; i-- {
        cur = min(cur, a[i])
        if cur == 0 {
            return -1
        }
        ans += cur
        cur--
    }

    return int64(ans)
}