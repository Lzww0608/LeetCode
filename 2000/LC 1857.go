const N int = 26
func largestPathValue(colors string, edges [][]int) (ans int) {
    n := len(colors)
    f := make([][N]int, n)
    g := make([][]int, n)
    q := make([]int, 0, n)
    in := make([]int, n)
    for _, edge := range edges {
        u, v := edge[0], edge[1]
        if u == v {
            return -1
        }
        g[u] = append(g[u], v)
        in[v]++
    }

    for i := 0; i < n; i++ {
        if in[i] == 0 {
            q = append(q, i)
        }
    }

    for len(q) > 0 {
        u := q[0]
        q = q[1:]
        x := int(colors[u]-'a')
        f[u][x]++
        ans = max(ans, f[u][x])

        for _, v := range g[u] {
            for i, x := range f[u] {
                f[v][i] = max(f[v][i], x)
            }
            if in[v]--; in[v] == 0 {
                q = append(q, v)
            }
        }
    }

   
    if cap(q) != 0 {
        return -1
    }
    return
}