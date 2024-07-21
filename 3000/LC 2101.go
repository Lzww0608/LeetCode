func maximumDetonation(bombs [][]int) (ans int) {
    n := len(bombs)
    g := make([][]int, n)
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if i == j {
                continue
            }
            if inRange(bombs[i], bombs[j]) {
                g[i] = append(g[i], j)
            }
        }
    }
    
    vis := make([]bool, n)
    var dfs func(int) int
    dfs = func(i int) int {
        vis[i] = true
        cnt := 1
        for _, v := range g[i] {
            if !vis[v] {
                cnt += dfs(v)
            }
        }
        return cnt
    }
    
    for i := 0; i < n; i++ {
        clear(vis)
        ans = max(ans, dfs(i))
    }
    
    return
}

func inRange(bomb1, bomb2 []int) bool {
    dx, dy := bomb1[0] - bomb2[0], bomb1[1] - bomb2[1]
    distanceSq := dx*dx + dy*dy
    radiusSq := bomb1[2] * bomb1[2]
    return distanceSq <= radiusSq
}

