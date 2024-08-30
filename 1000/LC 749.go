func containVirus(isInfected [][]int) (ans int) {
    m, n := len(isInfected), len(isInfected[0])
    vis := make([]bool, m * n)
    dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
    areas := []map[int]bool{}
    boundaries := []map[int]bool{}
    c := []int{}
    getMax := func() (ans int) {
        mx := len(boundaries[0])
        for i, row := range boundaries {
            if len(row) > mx {
                mx = len(row)
                ans = i
            } 
        }

        return 
    }

    var dfs func(int, int) 
    dfs = func(i, j int) {
        areas[len(areas)-1][i * n + j] = true

        for _, dir := range dirs {
            x, y := i + dir[0], j + dir[1]
            if x >= 0 && x < m && y >= 0 && y < n {
                if isInfected[x][y] == 1 && !vis[x * n + y] {
                    vis[x * n + y] = true
                    dfs(x, y)
                } else if isInfected[x][y] == 0 {
                    c[len(c)-1] += 1
                    boundaries[len(boundaries)-1][x * n + y] = true
                }
            }
        }
    }

    for {
        areas = []map[int]bool{}
        boundaries = []map[int]bool{}
        c = []int{}
        clear(vis)
        for i := 0; i < m; i++ {
            for j := 0; j < n; j++ {
                if isInfected[i][j] == 1 && !vis[i * n + j] {
                    vis[i * n + j] = true
                    c = append(c, 0)
                    boundaries = append(boundaries, map[int]bool{})
                    areas = append(areas, map[int]bool{})
                    dfs(i, j)
                }
            }
        }

        if len(areas) == 0 {
            break
        }
        idx := getMax()
        ans += c[idx]
        for k, area := range areas {
            if k == idx {
                for v := range area {
                    i, j := v / n, v % n
                    isInfected[i][j] = -1
                }
            } else {
                for v := range area {
                    i, j := v / n, v % n
                    for _, dir := range dirs {
                        x, y := i + dir[0], j + dir[1]
                        if x >= 0 && x < m && y >= 0 && y < n && isInfected[x][y] == 0 {
                            isInfected[x][y] = 1
                        }
                    }
                }
            }
        }
    }

    return
}