func flipChess(chessboard []string) (ans int) {
    m, n := len(chessboard), len(chessboard[0])
    dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}
       
    bfs := func(i, j int) (cnt int) {
        vis := make([][]bool, m)
        for i := range vis {
            vis[i] = make([]bool, n)
        }
        
        board := make([][]byte, m)
        for i := range board {
            board[i] = []byte(chessboard[i])
        }
        board[i][j] = 'X'

        q := [][]int{{i, j}}
        
        for len(q) > 0 {
            x, y := q[0][0], q[0][1]
            q = q[1:]
            
            for _, dir := range dirs {
                temp := [][]int{}
                nx, ny := x + dir[0], y + dir[1]
                for nx >= 0 && nx < m && ny >= 0 && ny < n && board[nx][ny] == 'O' {
                    temp = append(temp, []int{nx, ny})
                    nx += dir[0]
                    ny += dir[1]
                }

                if nx >= 0 && nx < m && ny >= 0 && ny < n && board[nx][ny] == 'X' {
                    cnt += len(temp)
                    for _, pos := range temp {
                        q = append(q, pos)
                        board[pos[0]][pos[1]] = 'X'
                    }
                }
            }
        }
        return cnt
    }

    for i := range chessboard {
        for j := range chessboard[i] {
            if chessboard[i][j] == '.' {
                ans = max(ans, bfs(i, j))
            }
        }
    }

    return
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
