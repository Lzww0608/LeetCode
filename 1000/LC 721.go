func accountsMerge(accounts [][]string) (ans [][]string) {
    // 离散化
    mails := []string{}
    m := map[string]int{}
    for _, s := range accounts {
        for i, x := range s {
            if i == 0 {
                continue
            }
            if _, ok := m[x]; !ok {
                mails = append(mails, x)
                m[x] = len(mails) - 1
            }
        }
    }

    l := len(mails)
    g := make([][]int, l)
    for _, s := range accounts {
        for i, x := range s {
            if i == 0 {
                continue
            }
            for j := i + 1; j < len(s); j++ {
                a, b := m[x], m[s[j]]
                g[a] = append(g[a], b)
                g[b] = append(g[b], a)
            }
        }
    }

    vis := make([]bool, l)
    var dfs func(int, *[]string) 
    dfs = func(idx int, tmp *[]string) {
        vis[idx] = true
        *tmp = append(*tmp, mails[idx])
        for _, e := range g[idx] {
            if vis[e] == false {
                dfs(e, tmp)
            }
        }
    }

    for _, s := range accounts {
        tmp := []string{s[0]}
        if vis[m[s[1]]] == false {
            dfs(m[s[1]], &tmp)
            sort.Strings(tmp[1:])
            ans = append(ans, tmp)
        }
    }
    return
}