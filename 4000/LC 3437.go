func permute(n int) (ans [][]int) {
    vis := make([]bool, n)
    tmp := []int{}

    var dfs func(int)
    dfs = func(pre int) {
        if len(tmp) == n {
            ans = append(ans, append([]int(nil), tmp...))
            return
        }

        for i := 0; i < n; i++ {
            if vis[i] {
                continue
            }
            if pre == -1 || pre & 1 != (i + 1) & 1 {
                vis[i] = true 
                tmp = append(tmp, i + 1)
                dfs(i + 1) 
                tmp = tmp[:len(tmp) - 1]
                vis[i] = false
            }
        }
    }
    dfs(-1)

    sort.Slice(ans, func(i, j int) bool {
        return slices.Compare(ans[i], ans[j]) == -1
    })

    return 
}