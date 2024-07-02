func minimumTime(n int, relations [][]int, time []int) int {
    g := make([][]int, n)
    deg := make([]int, n)
    for _, e := range relations {
        a, b := e[0] - 1, e[1] - 1
        g[a] = append(g[a], b)
        deg[b]++
    }

    f := make([]int, n)
    q := []int{}
    for i, x := range deg {
        if x == 0 {
            q = append(q, i)
        }
    }

    for len(q) > 0 {
        x := q[0]
        f[x] += time[x]
        q = q[1:]
        for _, y := range g[x] {
            f[y] = max(f[y], f[x])
            if deg[y]--; deg[y] == 0 {
                q = append(q, y)
            }
        }
    }

    return slices.Max(f)
}