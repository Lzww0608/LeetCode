func minimumSemesters(n int, relations [][]int) (ans int) {
    g := make([][]int, n)
    in := make([]int, n)
    for _, v := range relations {
        a, b := v[0] - 1, v[1] - 1
        g[a] = append(g[a], b)
        in[b]++
    }

    q := make([]int, 0, n)
    for i, x := range in {
        if x == 0 {
            q = append(q, i)
        }
    }

    for len(q) > 0 {
        ans++
        for sz := len(q); sz > 0; sz-- {
            u := q[0]
            q = q[1:]
            for _, v := range g[u] {
                if in[v]--; in[v] == 0 {
                    q = append(q, v)
                }
            }
        }
    }

    if cap(q) != 0 {
        return -1
    }

    return 
}