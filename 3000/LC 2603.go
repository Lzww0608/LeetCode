func collectTheCoins(coins []int, edges [][]int) int {
    n := len(coins)
    g := make([][]int, n)
    deg := make([]int, n)

    for _, e := range edges {
        u, v := e[0], e[1]
        g[u] = append(g[u], v)
        g[v] = append(g[v], u)
        deg[u]++
        deg[v]++
    } 

    cnt := n - 1
    q := []int{}
    for i, x := range deg {
        if x == 1 && coins[i] == 0 {
            deg[i]--
            q = append(q, i)
        }
    }

    for len(q) > 0 {
        cnt--
        u := q[0]
        q = q[1:]
        for _, v := range g[u] {
            if deg[v]--; deg[v] == 1 && coins[v] == 0 {
                q = append(q, v)
            }
        }
    }

    for i, x := range deg {
        if x == 1 && coins[i] == 1 {
            deg[i]--
            q = append(q, i)
        }
    }
    cnt -= len(q)
    for len(q) > 0 {
        u := q[0]
        q = q[1:]
        for _, v := range g[u] {
            if deg[v]--; deg[v] == 1 {
                cnt--
            }
        }
    }

    return max(0, cnt * 2)
}