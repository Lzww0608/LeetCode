const N int = 26
func maxSubstringLength(s string, k int) bool {
    if k == 0 {
        return true
    }

    pos := [N][]int{}
    for i, c := range s {
        x := int(c - 'a')
        pos[x] = append(pos[x], i)
    }

    g := [N][]int{}
    for i, v := range pos {
        if v == nil {
            continue
        }
        n := len(v)
        l, r := v[0], v[n - 1]
        for j, u := range pos {
            if j == i {
                continue
            }
            t := sort.SearchInts(u, l)
            if t < len(u) && u[t] <= r {
                g[i] = append(g[i], j)
            }
        }
    }

    vis := [N]bool{}
    var l, r int 
    var dfs func(int) 
    dfs = func(x int) {
        vis[x] = true 
        l = min(l, pos[x][0])
        r = max(r, pos[x][len(pos[x]) - 1])
        for _, y := range g[x] {
            if !vis[y] {
                dfs(y)
            }
        }
    }

    intervals := [][2]int{}
    for i, v := range pos {
        if v == nil {
            continue
        }

        vis = [N]bool{}
        l, r = len(s), 0 
        dfs(i)
        if l > 0 || r < len(s) - 1 {
            intervals = append(intervals, [2]int{l, r})
        }
    }

    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][1] < intervals[j][1] 
    })
    cur := -1
    for _, v := range intervals {
        if v[0] > cur {
            k--
            cur = v[1]
        }
    }

    return k <= 0
}