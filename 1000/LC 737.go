func areSentencesSimilarTwo(s []string, t []string, similarPairs [][]string) bool {
    if len(s) != len(t) {
        return false
    }

    m := map[string]int{}
    id := 0
    for _, v := range similarPairs {
        a, b := v[0], v[1]
        if _, ok := m[a]; !ok {
            m[a] = id
            id++
        }

        if _, ok := m[b]; !ok {
            m[b] = id
            id++
        }
    }

    fa := make([]int, id)
    for i := range fa {
        fa[i] = i
    }

    var find func(x int) int 
    find = func(x int) int {
        if fa[x] != x {
            fa[x] = find(fa[x])
        }

        return fa[x]
    }

    merge := func(x, y int) {
        rx, ry := find(x), find(y)
        if rx != ry {
            fa[ry] = rx
        }
    }

    for _, v := range similarPairs {
        a, b := v[0], v[1]
        merge(m[a], m[b])
    }

    for i := range s {
        if s[i] == t[i] {
            continue
        }

        if _, ok := m[s[i]]; !ok {
            return false
        }
        if _, ok := m[t[i]]; !ok {
            return false
        }

        if find(m[s[i]]) != find(m[t[i]]) {
            return false
        }
    }

    return true
}