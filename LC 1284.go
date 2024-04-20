func minFlips(mat [][]int) int {
    m, n := len(mat), len(mat[0])
    
    check := func(x int) bool {
        for i := range mat {
            for j := range mat[i] {
                if x & (1 << (i * n + j)) > 0 {
                    return false
                }
            }
        }
        return true
    }

    type pair struct {
        m int
        vis int
    }

    vis := 0
    mt := 0
    for i := range mat {
        for j, x := range mat[i] {
            if x == 1 {
                mt |= 1 << (i * n + j)
            }
        }
    }
    q := []pair{pair{mt, vis}}

    d := 0
    dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {0, 0}}
    for len(q) > 0 {
        for k := len(q); k > 0; k-- {
            t := q[0]
            q = q[1:]
            
            if check(t.m) {
                return d
            }
            
            for i := range mat {
                for j := range mat[i] {
                    if t.vis & (1 << (i * n + j)) == 0 {
                        v := t.m
                        f := t.vis
                        f |= 1 << (i * n + j)
                        for _, dir := range dirs {
                            x, y := i + dir[0], j + dir[1]
                            if x >= 0 && x < m && y >= 0 && y < n {
                                v ^= 1 << (x * n + y)
                            }
                        }
                        q = append(q, pair{v, f})
                    }
                }
            }
        }
        d++
    }

    return -1
}