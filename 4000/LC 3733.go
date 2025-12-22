func minimumTime(d []int, r []int) int64 {
    t := lcm(r[0], r[1])

    check := func(mid int) bool {
        return d[0] <= mid - mid / r[0] && d[1] <= mid - mid / r[1] && d[0] + d[1] <= mid - mid / t
    }

    l, rt := d[0] + d[1] - 1, (d[0] + d[1]) * 2 
    for l + 1 < rt {
        mid := l + ((rt - l) >> 1)
        if check(mid) {
            rt = mid
        } else {
            l = mid
        }
    }

    return int64(rt)
}

func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a % b
    }

    return a
}

func lcm(a, b int) int {
    return a * b / gcd(a, b)
}