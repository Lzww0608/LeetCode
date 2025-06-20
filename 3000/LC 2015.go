func averageHeightOfBuildings(buildings [][]int) (ans [][]int) {
    type pair struct {
        sum, cnt int
    }
    m := make(map[int]pair)
    for _, v := range buildings {
        p := m[v[0]]
        p.sum += v[2]
        p.cnt++
        m[v[0]] = p
        q := m[v[1]]
        q.sum -= v[2]
        q.cnt--
        m[v[1]] = q
    }
    a := [][3]int{}
    for k, v := range m {
        a = append(a, [3]int{k, v.sum, v.cnt})
    }
    sort.Slice(a, func(i, j int) bool {
        return a[i][0] < a[j][0]
    })

    l, sum, cnt := a[0][0], a[0][1], a[0][2]
    for _, v := range a[1:] {
        if cnt > 0 {
            avg := sum / cnt 
            if len(ans) > 0 && ans[len(ans)-1][2] == avg && ans[len(ans)-1][1] == l {
                ans[len(ans)-1][1] = v[0]
            } else {
                ans = append(ans, []int{l, v[0], sum / cnt})
            }
        }
        
        cnt += v[2]
        sum += v[1]
        l = v[0]
    }

    return
}