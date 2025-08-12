func minDays(bloomDay []int, m int, k int) int {
    n := len(bloomDay)
    if m * k > n {
        return -1
    }

    check := func(mx int) bool {
        cnt := 0
        cur := 0
        for _, x := range bloomDay {
            if x > mx {
                cur = 0
            } else {
                cur++
                if cur == k {
                    cur = 0
                    cnt++
                }
            }
        }

        return cnt >= m
    }

    l, r := 0, slices.Max(bloomDay) + 1
    for l < r {
        mid := l + ((r - l) >> 1)
        if check(mid) {
            r = mid
        } else {
            l = mid + 1
        }
    }

    return r
}