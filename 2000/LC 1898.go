func maximumRemovals(s string, p string, removable []int) (ans int) {
    n := len(removable)
    l, r := 0, n + 1

    check := func(mid int) bool {
        m := map[int]bool{}
        for _, x := range removable[:mid] {
            m[x] = true
        }
        i, j := 0, 0
        for i < len(s) && j < len(p) {
            if !m[i] && s[i] == p[j] {
                j++
            }
            i++
        }

        return j >= len(p)
    }


    for l < r {
        mid := l + ((r - l) >> 1)
        if check(mid) {
            ans = mid
            l = mid + 1
        } else {
            r = mid
        }
    }

    return
}