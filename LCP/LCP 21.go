func chaseGame(edges [][]int, startA int, startB int) int {
    startA--
    startB--
    n := len(edges)

    g := make([][]int, n)
    deg := make([]int, n)
    for _, edge := range edges {
        u, v := edge[0] - 1, edge[1] - 1
        if u == startA && v == startB || u == startB && v == startA {
            return 1
        }
        g[u] = append(g[u], v)
        g[v] = append(g[v], u)
        deg[u]++
        deg[v]++
    }

    q := []int{}
    s := make(map[int]bool)
    for i := 0; i < n; i++ {
        if deg[i] == 1 {
            q = append(q, i)
        }
    }

    for len(q) > 0 {
        u := q[0]
        q = q[1:]
        s[u] = true
        for _, v := range g[u] {
            if u == v {
                continue
            }
            if deg[v]--; deg[v] == 1 {
                q = append(q, v)
            }
        }
    }

    if !s[startA] && !s[startB] {
        if n - len(s) <= 3 {
            return 1
        } else {
            return -1
        }
    }

    bfs := func(x int) []int {
        dis := make([]int, n)
        vis := make([]bool, n)
        vis[x] = true
        q := []int{x}
        d := 0
        for len(q) > 0 {
            for sz := len(q); sz > 0; sz-- {
                u := q[0]
                q = q[1:]
                dis[u] = d
                for _, v := range g[u] {
                    if !vis[v] {
                        vis[v] = true
                        q = append(q, v)
                    }
                }
            }
            d++
        }
        return dis
    }

    disA := bfs(startA)
    disB := bfs(startB)

    findEntrance := func(x int) int {
        vis := make([]bool, n)
        q := []int{x}
        for len(q) > 0 {
            u := q[0]
            q = q[1:]
            if !s[u] {
                return u
            }
            for _, v := range g[u] {
                if !vis[v] {
                    vis[v] = true
                    q = append(q, v)
                }
            }
        }
        return -1
    }
    entrance := findEntrance(startB)
    if n - len(s) > 3 {
        if !s[startB] {
            return -1
        }
        if disA[entrance] > disB[entrance] + 1 {
            return -1
        }
    }

    ans := 0
    for i := 0; i < n; i++ {
        if disA[i] > disB[i] + 1 {
            ans = max(ans, disA[i])
        }
    }
    return ans
}