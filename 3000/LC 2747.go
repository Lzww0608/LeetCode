func countServers(n int, logs [][]int, x int, queries []int) []int {
    sort.Slice(logs, func(i, j int) bool {
        return logs[i][1] < logs[j][1]
    })   

    m := len(queries)
    ans := make([]int, m)
    id := make([]int, m)
    for i := range id {
        id[i] = i
    }
    sort.Slice(id, func(i, j int) bool {
        return queries[id[i]] < queries[id[j]]
    })

    vis := make(map[int]int)
    l, r := 0, 0
    for _, i := range id {
        v := queries[i]
        for r < len(logs) && logs[r][1] <= v {
            vis[logs[r][0]]++
            r++
        }

        for l < len(logs) && logs[l][1] < v - x {
            x := logs[l][0]
            vis[x]--
            if vis[x] == 0 {
                delete(vis, x)
            }
            l++
        }

        ans[i] = n - len(vis)
    }

    return ans
}