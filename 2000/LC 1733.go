
func minimumTeachings(n int, languages [][]int, friendships [][]int) int {
    m := len(languages)
    g := make([]map[int]bool, m + 1)
    for i, v := range languages {
        g[i+1] = make(map[int]bool)
        for _, x := range v {
            g[i+1][x] = true
        }
    }
    
    unTalked := [][]int{}
    for _, v := range friendships {
        a, b := v[0], v[1]
        f := true
        for i := 1; i <= n; i++ {
            if g[a][i] && g[b][i] {
                f = false
                break
            } 
        }
        if f {
            unTalked = append(unTalked, v)
        }
    }

    ans := m 
    for i := 1; i <= n; i++ {
        cnt := 0
        vis := make([]bool, m + 1)
        for _, v := range unTalked {
            a, b := v[0], v[1]
            if !vis[a] && !g[a][i] {
                cnt++
                vis[a] = true
            }
            if !vis[b] && !g[b][i] {
                cnt++
                vis[b] = true
            }
        }
        ans = min(ans, cnt)
    }

    return ans
}