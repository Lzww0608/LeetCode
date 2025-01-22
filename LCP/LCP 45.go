func bicycleYard(position []int, terrain [][]int, obstacle [][]int) (ans [][]int) {
    dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
    m, n := len(terrain), len(terrain[0])
    vis := make([][][102]bool, m)
    for i := range vis {
        vis[i] = make([][102]bool, n)
    }

    q := [][]int{{position[0], position[1], 1}}
    vis[position[0]][position[1]][1] = true
    for len(q) > 0 {
        cur := q[0]
        q = q[1:]
        i, j, s := cur[0], cur[1], cur[2]
        if s == 1 {
            ans = append(ans, []int{i, j})
        }
        for _, dir := range dirs {
            x, y := i + dir[0], j + dir[1]
            if x < 0 || x >= m || y < 0 || y >= n {
                continue
            }
            sp := s + terrain[i][j] - terrain[x][y] - obstacle[x][y]
            if sp > 0 && !vis[x][y][sp] {
                vis[x][y][sp] = true
                q = append(q, []int{x, y, sp})
            }
        }
    }

    ans = ans[1:]
    sort.Slice(ans, func(i, j int) bool {
        return ans[i][0] < ans[j][0] || ans[i][0] == ans[j][0] && ans[i][1] < ans[j][1]
    })

    return ans
}