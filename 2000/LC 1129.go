type pair struct {
    x int
    y bool
}
func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
    vis, dis, g1, g2 := make([][2]bool, n), make([]int, n), make([][]int, n), make([][]int, n)
    for i := range dis {
        dis[i] = -1
    }
    for _, e := range redEdges {
        g1[e[0]] = append(g1[e[0]], e[1])
    }
    for _, e := range blueEdges {
        g2[e[0]] = append(g2[e[0]], e[1])
    }
    vis[0] = [2]bool{true, true}
    q := []pair{pair{0,true},pair{0,false}}
    for d := 1; len(q) > 0; d++ {
        for i := len(q); i > 0; i-- {
            a, b := q[0].x, q[0].y
            q = q[1:]
            if b {
                for _, e := range g1[a] {
                    if vis[e][0] == false {
                        vis[e][0] = true
                        if dis[e] < 0 {
                            dis[e] = d
                        } else {
                            dis[e] = min(dis[e], d)
                        }
                        q = append(q, pair{e, !b})
                    }
                }
            } else {
                for _, e := range g2[a] {
                    if vis[e][1] == false {
                        vis[e][1] = true
                        if dis[e] < 0 {
                            dis[e] = d
                        } else {
                            dis[e] = min(dis[e], d)
                        }
                        q = append(q, pair{e, !b})
                    }
                }
            }
        }
    }
    dis[0] = 0
    return dis
}
