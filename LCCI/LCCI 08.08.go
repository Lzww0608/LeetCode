func permutation(str string) (ans []string) {
    path := []byte{}
    s := []byte(str)
    slices.Sort(s)
    n := len(s)
    m := map[int]bool{}

    var dfs func() 
    dfs = func() {
        if len(path) == n {
            ans = append(ans, string(path))
            return
        }

        for i := 0; i < n; i++ {
            if i > 0 && s[i] == s[i-1] && !m[i-1] || m[i] {
                continue
            }
            m[i] = true
            path = append(path, s[i])
            dfs()
            path = path[:len(path)-1]
            m[i] = false
        }
    }
    dfs()

    return
}