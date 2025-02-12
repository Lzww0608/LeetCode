func numberOfPaths(n int, corridors [][]int) int {
    g := make([][]int, n)
    edge := make([]map[int]bool, n)
    for i := 0; i < n; i++ {
        edge[i] = make(map[int]bool)
    }

    for _, v := range corridors {
        a, b := v[0] - 1, v[1] - 1
        g[a] = append(g[a], b)
        g[b] = append(g[b], a)
        edge[a][b] = true
        edge[b][a] = true
    }

    ans := 0
    for i := 0; i < n; i++ {
        for _, j := range g[i] {
            for _, k := range g[j] {
                if edge[k][i] {
                    ans++
                }
            } 
        }
    }

    return ans / 6
}