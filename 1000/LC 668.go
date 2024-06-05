func findKthNumber(m int, n int, k int) int {

    check := func(x int) bool {
        cnt := 0
        r, c := 1, n
        for r <= m {
            if r * c <= x {
                cnt += c
                if cnt >= k {
                    return false
                }
                r++
            } else {
                c--
            }
        }

        return true
    }

    l, r := 1, m * n + 1
    for l < r {
        mid := l + ((r - l) >> 1)
        if check(mid) {
            l = mid + 1
        } else {
            r = mid
        }
    }
    
    return l
}