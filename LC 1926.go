func nearestExit(maze [][]byte, entrance []int) int {
    m, n := len(maze), len(maze[0])
    type pair struct {
        x, y int
    }
    q := []pair{pair{entrance[0], entrance[1]}}
    vis := make([]bool, m * n)
    vis[entrance[0] * n + entrance[1]] = true
    
    d := 0
    dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
    for len(q) > 0 {
        for k := len(q); k > 0; k-- {
            i, j := q[0].x, q[0].y
            q = q[1:]
            if (i == 0 || i == m - 1 || j == 0 || j == n - 1) && (maze[i][j] != '+' && d > 0) {
                return d
            }
            for _, dir := range dirs {
                x, y := i + dir[0], j + dir[1]
                if x >= 0 && x < m && y >= 0 && y < n && !vis[x * n + y] && maze[x][y] == '.' {
                    vis[x * n + y] = true
                    q = append(q, pair{x, y})
                }
            } 
        }
        d++
    }

    return -1
}