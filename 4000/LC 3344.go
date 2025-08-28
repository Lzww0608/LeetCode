func maxSizedArray(s int64) int {
    
    check := func(m int) bool {
        if m == 0 || m == 1 {
            return s >= 0
        }

        sum := m * (m - 1) / 2
        cur := 0
        for j := 0; j < m; j++ {
            for k := 0; k < m; k++ {
                cur += (j | k)
                if cur * sum > int(s) {
                    return false
                }
            }
        }

        return true
    }


    l, r := 0, 1500
    for l + 1 < r {
        mid := l + ((r - l) >> 1)
        if check(mid) {
            l = mid
        } else {
            r = mid
        }
    }

    return l
}