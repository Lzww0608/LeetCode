func buildMatrix(k int, rowConditions [][]int, colConditions [][]int) [][]int {
    row := topoSort(k, rowConditions)
    col := topoSort(k, colConditions)
    if row == nil || col == nil {
        return nil
    }

    pos := make([]int, k)
    for i, v := range col {
        pos[v] = i
    }
    ans := make([][]int, k)
    for i, x := range row {
        ans[i] = make([]int, k)
        ans[i][pos[x]] = x + 1
    }

    return ans
}

func topoSort(k int, edges [][]int) []int {
    g := make([][]int, k)
    indeg := make([]int, k)
    for _, e := range edges {
        a, b := e[0] - 1, e[1] - 1
        g[a] = append(g[a], b)
        indeg[b]++
    }

    q := make([]int, 0, k)
    p := q
    for i, x := range indeg {
        if x == 0 {
            q = append(q, i)
        }
    }

    for len(q) > 0 {
        x := q[0]
        q = q[1:]
        for _, y := range g[x] {
            if indeg[y]--; indeg[y] == 0 {
                q = append(q, y)
            }
        }
    }
    if cap(q) > 0 {
        return nil
    }

    return p[:k]
}

