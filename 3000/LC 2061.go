func numberOfCleanRooms(room [][]int) int {
    dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
    m, n := len(room), len(room[0])
    vis := make([][4]bool, m * n)
    vis[0][0] = true
    ans := make(map[int]bool)

    q := [][]int{{0, 0}}
    for len(q) > 0 {
        cur := q[0]
        q = q[1:]
        i, j, k := cur[0] / n, cur[0] % n, cur[1]
        ans[cur[0]] = true
        for t := 0; t < 4; t++ {
            dir := dirs[(t+k)%4]
            x, y := i + dir[0], j + dir[1]
            if x < 0 || x >= m || y < 0 || y >= n || room[x][y] == 1 {
                    continue
            } else {
                d := (t + k) % 4
                if !vis[x * n + y][d] {
                    vis[x *n + y][d] = true
                    q = append(q, []int{x * n + y, d})
                }
                break
            }
        }
    }

    return len(ans)
}