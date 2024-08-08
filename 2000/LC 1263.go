import "container/list"
func minPushBox(grid [][]byte) int {
    m, n := len(grid), len(grid[0])

    f := func(i, j int) int {
        return i * n + j
    }

    check := func(i, j int) bool {
        return i >= 0 && i < m && j >= 0 && j < n && grid[i][j] != '#'
    }

    vis := make([][]bool, m * n)
    for i := range vis {
        vis[i] = make([]bool, m * n)
    }

    var si, sj, bi, bj int
    for i, row := range grid {
        for j, x := range row {
            if x == 'B' {
                bi, bj = i, j
            } else if x == 'S' {
                si, sj = i, j
            }
        }
    }

    dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
    q := list.New()
    q.PushFront([]int{f(si, sj), f(bi, bj), 0})
    vis[f(si, sj)][f(bi, bj)] = true
    for q.Len() > 0 {
        node := q.Remove(q.Front()).([]int)
        s, b, d := node[0], node[1], node[2]
        si, sj = s / n, s % n 
        bi, bj = b / n, b % n
        if grid[bi][bj] == 'T' {
            return d
        }
        for _, dir := range dirs {
            sx, sy := si + dir[0], sj + dir[1]
            if !check(sx, sy) {
                continue
            }

            if sx == bi && sy == bj {
                bx, by := bi + dir[0], bj + dir[1]
                if !check(bx, by) || vis[f(sx, sy)][f(bx, by)] {
                    continue
                }
                vis[f(sx, sy)][f(bx, by)] = true
                q.PushBack([]int{f(sx, sy), f(bx, by), d + 1})
            } else {
                if vis[f(sx, sy)][f(bi, bj)] {
                    continue
                }
                vis[f(sx, sy)][f(bi, bj)] = true
                q.PushFront([]int{f(sx, sy), f(bi, bj), d})
            }
        }
    }

    return -1
}