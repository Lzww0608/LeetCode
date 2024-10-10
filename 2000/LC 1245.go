func treeDiameter(edges [][]int) int {
    n := len(edges) + 1
    g := make([][]int, n)

    for _, e := range edges {
        u, v := e[0], e[1]
        g[u] = append(g[u], v)
        g[v] = append(g[v], u)
    }

    bfs := func(start int) (int, int) {
        vis := make([]bool, n)
        farthestNode, distance := start, 0

        q := []int{start}
        vis[start] = true

        for len(q) > 0 {
            for t := len(q); t > 0; t-- {
                u := q[0]
                q = q[1:]
                farthestNode = u 
                for _, v := range g[u] {
                    if !vis[v] {
                        vis[v] = true
                        q = append(q, v)
                    }
                }
            }
            distance++
        }

        return farthestNode, distance - 1
    }

    node, _ := bfs(0)
    _, ans := bfs(node)

    return ans
}