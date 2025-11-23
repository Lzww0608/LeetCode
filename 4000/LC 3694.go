func distinctPoints(s string, k int) (ans int) {
    type pair struct {
        x, y int
    }
    n := len(s)
    m := make(map[pair]bool)
    f := make([]pair, n + 1)

    cur := pair{0, 0}
    for i := range n {
        if s[i] == 'U' {
            cur.y++
        } else if s[i] == 'D' {
            cur.y--
        } else if s[i] == 'L' {
            cur.x--
        } else {
            cur.x++
        }
        f[i + 1] = cur

        if i >= k - 1 {
            pre := f[i - k + 1]
            x, y := cur.x - pre.x, cur.y - pre.y
            t := pair{x, y}
            if !m[t] {
                m[t] = true
                ans++
            }
        }
    }

    return
}