var MOD int = int(1e9 + 7)


func countPaths(n int, roads [][]int) int {
    g := make([][]int, n) // 邻接矩阵
	for i := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
			g[i][j] = math.MaxInt / 2 
		}
	}
	for _, r := range roads {
		x, y, d := r[0], r[1], r[2]
		g[x][y] = d
		g[y][x] = d
	}

    dis := make([]int, n)
    for i := range dis {
        dis[i] = math.MaxInt / 2
    }
    dis[0] = 0

    f, done := make([]int, n), make([]bool, n)
    f[0] = 1
    for true {
        x := -1
        for i := 0; i < n; i++ {
            if !done[i] && (x == -1 || dis[i] < dis[x]) {
                x = i
            }
        }

        if x == n - 1 {
            return f[n-1]
        } 

        done[x] = true
        for u, v := range g[x] {
            new_dis := dis[x] + v
            if new_dis < dis[u] {
                dis[u] = new_dis
                f[u] = f[x]
            } else if new_dis == dis[u] {
                f[u] = (f[u] + f[x]) % MOD
            }
        }
    }
    return f[n-1]
}