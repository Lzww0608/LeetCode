func constructGridLayout(n int, edges [][]int) [][]int {
    g := make([][]int, n)
    for _, edge := range edges {
        u, v := edge[0], edge[1]
        g[u] = append(g[u], v)
        g[v] = append(g[v], u)
    }

    deg := [5]int{-1, -1, -1, -1, -1}
    for k, v := range g {
        deg[len(v)] = k
    }

    var row []int
    if deg[1] != -1 {
        // 只有一列
        row = append(row, deg[1])
    } else if deg[4] == -1 {
        // 只有两列
        x := deg[2]
        for _, y := range g[x] {
            if len(g[y]) == 2 {
                row = append(row, y)
                break
            }
        }
        row = append(row, x)
    } else {
        // 至少3列，每一行必然是2333...2的序列
        x := deg[2]
        row = append(row, x)
        pre := x 
        x = g[x][0]
        for len(g[x]) == 3 {
            row = append(row, x)
            for _, y := range g[x] {
                if y != pre && len(g[y]) < 4 {
                    pre = x 
                    x = y
                    break
                }
            }
        }
        row = append(row, x)
    }

    k := len(row)
    ans := make([][]int, n / k)
    ans[0] = row
    vis := make([]bool, n)
    for _, x := range row {
        vis[x] = true
    }

    for i := 1; i < len(ans); i++ {
        ans[i] = make([]int, k)
        for j, x := range ans[i-1] {
            for _, y := range g[x] {
                if !vis[y] {
                    vis[y] = true
                    ans[i][j] = y 
                    break
                }
            }
        }
    }

    return ans
}