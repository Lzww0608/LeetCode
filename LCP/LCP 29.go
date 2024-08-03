func orchestraLayout(n int, x int, y int) int {
    t := min(x, y, n - x - 1, n - y - 1)

    //S = a1 * n + n * (n - 1) * d / 2
    m := t * (n - 1) - (t - 1) * t
    m <<= 2

    u, l := t, t
    d, r := n - t - 1, n - t - 1

    if x == u {
        m += y - l + 1
    } else if y == r {
        m += (r - l) + (x - u + 1)
    } else if x == d {
        m += (r - l) * 2 + (r - y + 1)
    } else if y == l {
        m += (r - l) * 3 + (d - x + 1)
    }

    if m % 9 == 0 {
        return 9
    }

    return m % 9
}