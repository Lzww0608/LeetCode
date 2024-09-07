func numIslands2(m int, n int, positions [][]int) []int {
    dirs := [][]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}

    cnt := 0
    fa := make([]int, m * n)
    g := make([][]int, m)
    for i := range g {
        g[i] = make([]int, n)
    }

    for i := range fa {
        fa[i] = i
    }

    var find func(int) int 
    find = func(x int) int {
        if fa[x] != x {
            fa[x] = find(fa[x])
        }

        return fa[x]
    }

    merge := func(x, y int) {
        rx, ry := find(x), find(y)
        if rx != ry {
            fa[ry] = rx
            cnt--
        }
    }

    ans := make([]int, len(positions))
    for i, v := range positions {
        x, y := v[0], v[1]
        if g[x][y] == 1 {
            ans[i] = cnt
            continue
        }
        g[x][y] = 1
        cnt++
        for _, dir := range dirs {
            nx, ny := x + dir[0],  y + dir[1]
            if nx >= 0 && nx < m && ny >= 0 && ny < n && g[nx][ny] == 1 {
                merge(x * n + y, nx * n + ny)
            }
        }

        ans[i] = cnt
    }

    return ans
}