func secondMinimum(n int, edges [][]int, time int, change int) int {
    g := make([][]int, n)
    for _, e := range edges {
        a, b := e[0] - 1, e[1] - 1
        g[a] = append(g[a], b)
        g[b] = append(g[b], a)
    }

    next := func(d int) int {
        if times := d / change; times & 1 == 1 {
            return (times + 1) * change + time
        }
        return d + time
    }

    dis := make([][2]int, n)
    for i := range dis {
        dis[i][0], dis[i][1] = math.MaxInt32, math.MaxInt32
    }
    dis[0][0] = 0
    q := [][]int{{0, 0}}

    for len(q) > 0 {
        node := q[0]
        u, d := node[0], node[1]
        q = q[1:]
        time := next(d)
        for _, v := range g[u] {
            if time < dis[v][0] {
                dis[v][0] = time
                q = append(q, []int{v, time})
            } else if time > dis[v][0] && time < dis[v][1] {
                dis[v][1] = time
                q = append(q, []int{v, time})
            }
        } 
    }

    return dis[n-1][1]
} 