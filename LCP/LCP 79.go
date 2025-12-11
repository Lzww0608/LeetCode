func extractMantra(matrix []string, mantra string) int {
    m, n := len(matrix), len(matrix[0])
    sz := len(mantra)
    dirs := [][]int{[]int{0, 1}, []int{0, -1}, []int{1, 0}, []int{-1, 0}}

    vis := make([][][]bool, m)
    for i := range m {
        vis[i] = make([][]bool, n)
        for j := range n {
            vis[i][j] = make([]bool, sz)
        }
    }

    type pair struct {
        x, y, z int
    }
    q := []pair{{0, 0, 0}}
    vis[0][0][0] = true

    d := 1
    for len(q) > 0 {
        for t := len(q); t > 0; t-- {
            cur := q[0]
            q = q[1:]
            i, j, k := cur.x, cur.y, cur.z 
            if matrix[i][j] == mantra[k] {
                if k == sz - 1 {
                    return d
                }
                if !vis[i][j][k + 1] {
                    vis[i][j][k + 1] = true
                    q = append(q, pair{i, j, k + 1})
                }
            } else {
                for _, dir := range dirs {
                    x, y := i + dir[0], j + dir[1]
                    if x < 0 || x >= m || y < 0 || y >= n || vis[x][y][k] {
                        continue
                    }
                    vis[x][y][k] = true
                    q = append(q, pair{x, y, k})
                }
            }
        }
        d++
    }

    return -1
}