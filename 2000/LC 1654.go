func minimumJumps(forbidden []int, a int, b int, x int) int {
    f := slices.Max(forbidden)
    limit := max(f + a + b, x + b)

    m := make([]int, limit + 1)
    for _, x := range forbidden {
        m[x] = 1
    }

    dist := make([][2]int, limit + 1)
    for i := range dist {
        dist[i][0] = math.MaxInt32
        dist[i][1] = math.MaxInt32
    }

    dist[0][0] = 0
    q := [][]int{{0, 0}}
    for len(q) > 0 {
        v := q[0]
        q = q[1:]
        i, pre := v[0], v[1]
        if i == x {
            return dist[i][pre]
        }
        if pre == 0 && i - b >= 0 && m[i-b] == 0 && dist[i][pre] + 1 < dist[i-b][1] {
            dist[i-b][1] = dist[i][pre] + 1
            q = append(q, []int{i - b, 1})
        }
        if i + a <= limit && m[i+a] == 0 && dist[i][pre] + 1 < dist[i+a][0] {
            dist[i+a][0] = dist[i][pre] + 1
            q = append(q, []int{i + a, 0})
        }
    }

    return -1
}