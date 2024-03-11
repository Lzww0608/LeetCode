func maximumCandies(a []int, k int64) int {
    l, r := 1, int(1e7) + 1

    cal := func(per int) bool {
        if per == 0 {
            return true
        }
        var res int64
        for _, x := range a {
            res += int64(x / per)
        }
        if res >= k {
            return true
        }
        return false
    }

    for l < r {
        mid := l + ((r - l) >> 1)
        if cal(mid) {
            l = mid + 1
        } else {
            r = mid
        }
    }
 
    return l - 1
}