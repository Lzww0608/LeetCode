func goodsOrder(goods string) (ans []string) {
    s := []byte(goods)
    sort.Slice(s, func(i, j int) bool {
        return s[i] < s[j] 
    })

    n := len(s)
    path := []byte{}
    m := map[int]bool{}
    var dfs func(int)
    dfs = func(idx int) {
        if idx == n {
            ans = append(ans, string(path))
            return
        }
        for i := 0; i < n; i++ {
            if m[i] {
                continue
            } else if i > 0 && s[i] == s[i-1] && !m[i-1] {
                continue
            }
            path = append(path, s[i])
            m[i] = true
            dfs(idx + 1)
            m[i] = false
            path = path[:len(path)-1]
        }
    }
    dfs(0)

    return
}