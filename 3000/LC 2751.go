func survivedRobotsHealths(positions []int, healths []int, directions string) (ans []int) {
    n := len(positions)
    m := make(map[int]int)
    id := make([]int, n)
    for i := 0; i < n; i++ {
        id[i] = i 
        m[i] = positions[i]
    }
    sort.Slice(id, func(i, j int) bool {
        return positions[id[i]] < positions[id[j]] || positions[id[i]] == positions[id[j]] && directions[id[i]] == 'R'
    })

    st := []int{}
    for _, i := range id {
        f := true
        for len(st) > 0 && directions[st[len(st)-1]] == 'R' && directions[i] == 'L' {
            x, y := healths[st[len(st)-1]], healths[i]
            if y >= x {
                j := st[len(st)-1]
                delete(m, j)
                st = st[:len(st)-1]
                healths[i] -= 1
                if y == x {
                    delete(m, i)
                    f = false
                    break
                }
            } else {
                delete(m, i)
                healths[st[len(st)-1]] -= 1
                f = false
                break
            }
        }
        if f {
            st = append(st, i)
        }
    }

    for i := 0; i < n; i++ {
        if _, ok := m[i]; ok {
            ans = append(ans, healths[i])
        }
    }

    return
}