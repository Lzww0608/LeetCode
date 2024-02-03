func partition(s string) (ans [][]string) {
    n := len(s)
    path := []string{}
    var dfs func(int)
    dfs = func(i int) {
        if i == n {
            ans = append(ans, append([]string(nil), path...))
        }
        t := ""
        for j := i; j < n; j++ {
            t += string(s[j])
            f := true
            for k := 0; k * 2 < len(t); k++ {
                if t[k] != t[len(t)-k-1] {
                    f = false
                    break
                }
            }
            if f {
                path = append(path, t)
                dfs(j+1)
                path = path[:len(path)-1]
            }
        }
    }
    dfs(0)
    return
}