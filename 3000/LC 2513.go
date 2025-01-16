func minimizeSet(divisor1 int, divisor2 int, uniqueCnt1 int, uniqueCnt2 int) int {

    g := divisor1 * divisor2 / gcd(divisor1, divisor2)

    check := func(k int) bool {
        sum := k - k / divisor1 - k / divisor2 + k / g // for 1 and 2
        a := k / divisor1 - k / g // only for 2
        b := k / divisor2 - k / g // only for 1
        return sum >= max(uniqueCnt1 - b, 0) + max(uniqueCnt2 - a, 0)
    }

    l, r := 0, int(2e9)
    for l + 1 < r {
        mid := l + ((r - l) >> 1)
        if check(mid) {
            r = mid
        } else {
            l = mid
        }
    }

    return r
}

func gcd(x, y int) int {
    for y != 0 {
        x, y = y, x % y
    }
    return x
}