func sumRemoteness(grid [][]int) (ans int64) {
    dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
    sum, cnt := 0, 0
    n := len(grid)
    id := make(map[int]int)
    fa := []int{}
    sz := []int{}
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] > 0 {
                sum += grid[i][j]
                id[i * n + j] = cnt
                fa = append(fa, cnt)
                sz = append(sz, grid[i][j])
                cnt++
            }
        }
    }

    var find func(int) int
    find = func(x int) int {
        if x != fa[x] {
            fa[x] = find(fa[x])
        }
        return fa[x]
    }

    merge := func(x, y int) {
        rx, ry := find(x), find(y)
        if rx == ry {
            return
        }
        if sz[rx] > sz[ry] {
            rx, ry = ry, rx
        }
        fa[rx] = ry
        sz[ry] += sz[rx]
    }

    bfs := func(sx, sy int) (res int) {
        q := [][]int{{sx, sy}}
        for len(q) > 0 {
            cur := q[0]
            q = q[1:]
            i, j := cur[0], cur[1]
            merge(id[sx * n + sy], id[i * n + j])
            res++
            for _, dir := range dirs {
                x, y := i + dir[0], j + dir[1]
                if x >= 0 && x < n && y >= 0 && y < n && grid[x][y] > 0{
                    q = append(q, []int{x, y})
                    grid[x][y] = -1
                }
            }
        }

        return
    }

    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] > 0 {
                grid[i][j] = -1
                t := bfs(i, j)
                ans += int64((sum - sz[find(id[i * n + j])]) * t)
            }
        }
    }

    return
}