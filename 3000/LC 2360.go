func longestCycle(edges []int) int {
    n := len(edges)
    ans := -1
    
    in := make([]int, n)
    for _, x := range edges {
        if x != -1 {
            in[x]++
        }
    }

    q := []int{}
    for i, x := range in {
        if x == 0 {
            q = append(q, i)
        }
    }

    for len(q) > 0 {
        u := q[0]
        q = q[1:]
        if edges[u] != -1 {
            if in[edges[u]]--; in[edges[u]] == 0 {
                q = append(q, edges[u])
            }
        }
    }

    vis := make([]bool, n)
    var dfs func(int) int 
    dfs = func(i int) (res int) {
        if vis[i] {
            return
        }
        vis[i] = true
        res += 1 + dfs(edges[i])
        return
    }

    
    for i, x := range in {
        if x > 0 && !vis[i] {
            ans = max(ans, dfs(i))
        }
    }

    return ans
}