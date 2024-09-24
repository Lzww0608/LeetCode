func containsCycle(grid [][]byte) bool {
    dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
    m, n := len(grid), len(grid[0])
    vis := make([]bool, m * n)
    
bfs := func(i, j int, c byte) bool {
        q := [][]int{{i, j, -1, -1}} // Start BFS with no parent node

        for len(q) > 0 {
            node := q[0]
            q = q[1:]
            x, y := node[0], node[1]

            for _, dir := range dirs {
                nx, ny := x + dir[0], y + dir[1]
                
                if nx >= 0 && nx < m && ny >= 0 && ny < n && grid[nx][ny] == c {
                    if vis[nx * n + ny] {
                        if nx != node[2] || ny != node[3] {
                            return true 
                        }
                    } else {
                        vis[nx * n + ny] = true
                        q = append(q, []int{nx, ny, x, y}) 
                    }
                }
            }
        }

        return false
    }

    for i, row := range grid {
        for j := range row {
            if !vis[i * n + j] {
                vis[i * n + j] = true
                if bfs(i, j, row[j]) {
                    return true
                }
            }
        }
    }
    
    return false
}