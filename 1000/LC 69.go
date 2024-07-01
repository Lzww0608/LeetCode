func mySqrt(x int) int {
    t := int64(x)
    l, r := int64(-1), t + 1

    for l + 1 < r {
        mid := l + ((r - l) >> 1)
        if mid * mid < t {
            l = mid
        } else if mid * mid > t {
            r = mid
        } else {
            return int(mid)
        }
    }
    return int(l)
}
