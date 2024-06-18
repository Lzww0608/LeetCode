func shortestPathLength(graph [][]int) (ans int) {
    n := len(graph)
    vis := make([][]bool, n)
    for i := range vis {
        vis[i] = make([]bool, 1 << n)
    }
    type tuple struct {
        cur, mask, dis int
    }
    q := []tuple{}
    for i := range graph {
        q = append(q, tuple{i, 1 << i, 0})
        vis[i][1<<i] = true
    }

    for len(q) > 0 {
        tmp := q[0]
        q = q[1:]
        cur, mask, dis := tmp.cur, tmp.mask, tmp.dis
        if mask == (1 << n) - 1{
            ans = dis
            break
        }
        for _, x := range graph[cur] {
            mask_cur := mask | (1 << x)
            if !vis[x][mask_cur] {
                vis[x][mask_cur] = true
                q = append(q, tuple{x, mask_cur, dis + 1})
            }
        }
    }

    return
}