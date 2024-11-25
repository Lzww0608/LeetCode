func networkDelayTime(times [][]int, n int, k int) int {
    dis := make([]int, n)
    g := make([][]int, n)
    for i := range dis {
        g[i] = make([]int, n)
        for j := range g[i] {
            g[i][j] = math.MaxInt32 / 2
        }
        dis[i] = math.MaxInt32 / 2
    }
    dis[k-1] = 0

    
    vis := make([]bool, n)
    for _, v := range times {
        a, b := v[0] - 1, v[1] - 1
        g[a][b] = v[2]
    } 

    for {
        x := -1
        for i := 0; i < n; i++ {
            if !vis[i] && (x == -1 || dis[i] < dis[x]) {
                x = i 
            }
        }
        if x < 0 {
            return slices.Max(dis)
        } else if dis[x] == math.MaxInt32 / 2 {
            return -1
        }

        vis[x] = true
        for i := 0; i < n; i++ {
            dis[i] = min(dis[i], dis[x] + g[x][i])
        }
    }
}