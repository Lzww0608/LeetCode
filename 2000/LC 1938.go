type Node struct {
    son [2]*Node
    cnt int
}

var root *Node

func (c *Node) put(x int) {
    cur := c
    for i := 20; i >= 0; i-- {
        t := (x >> i) & 1 
        y := 0
        if t == 1 {
            y = 1
        }
        if cur.son[y] == nil {
            cur.son[y] = &Node{cnt: 1}
        } else {
            cur.son[y].cnt += 1
        }
        cur = cur.son[y]
    }
}

func (c *Node) del(x int) {
    cur := c
    for i := 20; i >= 0; i-- {
        t := (x >> i) & 1 
        y := 0
        if t == 1 {
            y = 1
        }
        cur.son[y].cnt -= 1
        cur = cur.son[y]
    }
}

func (c *Node) maxXor(x int) (ans int) {
    cur := c
    for i := 20; i >= 0 && cur != nil; i-- {
        t := (x >> i) & 1 
        y := 1
        if t == 1 {
            y = 0
        }
        if cur.son[y] != nil && cur.son[y].cnt > 0 {
            ans |= 1 << i 
            cur = cur.son[y]
        } else {
            cur = cur.son[y ^ 1]
        }
    }

    return
}

func maxGeneticDifference(parents []int, queries [][]int) []int {
    n := len(parents)
    type query struct {
        v, i int
    }
    g := make([][]int, n)
    tree := -1
    for i, x := range parents {
        if x == -1 {
            tree = i 
        } else {
            g[x] = append(g[x], i)
        }
    }

    q := make([][]query, n)
    for i, v := range queries {
        a, b := v[0], v[1]
        q[a] = append(q[a], query{b, i})
    }

    ans := make([]int, len(queries))
    root = &Node{}

    var dfs func(int)
    dfs = func(u int) {
        root.put(u)
        for _, t := range q[u] {
            ans[t.i] = root.maxXor(t.v)
        }
        
        for _, v := range g[u] {
            dfs(v)
        }
        root.del(u)
    }
    dfs(tree)

    return ans
}