func minTime(s string, order []int, k int) int {
    n := len(s)
    if k > n * (n + 1) / 2 {
        return -1
    }
    a := make([]int, n)

    check := func(mid int) bool {
        for i := range mid + 1 {
            a[order[i]] = mid + 1
        }

        cur := 0
        last := -1
        for i := range n {
            if a[i] == mid + 1 {
                last = i
            }

            cur += last + 1
            if cur >= k {
                return true
            }
        }

        return false
    }

    l, r := -1, n
    for l + 1 < r {
        mid := l + ((r - l) >> 1)
        if check(mid) {
            r = mid
        } else {
            l = mid
        }
    }

    if r == n {
        return -1
    }

    return r
}