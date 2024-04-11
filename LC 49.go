func groupAnagrams(strs []string) (ans [][]string) {
    m := map[string][]string{}

    for _, s := range strs {
        t := []byte(s)
        sort.Slice(t, func(i, j int) bool {
            return t[i] < t[j]
        })

        m[string(t)] = append(m[string(t)], s)
    }

    for _, v := range m {
        ans = append(ans, v)
    } 

    return
}