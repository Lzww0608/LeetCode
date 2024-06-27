func findShortestCycle(n int, edges [][]int) int {
    g := make([][]int, n)
    
    for _, e := range edges {
        a, b := e[0], e[1]
        g[a] = append(g[a], b)
        g[b] = append(g[b], a)
    }
    
    ans := math.MaxInt32
    type pair struct {
        x, fa int
    }
    
    for i := 0; i < n; i++ {
        q := []pair{pair{i, -1}}
        dis := make([]int, n)
        for i := range dis {
            dis[i] = -1
        }
        dis[i] = 0
        for len(q) > 0 {
            for k := len(q); k > 0; k-- {
                t := q[0]
                q = q[1:]
                for _, x := range g[t.x] {
                    if dis[x] < 0 {
                        dis[x] = dis[t.x] + 1
                        q = append(q, pair{x, t.x})
                    } else if x != t.fa {
                        ans = min(ans, dis[x] + dis[t.x] + 1)
                    }
                }
            }
        }
    }

    if ans == math.MaxInt32 {
        return -1
    }

    return ans
}
