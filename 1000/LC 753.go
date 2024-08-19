func crackSafe(n int, k int) string {
    mod := int(math.Pow(10, float64(n - 1)))
    ans := []byte{}
    vis := map[int]bool{}

    var dfs func(int)
    dfs = func(u int) {
        for i := 0; i < k; i++ {
            v := u * 10 + i 
            if vis[v] {
                continue
            }
            vis[v] = true
            dfs(v % mod)
            ans = append(ans, byte('0' + i))
        }
    }
    dfs(0)
    for i := 0; i < n - 1; i++ {
        ans = append(ans, '0')
    }

    return string(ans)
}