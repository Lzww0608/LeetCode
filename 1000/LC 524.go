func updateMatrix(mat [][]int) [][]int {
    dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
    m, n := len(mat), len(mat[0])
    vis := make([]bool, m * n)

    q := [][]int{}
    for i, row := range mat {
        for j, x := range row {
            if x == 0 {
                q = append(q, []int{i, j})
                vis[i * n + j] = true
            }
        }
    }

    for len(q) > 0 {
        node := q[0]
        q = q[1:]
        i, j := node[0], node[1]

        for _, dir := range dirs {
            x, y := i + dir[0], j + dir[1]
            if x >= 0 && x < m && y >= 0 && y < n && !vis[x * n + y] {
                vis[x * n + y] = true
                mat[x][y] = mat[i][j] + 1
                q = append(q, []int{x, y})
            }
        }
    }

    return mat
}