func canFinish(numCourses int, prerequisites [][]int) bool {
    in := make([]int, numCourses)
    g := make([][]int, numCourses)
    for _, v := range prerequisites {
        a, b := v[0], v[1]
        in[a]++
        g[b] = append(g[b], a)
    }

    q := make([]int, 0, numCourses)
    for i, x := range in {
        if x == 0 {
            q = append(q, i)
        }
    }

    for len(q) > 0 {
        u := q[0]
        q = q[1:]
        for _, v := range g[u] {
            in[v]--
            if in[v] == 0 {
                q = append(q, v)
            }
        }
    }

    return cap(q) == 0
}