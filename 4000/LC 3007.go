func findMaximumNumber(k int64, x int) int64 {

    check := func(mid int) bool {
        i := x - 1
        res := 0
        for n := mid >> i; n > 0; n, i = n >> x, i + x {
            res += (n >> 1) << i
            if n & 1 == 1 {
                mask := (1 << i) - 1
                res += (mid & mask) + 1
            }
        }

        return res <= int(k)
    }   

    l, r := 0, int(k + 1) << x
    for l < r {
        mid := l + ((r - l) >> 1)
        if check(mid) {
            l = mid + 1
        } else {
            r = mid
        }
    }

    return int64(l-1)
}