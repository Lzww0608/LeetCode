var parent map[string]string = map[string]string{}

func find(x string) string {
    if x != parent[x] {
        parent[x] = find(parent[x])
    }
    return parent[x]
}

func merge(x, y string) {
    rx, ry := find(x), find(y)
    if rx != ry {
        parent[ry] = rx
    }
}

func numSimilarGroups(strs []string) int {
    n := len(strs)
    check := func(x, y int) bool {
        a, b := strs[x], strs[y]
        cnt := 0
        for k := range a {
            if a[k] != b[k] {
                cnt++
                if cnt > 2 {
                    return false
                }
            }
        }
        return true
    }
    for _, s := range strs {
        parent[s] = s
    }
    for i, s := range strs {
        for j := i + 1; j < n; j++ {
            if check(i, j) {
                merge(s, strs[j])
            }
        }
    }
    ans := 0
    m := map[string]bool{}
    for _, s := range strs {
        p := find(s)
        if _, ok := m[p]; !ok {
            ans++
            m[p] = true
        }
    }
    return ans
}