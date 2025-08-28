func minimumCost(target string, words []string, costs []int) int {
    m := make(map[string]int)
    for i, word := range words {
        if v, ok := m[word]; !ok {
            m[word] = costs[i]
        } else {
            m[word] = min(v, costs[i])
        }
    }

    n := len(target)
    f := make([]int, n + 1)
    for i := range f {
        f[i] = math.MaxInt32
    }
    f[0] = 0

    for i := 0; i < n; i++ {
        for k, v := range m {
            if i + len(k) <= n && k == target[i: i + len(k)] {
                f[i + len(k)] = min(f[i + len(k)], f[i] + v)
            }
        }
    }

    if f[n] == math.MaxInt32 {
        return -1
    }

    return f[n]
}