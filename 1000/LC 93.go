func restoreIpAddresses(s string) (ans []string) {
    path := []byte{}
    n := len(s)

    var dfs func(int, int)
    dfs = func(idx, cnt int) {
        if idx == n && cnt == 4 {
            ans = append(ans, string(path[:len(path)-1]))
        } else if idx >= n || cnt >= 4 {
            return
        } 

        t, m := 0, len(path) 
        for i := idx; i < n; i++ {
            t  = t * 10 + int(s[i] - '0')
            if i > idx && s[idx] == '0' || t > 255{
                break
            }
            path = append(path, s[i])
            path = append(path, '.')
            dfs(i + 1, cnt + 1)
            path = path[:len(path)-1]
        }
        path = path[:m]
    }
    dfs(0, 0)

    return 
}
