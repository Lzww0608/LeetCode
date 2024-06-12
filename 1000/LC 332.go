func findItinerary(tickets [][]string) (res []string) {
    var (
        m = map[string][]string{}
    )

    for _, ticket := range tickets {
        a, b := ticket[0], ticket[1]
        m[a] = append(m[a], b)
    }

    for k := range m {
        sort.Strings(m[k])
    }

    var dfs func(string)
    dfs = func(cur string) {
        for {
            if v, ok := m[cur]; !ok || len(v) == 0 {
                break
            }
            tmp := m[cur][0]
            m[cur] = m[cur][1:]
            dfs(tmp)
        }
        res = append(res, cur)
    }
    dfs("JFK")
    
    for i, j := 0, len(res) - 1; i < j; i, j = i + 1, j - 1 {
        res[i], res[j] = res[j], res[i]
    }

    return 
}
