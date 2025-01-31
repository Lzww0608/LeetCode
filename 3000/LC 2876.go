func countVisitedNodes(edges []int) []int {
    n := len(edges)
    ans := make([]int, n)
    in := make([]int, n)
    for _, x := range edges {
        in[x]++
    }

    q := []int{}
    g := make([][]int, n)
    for i, x := range in {
        if x == 0 {
            q = append(q, i)
        }
    }

    for len(q) > 0 {
        u := q[0]
        q = q[1:]
        v := edges[u]
        g[u] = append(g[u], v)
        if in[v]--; in[v] == 0 {
            q = append(q, v)
        }
    }

    maxDepth := make([]int, n)
    var dfs func(int) int
    dfs = func(x int) int {
        if maxDepth[x] != 0 {
            return maxDepth[x]
        } else if len(g[x]) == 0 {
            maxDepth[x] = 1 + ans[edges[x]] - 1
            return maxDepth[x]
        }
        t := 1
        for _, y := range g[x] {
            t = max(t, dfs(y) + 1)
        }
        maxDepth[x] = t
        return t
    }

    find := func(x int) int {
        if maxDepth[x] == 0 {
            maxDepth[x] = dfs(x)
        }
        return maxDepth[x]
    }

    for i := 0; i < n; i++ {
        if in[i] != 0 {
            a := []int{}
            cnt := 0
            for x := i; in[x] != 0; x = edges[x] {
                in[x] = 0
                a = append(a, x)
                cnt++
            }
            for _, x := range a {
                ans[x] = cnt
            }
        }
    }
    for i := 0; i < n; i++ {
        if ans[i] != 0 {
            continue
        }

        ans[i] = find(i)
    }
    
    return ans
}