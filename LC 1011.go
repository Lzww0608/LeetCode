func shipWithinDays(weights []int, days int) int {
    l, r := 0, 0
    for _, x := range weights {
        l = max(l, x)
        r += x
    }

    cal := func(cap int) bool {
        t, res := 0, 1
        for _, x := range weights {
            if t + x > cap {
                t = 0
                res++
            } 
            t += x
        }

        return res <= days
    }

    for l < r {
        mid := l + ((r - l) >> 1)
        
        if cal(mid) {
            r = mid
        } else {
            l = mid + 1
        }
    }

    return l
}