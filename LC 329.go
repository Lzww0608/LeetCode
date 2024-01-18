var dir = [][]int{[]int{0,1}, []int{0,-1}, []int{1,0}, []int{-1,0}}
func longestIncreasingPath(matrix [][]int) int {
    n, m := len(matrix), len(matrix[0])
    vis := make([][]int, n)
    for i := range vis {
        vis[i] = make([]int, m)
    }
    var dfs func(i, j int) int 
    dfs = func(i, j int) int {
        if vis[i][j] > 0 {
            return vis[i][j]
        }
        res := 1
        for k := 0; k < 4; k++ {
            x, y := i + dir[k][0], j + dir[k][1]
            if x >= 0 && x < n && y >= 0 && y < m && matrix[x][y] > matrix[i][j] {
                res = max(res, 1 + dfs(x, y))
            }
        }
        vis[i][j] = res
        return res
    }
    ans := 0
    for i := range matrix {
        for j := range matrix[i] {
            if vis[i][j] == 0 {
                ans = max(ans, dfs(i, j))    
            } 
        }
    }
    return ans
}