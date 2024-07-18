func minimumMoves(grid [][]int) int {
    n := len(grid)
    dirs := [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
    vis := make([][][2]bool, n)
    for i := range vis {
        vis[i] = make([][2]bool, n)
    }

    q := [][]int{{0, 0, 0}}
    for d := 1; len(q) > 0; d++ {
        for t := len(q); t > 0; t-- {
            node := q[0]
            q = q[1:]
            for _, dir := range dirs {
                x, y, z := node[0] + dir[0], node[1] + dir[1], node[2] ^ dir[2]
                i, j := x + z, y + (z ^ 1)
                if i < n && j < n && grid[x][y] == 0 && grid[i][j] == 0 && !vis[x][y][z] &&
                    (dir[2] == 0 || grid[x+1][y+1] == 0) {
                        if x == n - 1 && y == n - 2 {
                            return d
                        }
                        vis[x][y][z] = true
                        q = append(q, []int{x, y, z})
                    }
            }
        }
        
    }

    return -1
}