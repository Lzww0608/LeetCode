func rampartDefensiveLine(a [][]int) int {
    l, r := -1, int(1e8)
    n := len(a)

    check := func(mid int) bool {
        rem := a[1][0] - a[0][1]
        for i := 1; i < n - 1; i++ {
            cur := a[i + 1][0] - a[i][1]
            if rem + cur < mid {
                return false
            } else {
                cur -= max(0, mid - rem)
            }

            rem = cur
        }

        return true
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