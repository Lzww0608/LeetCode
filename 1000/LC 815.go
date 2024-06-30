func numBusesToDestination(routes [][]int, source int, target int) int {
    if source == target {
        return 0
    }
    m := map[int]int{}
    mp := map[int]bool{}
    stop := map[int][]int{}

    q := []int{}
    for i, e := range routes {
        for _, x := range e {
            if x == source && m[x] == 0{
                m[x] = 1
                q = append(q, source)
            }
            stop[x] = append(stop[x], i)
        }
    }

    for len(q) > 0 {
        idx := q[0]
        if idx == target {
            return m[idx] - 1
        }
        q = q[1:]
        for _, t := range stop[idx] {
            if _, ok := mp[t]; ok {
                continue
            }
            mp[t] = true
            for _, x := range routes[t] {
                if x == target {
                    return m[idx]
                }
                if _, ok := m[x]; !ok {
                    m[x] = m[idx] + 1
                    q = append(q, x)
                }
            }
        }
    }

    return -1
}
