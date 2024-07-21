func numSimilarGroups(strs []string) (ans int) {
    n := len(strs)
    f := make([]int, n)
    r := make([]int, n)
    for i := range f {
        f[i] = i
        r[i] = 1
    }

    var find func(int) int
    find = func(x int) int {
        if f[x] != x {
            f[x] = find(f[x])
        }
        return f[x]
    }

    merge := func(x, y int) {
        rx, ry := find(x), find(y)
        if rx != ry {   
            if r[rx] < r[ry] {
                rx, ry = ry, rx
            }
            f[ry] = rx
            if r[rx] == r[ry] {
                r[rx]++
            }
        }
    }

    check := func(a, b string) bool {
        cnt := 0
        for i := range a {
            if a[i] != b[i] {
                cnt++
                if cnt > 2 {
                    return false
                }
            }
        }

        return true
    }

    for i := range strs {
        for j := i + 1; j < n; j++ {
            if check(strs[i], strs[j]) {
                merge(i, j)
            }
        }
    }

    for i := range strs {
        if f[i] == i {
            ans++
        }
    }

    return 
}