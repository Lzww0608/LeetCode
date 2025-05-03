func mostVisitedPattern(username []string, timestamp []int, website []string) []string {
    type Node struct {
        name string 
        time int 
        website string
    }
    n := make([]Node, len(username))
    for i := range username {
        n[i] = Node{username[i], timestamp[i], website[i]}
    }
    sort.Slice(n, func(i, j int) bool {
        return n[i].time < n[j].time
    })
    m := make(map[string][]string)
    for i := range n {
        m[n[i].name] = append(m[n[i].name], n[i].website) 
    }

    cnt := make(map[[3]string]int)
    for _, v := range m {
        sz := len(v)
        tmp := make(map[[3]string]bool)
        for i := 0; i < sz; i++ {
            for j := i + 1; j < sz; j++ {
                for k := j + 1; k < sz; k++ {
                    tmp[[3]string{v[i], v[j], v[k]}] = true
                }
            }
        }
        for k := range tmp {
            cnt[k]++
        }
    }

    ans := []string{}
    mx := 0
    for k, v := range cnt {
        if v > mx {
            mx = v 
            ans = []string{k[0], k[1], k[2]}
        } else if v == mx {
            if k[0] < ans[0] || k[0] == ans[0] && k[1] < ans[1] || k[0] == ans[0] && k[1] == ans[1] && k[2] < ans[2] {
                ans = []string{k[0], k[1], k[2]}
            }
        }
    }

    return ans
}