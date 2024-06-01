func waysToBuyPensPencils(total int, cost1 int, cost2 int) int64 {
    n := total / cost1
    return euclid(cost1, total % cost1, cost2, n) + int64(n + 1)
}

func euclid(a, b, c, n int) int64 {
    if n == 0 || a == 0 {
        return int64(b / c)
    }

    if a >= c || b >= c {
        return int64(n * (n + 1) / 2 * (a / c)) + int64((n + 1) * (b / c)) + euclid(a % c, b % c, c, n)
    }

    m := (a * n + b) / c
    return int64(n * m) - euclid(c, c - b - 1, a, m - 1) 
}
