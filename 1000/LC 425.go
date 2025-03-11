func wordSquares(words []string) (ans [][]string) {
    n := len(words[0])
    m := make(map[string][]string)
    m[""] = words

    for _, s := range words {
        for i := range s {
            m[s[:i+1]] = append(m[s[:i+1]], s)
        }
    }

    path := []string{}

    var dfs func(int) 
    dfs = func(idx int) {
        if idx == n {
            ans = append(ans, append([]string(nil), path...))
            return
        }

        s := []byte{}
        for i := range idx {
            s = append(s, path[i][idx])
        }
        for _, t := range m[string(s)] {
            path = append(path, t)
            dfs(idx + 1)
            path = path[:len(path)-1]
        }
    }
    dfs(0)

    return
}