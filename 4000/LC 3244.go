func shortestDistanceAfterQueries(n int, queries [][]int) []int {
    m := map[int]int{}
    for i := 0; i < n - 1; i++ {
        m[i] = i + 1
    }

    ans := make([]int, len(queries))
    for k, q := range queries {
        l, r := q[0], q[1]
        
        if _, ok := m[l]; ok && m[l] < r {
            v := l
            for v < r {
                t := v
                v = m[v]
                delete(m, t)
            }
            m[l] = r
        }

        ans[k] = len(m)
    }

    return ans
}