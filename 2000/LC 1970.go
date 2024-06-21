var dirs = [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
func latestDayToCross(row int, col int, cells [][]int) int {
    top := row * col
    bottom := top + 1
    fa := make([]int ,bottom + 1)
    for i := range fa {
        fa[i] = i
    }

    m := map[int]bool{}

    var find func(int) int
    find = func(x int) int {
        if x != fa[x] {
            fa[x] = find(fa[x])
        }
        return fa[x]
    }

    merge := func(x, y int) {
        rx, ry := find(x), find(y)
        if rx != ry {
            fa[rx] = y
        }
    }

    for i := len(cells) - 1; i >= 0; i-- {
        v := cells[i]
        r, c := v[0] - 1, v[1] - 1
        pos := r * col + c
        for _, dir := range dirs {
            x, y := r + dir[0], c + dir[1]
            if x >= 0 && x < row && y >= 0 && y < col && m[x * col + y] {
                merge(pos, x * col + y)
            }
        }

        if r == 0 {
            merge(pos, top)
        } else if r == row - 1 {
            merge(pos, bottom)
        }

        if find(top) == find(bottom) {
            return i
        }

        m[r * col + c] = true
    }

    return 0
}
