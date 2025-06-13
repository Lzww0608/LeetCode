
import "math/bits"
func maximumCost(n int, highways [][]int, k int) (ans int) {
    if k >= n {
        return -1
    }

    g := make([][][2]int, n)
    for _, v := range highways {
        a, b := v[0], v[1]
        g[a] = append(g[a], [2]int{b, v[2]})
        g[b] = append(g[b], [2]int{a, v[2]})
    }

    type pair struct {
        x, vis, sum int 
    }
    q := []pair{}
    vis := make(map[pair]bool)
    for i := 0; i < n; i++ {
        q = append(q, pair{i, 1 << i, 0})
        vis[pair{i, 1 << i, 0}] = true
    }
    

    f := false
    for len(q) > 0 {
        cur := q[0]
        q = q[1:]
        if bits.OnesCount(uint(cur.vis)) == k + 1 {
            f = true
            ans = max(ans, cur.sum)
            continue
        }
        for _, v := range g[cur.x] {
            u, cost := v[0], v[1]
            if cur.vis & (1 << u) != 0 {
                continue
            }
            nxt := pair{u, cur.vis | (1 << u), cur.sum + cost}
            if !vis[nxt] {
                vis[nxt] = true
                q = append(q, nxt)
            }
        }
    }

    if !f {
        return -1
    }

    return 
}