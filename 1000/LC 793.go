func preimageSizeFZF(k int) int {
    l, r := int64(k) * 4, int64(k) * 5
    for l < r {
        mid := l + ((r - l) >> 1)
        if zero(mid) < k {
            l = mid + 1
        } else {
            r = mid
        }
    }
    if zero(l) == k {
        return 5
    }
    return 0
}

func zero(n int64) (res int) {
    for n > 0 {
        res += int(n / 5)
        n /= 5
    }

    return
}