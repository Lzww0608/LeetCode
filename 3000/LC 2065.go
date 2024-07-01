func maximalPathQuality(values []int, edges [][]int, maxTime int) (ans int) {
    type pair struct {
        d, w int
    }
    n := len(values)
    g := make([][]pair, n)
    for _, e := range edges {
        a, b, c := e[0], e[1], e[2]
        g[a] = append(g[a], pair{b, c})
        g[b] = append(g[b], pair{a, c})
    }

    m := map[int]bool{0: true}

    var dfs func(int, int, int)
    dfs = func(node, cur, time int) {
        if node == 0 {
            ans = max(ans, cur)
        }

        for _, v := range g[node] {
            a, b := v.d, v.w
            if time + b > maxTime {
                continue
            }

            if m[a] {
                dfs(a, cur, time + b)
            } else {
                m[a] = true
                dfs(a, cur + values[a], time + b)
                m[a] = false
            }
        }

    }
    dfs(0, values[0], 0)

    return
}