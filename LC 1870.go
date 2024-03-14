func minSpeedOnTime(dist []int, hour float64) int {
    l, r := 1, int(1e9)
    n := len(dist)

    check := func(sp int) bool {
        var res float64
        for i := 0; i < n - 1; i++{
            res += float64((dist[i] + sp - 1) / sp)
        }
        res += float64(dist[n-1]) / float64(sp)
        
        return res <= hour
    }

    if !check(r) {
        return -1
    } 

    for l <= r {
        mid := l + ((r - l) >> 1)
        
        if check(mid) {
            r = mid - 1
        } else {
            l = mid + 1
        }
    }
    return l
}