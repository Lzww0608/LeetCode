const N int = 26
func findSubtreeSizes(parent []int, s string) []int {
    n := len(s)
    ans := make([]int, n)
    for i := range ans {
        ans[i] = 1
    }
    g := make([][]int, n)
    for i := 1; i < n; i++ {
        j := parent[i]
        g[j] = append(g[j], i)
    }

    fa := [N]int{}
    for i := range fa {
        fa[i] = -1
    }
    var dfs func(int) 
    dfs = func(u int) {
        x := int(s[u] - 'a')
        pre := fa[x]
        fa[x] = u 
        for _, v := range g[u] {
            dfs(v)
            cur := fa[int(s[v] - 'a')]
            if cur < 0 {
                ans[u] += ans[v] 
            } else {
                ans[cur] += ans[v]  
            }
        }
        fa[x] = pre
    }
    dfs(0)

    return ans
}