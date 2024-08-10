func minTime(time []int, m int) (ans int) {
    if m >= len(time) {
        return 0
    }

    l, r := 0, 0
    for i := m; i < len(time); i++ {
        r += time[i]
    }

    check := func(limit int) bool {
        cur, cnt, maxInDay := 0, 1, 0

        for _, x := range time {
            maxInDay = max(maxInDay, x)  
            if cur + x - maxInDay > limit {
                cnt++
                cur = 0
                maxInDay = x
            }
            cur += x
        }

        return cnt <= m
    }


    for l < r {
        mid := l + ((r - l) >> 1)
            if check(mid) {
                ans = mid
                r = mid
            } else {
                l = mid + 1
            }
    }

    return 
}