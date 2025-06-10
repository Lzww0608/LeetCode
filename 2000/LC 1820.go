func maximumInvitations(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    cx := make([]int, m)
    cy := make([]int, n)
    g := make([][]int, m)
    vis := make([]bool, n)
    for i := range cx {
        cx[i] = -1
    }
    for i := range cy {
        cy[i] = -1
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                g[i] = append(g[i], j)
            }
        }
    }

    var dfs func(int) int 
    dfs = func(x int) int {
        for i := 0; i < len(g[x]); i++ {
            y := g[x][i]
            if !vis[y] {
                vis[y] = true
                if cy[y] == -1 || dfs(cy[y]) != 0 {
                    cx[x] = y
                    cy[y] = x
                    return 1
                }
            }
        }
        return 0
    }

    ans := 0
    for i := 0; i < m; i++{
        if cx[i] == -1 {
            for j := 0; j < n; j++ {
                vis[j] = false
            }
            ans += dfs(i)
        }
    }

    return ans
}