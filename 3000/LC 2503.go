func maxPoints(grid [][]int, queries []int) []int {
    dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
    m, n := len(grid), len(grid[0])
    k := len(queries)
    ans := make([]int, k)
    id := make([]int, k)
    for i := range id {
        id[i] = i
    }
    sort.Slice(id, func(i, j int) bool {
        return queries[id[i]] < queries[id[j]]
    })

    pos := make([]int, m * n)
    fa := make([]int, m * n)
    sz := make([]int, m *n )
    for i := range pos {
        pos[i] = i
        fa[i] = i
        sz[i] = 1
    }
    sort.Slice(pos, func(i, j int) bool {
        x1, y1 := pos[i] / n, pos[i] % n 
        x2, y2 := pos[j] / n, pos[j] % n 
        return grid[x1][y1] < grid[x2][y2]
    })
    //fmt.Println(pos)

    var find func(int)int
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
        if sz[rx] < sz[ry] {
            rx, ry = ry, rx
        }
        fa[ry] = rx
        sz[rx] += sz[ry]
    }

    j := 0
    for _, i := range id {
        for j < m * n && grid[pos[j]/n][pos[j]%n] < queries[i] {
            for _, dir := range dirs {
                x, y := pos[j] / n + dir[0], pos[j] % n + dir[1]
                if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] < queries[i] {
                    merge(pos[j], x * n + y)
                }
            }
            j++
        }

        if grid[0][0] < queries[i] {
            ans[i] = sz[find(0)]
        }
    }

    return ans
}