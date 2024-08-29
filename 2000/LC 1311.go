func watchedVideosByFriends(watchedVideos [][]string, friends [][]int, id int, level int) (ans []string) {
    n := len(watchedVideos)
    q := []int{id}
    vis := make([]bool, n)
    vis[id] = true
    for level > 0 {
        for i := len(q); i > 0; i-- {
            cur := q[0]
            q = q[1:]
            for _, x := range friends[cur] {
                if !vis[x] {
                    vis[x] = true
                    q = append(q, x)
                }
            }
        }
        level--
    }

    m := map[string]int{}
    for len(q) > 0 {
        cur := q[0]
        q = q[1:]
        for _, s := range watchedVideos[cur] {
            m[s]++
        }
    }

    for k := range m {
        ans = append(ans, k)
    }

    sort.Slice(ans, func(i, j int) bool {
        return m[ans[i]] < m[ans[j]] || (m[ans[i]] == m[ans[j]] && ans[i] < ans[j])
    })

    return
}