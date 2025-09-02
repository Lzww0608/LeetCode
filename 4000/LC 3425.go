func longestSpecialPath(edges [][]int, nums []int) []int {
    n := len(edges) + 1
    type pair struct {
        u, w int 
    }
    g := make([][]pair, n)
    for _, edge := range edges {
        u, v, w := edge[0], edge[1], edge[2]
        g[u] = append(g[u], pair{v, w}) 
        g[v] = append(g[v], pair{u, w}) 
    }

    ans := []int{0, 1}
    dis := []int{0}
    m := make(map[int]int)
    var dfs func(int, int, int) 
    dfs = func(u, fa, d int) {
        x := nums[u]
        pre := m[x]
        d = max(d, pre)

        cur := dis[len(dis) - 1] - dis[d]
        node := len(dis) - d 
        if cur > ans[0] || cur == ans[0] && node < ans[1] {
            ans[0], ans[1] = cur, node
        }

        m[x] = len(dis)
        for _, v := range g[u] {
            if v.u == fa {
                continue
            }
            dis = append(dis, dis[len(dis) - 1] + v.w)
            dfs(v.u, u, d)
            dis = dis[:len(dis) - 1]
        }
        m[x] = pre
    }
    dfs(0, -1, 0)

    return ans
}