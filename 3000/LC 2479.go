const N int = 60
type Node struct {
    Left, Right *Node
}

func (c *Node) insert(val int) {
    for i := N; i >= 0; i-- {
        if val & (1 << i) != 0 {
            if c.Left == nil {
                c.Left = &Node{}
            }
            c = c.Left
        } else {
            if c.Right == nil {
                c.Right = &Node{}
            }
            c = c.Right
        }
    }
}

func (c *Node) query(val int) (ans int) {
    for i := N; i >= 0 && c != nil; i-- {
        if val & (1 << i) != 0 {
            if c.Right != nil {
                ans |= 1 << i 
                c = c.Right
            } else {
                c = c.Left
            }
        } else {
            if c.Left != nil {
                ans |= 1 << i 
                c = c.Left
            } else {
                c = c.Right
            }
        }
    }

    return
}

func maxXor(n int, edges [][]int, values []int) int64 {
    root := &Node{}

    g := make([][]int, n)
    for _, edge := range edges {
        a, b := edge[0], edge[1]
        g[a] = append(g[a], b)
        g[b] = append(g[b], a)
    }

    s := make([]int, n)
    var dfs1 func(int, int) int
    dfs1 = func(u, fa int) int {
        cur := values[u]
        for _, v := range g[u] {
            if v == fa {
                continue
            }
            cur += dfs1(v, u)
        }
        s[u] = cur
        return cur
    }
    dfs1(0, -1)

    ans := 0
    var dfs2 func(int, int)
    dfs2 = func(u, fa int) {
        ans = max(ans, root.query(s[u]))
        for _, v := range g[u] {
            if v == fa {
                continue
            }
            dfs2(v, u)
        }
        root.insert(s[u])
    } 
    dfs2(0, -1)

    return int64(ans)
}