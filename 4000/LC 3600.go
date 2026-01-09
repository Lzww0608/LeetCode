func maxStability(n int, edges [][]int, k int) int {
    if acycleOrUnconnected(n, edges) {
        return -1
    }

    sort.Slice(edges, func(i, j int) bool {
        if edges[i][3] != edges[j][3] {
            return edges[i][3] == 1
        }
        return edges[i][2] > edges[j][2]
    })
    d := NewDSU(n)

    check := func(mid int) bool {
        d.clear()
        cnt := 0
        for _, e := range edges {
            u, v, s, must := e[0], e[1], e[2], e[3]
            if must == 1 {
                if s < mid {
                    return false 
                }
                d.merge(u, v)
            } else {
                if s * 2 < mid || d.same(u, v) {
                    continue 
                } else if s < mid {
                    if cnt++; cnt > k {
                        return false
                    }
                } 
                d.merge(u, v)
            }
        }

        for i := range n {
            if d.find(i) != d.find(0) {
                return false
            }
        }

        return true 
    }


    l, r := 0, int(1e10 + 1)
    for l + 1 < r {
        mid := l + ((r - l) >> 1)
        if check(mid) {
            l = mid 
        } else {
            r = mid 
        }
    }

    return l
}

func acycleOrUnconnected(n int, edges [][]int) bool {
    d := NewDSU(n)
    for _, e := range edges {
        u, v := e[0], e[1]
        if e[3] == 1 {
            if d.same(u, v) {
                return true
            }
            d.merge(u, v)
        } 
    }

    for _, e := range edges {
        u, v := e[0], e[1]
        d.merge(u, v)
    }

    for i := range n {
        if d.find(i) != d.find(0) {
            return true
        }
    }

    return false
}

type DSU struct {
    n int 
    fa, sz []int 
}

func NewDSU(n int) *DSU {
    fa := make([]int, n)
    sz := make([]int, n)
    for i := range n {
        fa[i] = i 
        sz[i] = 1
    }
    return &DSU {n, fa, sz}
}

func (d *DSU) find(x int) int {
    if d.fa[x] != x {
        d.fa[x] = d.find(d.fa[x])
    }
    return d.fa[x]
}

func (d *DSU) merge(x, y int) {
    rx, ry := d.find(x), d.find(y)
    if d.sz[rx] < d.sz[ry] {
        rx, ry = ry, rx
    }

    d.sz[rx] += d.sz[ry]
    d.fa[ry] = d.fa[rx]
}

func (d *DSU) same(x, y int) bool {
    return d.find(x) == d.find(y)
}

func (d *DSU) clear() {
    for i := range d.n {
        d.fa[i] = i 
        d.sz[i] = 1
    }
}