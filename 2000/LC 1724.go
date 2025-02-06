type Node struct {
    limit, parent int
}

type DistanceLimitedPathsExist struct {
    f []Node
}

func (c *DistanceLimitedPathsExist)find(x int) int {
    for c.f[x].parent >= 0 {
        x = c.f[x].parent
    }
    return x
}

func (c *DistanceLimitedPathsExist)findWithLimit(x, limit int) int {
    for c.f[x].parent >= 0 && c.f[x].limit < limit {
        x = c.f[x].parent
    }
    return x
}

func (c *DistanceLimitedPathsExist)merge(x, y, limit int) bool {
    rx, ry := c.find(x), c.find(y)
    if (rx == ry) {
        return false
    }

    fx, fy := c.f[rx].parent, c.f[ry].parent
    if fx < fy {
        rx, ry = ry, rx
    }
    c.f[rx].parent = fx + fy
    c.f[ry].parent = rx 
    c.f[ry].limit = limit
    return true
}


func Constructor(n int, edgeList [][]int) DistanceLimitedPathsExist {
    f := make([]Node, n)
    for i := range f {
        f[i] = Node{0, -1}
    }
    sort.Slice(edgeList, func(i, j int) bool {
        return edgeList[i][2] < edgeList[j][2]
    })
    c := DistanceLimitedPathsExist{f}
    for _, v := range edgeList {
        c.merge(v[0], v[1], v[2])
    }

    return c
}


func (c *DistanceLimitedPathsExist) Query(p int, q int, limit int) bool {
    return c.findWithLimit(p, limit) == c.findWithLimit(q, limit)
}


/**
 * Your DistanceLimitedPathsExist object will be instantiated and called as such:
 * obj := Constructor(n, edgeList);
 * param_1 := obj.Query(p,q,limit);
 */