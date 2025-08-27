func wordPuzzle(grid [][]byte, target string) bool {
    m, n := len(grid), len(grid[0])
    dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
    
    vis := map[int]bool{}

    var dfs func(int, int, int) bool 
    dfs = func(i, j, k int) bool {
        if k == len(target) {
            return true
        }
        
        for _, dir := range dirs {
            x, y := i + dir[0], j + dir[1]
            if x >= 0 && x < m && y >= 0 && y < n && !vis[x * n + y] && grid[x][y] == target[k] {
                vis[x * n + y] = true
                if dfs(x, y, k + 1) {
                    return true
                }
                vis[x * n + y] = false
            }
        }
        vis[i * n + j] = false
        return false
    }
    

    for i := range grid {
        for j := range grid[i] {
            if grid[i][j] == target[0] {
                vis[i * n + j] = true
                if dfs(i, j, 1) {
                    return true
                }
                vis[i * n + j] = false
            }
        }
    }
    
    return false
}
