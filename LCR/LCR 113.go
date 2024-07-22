func findOrder(numCourses int, prerequisites [][]int) (ans []int) {
    g := make([][]int, numCourses)
    deg := make([]int, numCourses)

    for _, e := range prerequisites {
        a, b := e[0], e[1]
        g[b] = append(g[b], a)
        deg[a]++
    }

    q := []int{}
    for i, x := range deg {
        if x == 0 {
            q = append(q, i)
        }
    }

    for len(q) > 0 {
        node := q[0]
        q = q[1:]
        ans = append(ans, node)
        for _, x := range g[node] {
            deg[x]--
            if deg[x] == 0 {
                q = append(q, x)
            }
        }
    }

    if len(ans) < numCourses {
        return nil
    }

    return
} 