type pair struct {
    x, y int
}
func visitOrder(points [][]int, direction string) (ans []int) {
    n := len(points)
    vis := make([]bool, n)
    vp := make([]pair, n)
    for i := range n {
        vp[i] = pair{points[i][0], points[i][1]}
    }

    calc := func(q, m, j int) int {
        a, b, c := vp[q], vp[m], vp[j]
        l := pair{a.x - b.x, a.y - b.y}
        r := pair{c.x - b.x, c.y - b.y}

        return l.x * r.y - l.y * r.x
    }

    m := 0
    for i := range n {
        if vp[i].x < vp[m].x {
            m = i
        }
    }

    ans = append(ans, m)
    vis[m] = true 
    for i := 0; i < n - 2; i++ {
        q := -1
        for j := range n {
            if vis[j] {
                continue
            }

            if q == -1 {
                q = j
            } else {
                t := calc(q, m, j)
                if direction[i] == 'L' {
                    if t < 0 {
                        q = j
                    }
                } else if t > 0 {
                    q = j
                }
            }
        }
        m = q
        vis[q] = true 
        ans = append(ans, m)
    }

    for i := range n {
        if !vis[i] {
            ans = append(ans, i)
        }
    }

    return
}