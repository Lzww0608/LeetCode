func temperatureTrend(a []int, b []int) (ans int) {
    n := len(a)
    cnt := 0
    for i := 1; i < n; i++ {
        if cmp.Compare(a[i-1], a[i]) == cmp.Compare(b[i-1], b[i]) {
            cnt++
        } else {
            cnt = 0
        }
        ans = max(ans, cnt)
    }
    
    return 
}