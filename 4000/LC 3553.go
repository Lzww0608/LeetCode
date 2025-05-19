type TreeAncestor struct {
    depth []int
    dis   []int
    pa    [][]int
}

func Constructor(edges [][]int) *TreeAncestor {
    n := len(edges) + 1
    m := bits.Len(uint(n))
    type edge struct {
        to, wt int
    }
    g := make([][]edge, n)
    for _, e := range edges {
        x, y := e[0], e[1] // 节点编号从 0 开始
        g[x] = append(g[x], edge{y, e[2]})
        g[y] = append(g[y], edge{x, e[2]})
    }

    depth := make([]int, n)
    dis := make([]int, n)
    pa := make([][]int, n)
    var dfs func(int, int)
    dfs = func(x, fa int) {
        pa[x] = make([]int, m)
        pa[x][0] = fa
        for _, e := range g[x] {
            if e.to != fa {
                y := e.to
                depth[y] = depth[x] + 1
                dis[y] = dis[x] + e.wt
                dfs(y, x)
            }
        }
    }
    dfs(0, -1)

    for i := 0; i < m-1; i++ {
        for x := 0; x < n; x++ {
            if p := pa[x][i]; p != -1 {
                pa[x][i+1] = pa[p][i]
            } else {
                pa[x][i+1] = -1
            }
        }
    }
    return &TreeAncestor{depth, dis, pa}
}

func (t *TreeAncestor) GetKthAncestor(node, k int) int {
    for ; k > 0; k &= k - 1 {
        node = t.pa[node][bits.TrailingZeros(uint(k))]
    }
    return node
}

// 返回 x 和 y 的最近公共祖先（节点编号从 0 开始）
func (t *TreeAncestor) GetLCA(x, y int) int {
    if t.depth[x] > t.depth[y] {
        x, y = y, x
    }
    y = t.GetKthAncestor(y, t.depth[y]-t.depth[x]) // 使 y 和 x 在同一深度
    if y == x {
        return x
    }
    for i := len(t.pa[x]) - 1; i >= 0; i-- {
        px, py := t.pa[x][i], t.pa[y][i]
        if px != py {
            x, y = px, py // 同时上跳 2^i 步
        }
    }
    return t.pa[x][0]
}

func (t *TreeAncestor) GetDis(x, y int) int {
    return t.dis[x] + t.dis[y] - t.dis[t.GetLCA(x, y)] * 2
}

func minimumWeight(edges [][]int, queries [][]int) []int {
    t := Constructor(edges)

    ans := make([]int, len(queries))
    for i, q := range queries {
        a, b, c := q[0], q[1], q[2]
        ans[i] = (t.GetDis(a, b) + t.GetDis(b, c) + t.GetDis(a, c)) / 2
    }

    return ans
}