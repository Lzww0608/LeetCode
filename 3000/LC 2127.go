func maximumInvitations(favorite []int) int {
    n := len(favorite)
    in := make([]int, n)
    for _, f := range favorite {
        in[f]++
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
        v := favorite[u]
        g[v] = append(g[v], u)
        if in[v]--; in[v] == 0 {
            q = append(q, v)
        }
    }

    var dfs func(int) int
    dfs = func(x int) int {
        maxDepth := 1
        for _, y := range g[x] {
            maxDepth = max(maxDepth, dfs(y) + 1)
        }
        return maxDepth
    }

    maxRing, maxChain := 0, 0
    for i := 0; i < n; i++ {
        if in[i] == 0 {
            continue
        }

        cnt := 0
        for x := i; in[x] != 0; x = favorite[x] {
            in[x] = 0
            cnt++
        }

        if cnt == 2 {
            maxChain += dfs(i) + dfs(favorite[i])
        } else {
            maxRing = max(maxRing, cnt)
        }
    }

    return max(maxChain, maxRing)
}