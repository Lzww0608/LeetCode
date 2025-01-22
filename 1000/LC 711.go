func numDistinctIslands2(grid [][]int) int {
    dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
    m, n := len(grid), len(grid[0])
    vis := make([]bool, m * n)
    a := []island{}
    
    bfs := func(sx, sy, col int) island {
        tmp := island{100, 100, 0, 0, 0, col}
        q := [][]int{{sx, sy}}
        vis[sx * n + sy] = true
        for len(q) > 0 {
            cur := q[0]
            i, j := cur[0], cur[1]
            q = q[1:]
            tmp.x1 = min(tmp.x1, i)
            tmp.y1 = min(tmp.y1, j)
            tmp.x2 = max(tmp.x2, i)
            tmp.y2 = max(tmp.y2, j)
            tmp.cnt++
            grid[i][j] = col
            for _, dir := range dirs {
                x, y := i + dir[0], j + dir[1]
                if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == 1 && !vis[x * n + y] {
                    vis[x * n + y] = true
                    q = append(q, []int{x, y})
                }
            }
        }

        return tmp
    }

    check := func(tmp island) bool {
        dx1 := tmp.x2 - tmp.x1
        dy1 := tmp.y2 - tmp.y1
        c := tmp.col
        for _, v := range a {
            if v.cnt != tmp.cnt {
                continue
            }
            cc := v.col
            dx2 := v.x2 - v.x1
            dy2 := v.y2 - v.y1
            if dx1 == dx2 && dy1 == dy2 {
                f := true
                for i := 0; i <= dx1 && f; i++ {
                    for j := 0; j <= dy1; j++ {
                        if grid[tmp.x1+i][tmp.y1+j] == c && grid[v.x1+i][v.y1+j] != cc {   
                            f = false
                            break
                        }
                    }
                }
                if f {
                    return true
                }
                f = true
                for i := 0; i <= dx1 && f; i++ {
                    for j := 0; j <= dy1; j++ {
                        if grid[tmp.x1+i][tmp.y1+j] == c && grid[v.x1+i][v.y2-j] != cc {   
                            f = false
                            break
                        }
                    }
                }
                if f {
                    return true
                }
                f = true
                for i := 0; i <= dx1 && f; i++ {
                    for j := 0; j <= dy1; j++ {
                        if grid[tmp.x1+i][tmp.y1+j] == c && grid[v.x2-i][v.y1+j] != cc {   
                            f = false
                            break
                        }
                    }
                }
                if f {
                    return true
                }
                f = true
                for i := 0; i <= dx1 && f; i++ {
                    for j := 0; j <= dy1; j++ {
                        if grid[tmp.x1+i][tmp.y1+j] == c && grid[v.x2-i][v.y2-j] != cc {   
                            f = false
                            break
                        }
                    }
                }
                if f {
                    return true
                }
            } 
            if dx1 == dy2 && dy1 == dx2 {
                f := true
                for i := 0; i <= dx1 && f; i++ {
                    for j := 0; j <= dy1; j++ {
                        if grid[tmp.x1+i][tmp.y1+j] == c && grid[v.x1+j][v.y1+i] != cc {   
                            f = false
                            break
                        }
                    }
                }
                if f {
                    return true
                }
                f = true
                for i := 0; i <= dx1 && f; i++ {
                    for j := 0; j <= dy1; j++ {
                        if grid[tmp.x1+i][tmp.y1+j] == c && grid[v.x1+j][v.y2-i] != cc {   
                            f = false
                            break
                        }
                    }
                }
                if f {
                    return true
                }
                f = true
                for i := 0; i <= dx1 && f; i++ {
                    for j := 0; j <= dy1; j++ {
                        if grid[tmp.x1+i][tmp.y1+j] == c && grid[v.x2-j][v.y1+i] != cc {   
                            f = false
                            break
                        }
                    }
                }
                if f {
                    return true
                }
                f = true
                for i := 0; i <= dx1 && f; i++ {
                    for j := 0; j <= dy1; j++ {
                        if grid[tmp.x1+i][tmp.y1+j] == c && grid[v.x2-j][v.y2-i] != cc {   
                            f = false
                            break
                        }
                    }
                }
                if f {
                    return true
                }
            }
        }
        return false
    }

    col := 1
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if !vis[i * n + j] && grid[i][j] == 1 {
                tmp := bfs(i, j, col)
                col++
                if !check(tmp) {
                    a = append(a, tmp)
                }
            }
        }
    }

    return len(a)
}

type island struct {
    x1, y1, x2, y2 int
    cnt int
    col int
}