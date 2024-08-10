func maximumSafenessFactor(grid [][]int) int {
    n := len(grid)
    q := [][]int{}
    dis := make([][]int, n)

    for i, row := range grid {
        dis[i] = make([]int, n)
        for j, x := range row {
            if x == 1 {
                q = append(q, []int{i, j})
            } else {
                dis[i][j] = -1
            }
        }
    }

    dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
    groups := [][][]int{q}
    for len(q) > 0 {
        tmp := q
        q = nil 
        for _, v := range tmp {
            for _, dir := range dirs {
                x, y := v[0] + dir[0], v[1] + dir[1]
                if x >= 0 && x < n && y >= 0 && y < n && dis[x][y] < 0 {
                    q = append(q, []int{x, y})
                    dis[x][y] = len(groups)
                } 
            }
        }

        groups = append(groups, q)
    }

    fa := make([]int, n * n)
    for i := range fa {
        fa[i] = i
    }

    var find func(int) int 
    find = func(x int) int {
        if x != fa[x] {
            fa[x] = find(fa[x])
        }

        return fa[x]
    }

    for ans := len(groups) - 2; ans > 0; ans-- {
        for _, v := range groups[ans] {
            i, j := v[0], v[1]
            for _, dir := range dirs {
                x, y := i + dir[0], j + dir[1]
                if x >= 0 && x < n && y >= 0 && y < n && dis[x][y] >= dis[i][j] {
                    fa[find(x * n + y)] = find(i * n + j)
                }
            }
        }

        if find(0) == find(n * n - 1) {
            return ans
        }
    }

    return 0
}