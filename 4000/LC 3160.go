func queryResults(limit int, queries [][]int) []int {
    m := make(map[int]int)
    color := make(map[int]int)
    ans := make([]int, len(queries))

    for i, q := range queries {
        x, y := q[0], q[1]
        if v, ok := m[x]; !ok {
            m[x] = y
            color[y]++
        } else {
            color[v]--
            if color[v] == 0 {
                delete(color, v)
            }
            m[x] = y
            color[y]++
        }

        ans[i] = len(color)
    }

    return ans
}