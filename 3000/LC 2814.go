func minimumSeconds(land [][]string) int {
    dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
    m, n := len(land), len(land[0])
    vis := make([]bool, m * n)
    q1 := [][]int{}
    q2 := [][]int{}
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if land[i][j] == "*" {
                q2 = append(q2, []int{i, j})
            } else if land[i][j] == "S" {
                q1 = append(q1, []int{i, j})
                vis[i * n + j] = true
            }
        }
    }

    d := 1
    for len(q1) > 0 {
        for sz := len(q2); sz > 0; sz-- {
            cur := q2[0]
            q2 = q2[1:]
            i, j := cur[0], cur[1]
            for _, dir := range dirs {
                x, y := i + dir[0], j + dir[1]
                if x >= 0 && x < m && y >= 0 && y < n && land[x][y] == "." {
                    land[x][y] = "*"
                    q2 = append(q2, []int{x, y})
                }
            }
        }

        for sz := len(q1); sz > 0; sz-- {
            cur := q1[0]
            q1 = q1[1:]
            i, j := cur[0], cur[1]
            for _, dir := range dirs {
                x, y := i + dir[0], j + dir[1]
                if x < 0 || x >= m || y < 0 || y >= n || vis[x * n + y] || land[x][y] == "*" || land[x][y] == "X" {
                    continue
                }
                if land[x][y] == "D" {
                    return d
                }
                vis[x * n + y] = true
                q1 = append(q1, []int{x, y})
            }
        }
        d++
    }

    return -1
}