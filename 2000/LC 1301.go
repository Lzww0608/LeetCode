const MOD int = 1e9 + 7
func pathsWithMaxScore(board []string) []int {
    n := len(board)
    type pair struct {
        x, y int
    }
    f := make([][]pair, n)
    for i := range f {
        f[i] = make([]pair, n)
        for j := range f[i] {
            f[i][j] = pair{-1, 0}
        }
    }
    f[n-1][n-1] = pair{0, 1}

    update := func(i, j, x, y int) {
        if x >= n || y >= n || f[x][y].x == -1 {
            return
        }

        if f[i][j].x < f[x][y].x {
            f[i][j] = f[x][y]
        } else if f[i][j].x == f[x][y].x {
            f[i][j].y += f[x][y].y
            f[i][j].y %= MOD
        }

        return 
    }

    for i := n - 1; i >= 0; i-- {
        for j := n - 1; j >= 0; j-- {
            c := board[i][j]
            if !(i == n - 1 && j == n - 1) && c != 'X' {
                update(i, j, i + 1, j)
                update(i, j, i, j + 1)
                update(i, j, i + 1, j + 1)
                if f[i][j].x != -1 && c != 'E' {
                    f[i][j].x += int(c - '0')
                }
            }


        }
    }

    if f[0][0].x == -1 {
        return []int{0, 0}
    }
    
    return []int{f[0][0].x, f[0][0].y}
}
