func updateBoard(board [][]byte, click []int) [][]byte {
    m, n := len(board), len(board[0])
    i, j := click[0], click[1]

    if board[i][j] == 'M' {
        board[i][j] = 'X'
        return board
    }

    q := []int{i * n + j}
    vis := make([]bool, m * n)
    vis[i * n + j] = true

    calc := func(i, j int) (int, []int) {
        cnt := 0
        p := []int{}
        for x := -1; x <= 1; x++ {
            for y := -1; y <= 1; y++ {
                nx, ny := i + x, j + y 
                if nx >= 0 && nx < m && ny >= 0 && ny < n {
                    if board[nx][ny] == 'X' || board[nx][ny] == 'M' {
                        cnt++
                    } else if cnt == 0 && !vis[nx * n + ny] {
                        p = append(p, nx * n + ny)
                    }
                }
            }
        }

        return cnt, p
    }

    for len(q) > 0 {
        cur := q[0]
        q = q[1:]
        i, j = cur / n, cur % n
        cnt, p := calc(i, j)
        if cnt != 0 {
            board[i][j] = byte(cnt + '0')
        } else {
            board[i][j] = 'B'
            for _, x := range p {
                if !vis[x] {
                    q = append(q, x)
                    vis[x] = true
                }
            }
        }
    }

    return board
}