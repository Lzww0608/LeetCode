func pathWithObstacles(obstacleGrid [][]int) (ans [][]int) {
    m, n := len(obstacleGrid), len(obstacleGrid[0])
    path := [][]int{}
    vis := map[int]bool{}

    var dfs func(int, int) bool
    dfs = func(i, j int) bool {
        if i < 0 || i >= m || j < 0 || j >= n || obstacleGrid[i][j] == 1 || vis[i * n + j] {
            return false
        } else if ans != nil {
            return true
        } else if i == m - 1 && j == n - 1 {
            path = append(path, []int{i, j})
            ans = path
            return true
        } 

        path = append(path, []int{i, j}) 
        vis[i * n + j] = true

        if dfs(i + 1, j) {
            return true
        } else if dfs(i, j + 1) {
            return true
        }
        path = path[:len(path)-1]
        return false
    }
    dfs(0, 0)

    return
}