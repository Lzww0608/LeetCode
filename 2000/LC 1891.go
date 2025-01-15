func maxLength(ribbons []int, k int) (ans int) {
    l, r := 0, slices.Max(ribbons) + 1

    check := func(d int) bool {
        if d == 0 {
            return false
        }
        cnt := 0
        for _, x := range ribbons {
            cnt += x / d
        }
        return cnt >= k
    }

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