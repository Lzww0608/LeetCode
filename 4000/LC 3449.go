func maxScore(points []int, m int) int64 {
    n := len(points)

    check := func(mid int) bool {
        rem, pre := m, 0 

        for i, x := range points {
            k := (mid - 1) / x + 1 - pre 
            if i == n - 1 && k <= 0 {
                break
            }
            k = max(k, 1)
            if rem -= k * 2 - 1; rem < 0 {
                return false
            }
            pre = k - 1
        }

        return true
    }

    l, r := 0, (m + 1) / 2 * slices.Min(points) + 1
    for l + 1 < r {
        mid := l + ((r - l) >> 1)
        if check(mid) {
            l = mid
        } else {
            r = mid 
        }
    }

    return int64(l)
}