func minInitialStrength(monsters []int, boosts [][]int) int64 {
    l, r := -1, int(1e14)
    n := len(monsters)
    f := make([]int, n + 1)
    for _, v := range boosts {
        l, r, x := v[0], v[1], v[2]
        f[l] += x 
        f[r + 1] -= x
    }
    for i := 1; i <= n; i++ {
        f[i] += f[i - 1]
    }

    check := func(mid int) bool {
        for i, x := range monsters {
            cur := mid + f[i]
            if cur < x {
                return false
            }
            mid = max(0, mid - x)
        }

        return true
    }

    for l + 1 < r {
        mid := l + ((r - l) >> 1)
        if check(mid) {
            r = mid
        } else {
            l = mid
        }
    }

    return int64(r)
}