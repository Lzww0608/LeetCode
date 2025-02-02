const N int = 26
func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
    m := make(map[int]int)
    for i := range original {
        x := int(original[i] - 'a')
        y := int(changed[i] - 'a')
        t := (1 << x) + (1 << (y + 32))
        if _, ok := m[t]; !ok {
            m[t] = cost[i]
        } else {
            m[t] = min(m[t], cost[i])
        }
    }

    for k := 0; k < N; k++ {
        for i := 0; i < N; i++ {
            for j := 0; j < N; j++ {
                t := (1 << i) + (1 << (j + 32))
                x := (1 << i) + (1 << (k + 32))
                y := (1 << k) + (1 << (j + 32))
                _, ok1 := m[x]
                _, ok2 := m[y]
                if ok1 && ok2 {
                    if _, ok := m[t]; !ok {
                        m[t] = m[x] + m[y]
                    } else {
                        m[t] = min(m[t], m[x] + m[y])
                    }
                }
            }
        }
    }

    ans := 0
    for i := range source {
        if source[i] == target[i] {
            continue
        }
        x := int(source[i] - 'a')
        y := int(target[i] - 'a')
        t := (1 << x) + (1 << (y + 32))
        if v, ok := m[t]; !ok {
            return -1
        } else {
            ans += v
        }
    }

    return int64(ans)
}