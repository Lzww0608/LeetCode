func isBipartite(graph [][]int) bool {
    n := len(graph)

    col := make([]int, n)
    for i := range col {
        col[i] = -1
    }

    for i := range col {
        if col[i] == -1 {
            q := []int{i}
            col[i] = 0

            for len(q) > 0 {
                u := q[0]
                q = q[1:]
                for _, v := range graph[u] {
                    if col[v] == -1 {
                        col[v] = 1 - col[u]
                        q = append(q, v)
                    } else if col[v] == col[u] {
                        return false
                    }
                }
            }
        }
    }

    return true
}