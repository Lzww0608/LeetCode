func racecar(target int) int {
    type pair struct {
        pos, sp int
    }

    vis := make(map[pair]bool)
    vis[pair{0, 1}] = true

    q := []pair{{0, 1}}
    d := 0
    for len(q) > 0 {
        for t := len(q); t > 0; t-- {
            cur := q[0]
            q = q[1:]
            if cur.pos == target {
                return d 
            }

            next := pair{cur.pos + cur.sp, cur.sp * 2}
            if next.pos > 0 && next.pos < target * 2 && !vis[next] {
                vis[next] = true
                q = append(q, next)
            }
            
            next = pair{cur.pos, 1}
            if cur.sp > 0 {
                next.sp = -1
            }
            
            if next.pos > 0 && next.pos < target * 2 && !vis[next] {
                vis[next] = true
                q = append(q, next)
            }
             
        }
        d++
    }

    return -1
}