func maxPossibleScore(start []int, d int) int {
    sort.Ints(start)

    check := func(mid int) bool {
        pre := start[0]
        for _, x := range start[1:] {
            if pre + mid > x + d {
                return false
            } else {
                pre = max(x, pre + mid)
            }
        }

        return true
    }

    l, r := -1, int(2e9) + 1
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