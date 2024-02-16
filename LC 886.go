func possibleBipartition(n int, dislikes [][]int) bool {
    g := make([][]int, n + 1)
    for _, e := range dislikes {
        a, b := e[0], e[1]
        g[a] = append(g[a], b)
        g[b] = append(g[b], a)
    }
    col := make([]int, n + 1)
    for i := range g {
        if col[i] == 0 {
            col[i] = 1
            q := []int{i}
            for len(q) > 0 {
                for k := len(q); k > 0; k-- {
                    t := q[0]
                    q = q[1:]
                    for _, e := range g[t] {
                        if col[e] == col[t] {
                            return false
                        } else if col[e] == 0 {
                            q = append(q, e)
                        }
                        col[e] = 3 - col[t]
                    }
                }
            }
        }
    }
    
    return true
}