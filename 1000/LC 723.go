func candyCrush(board [][]int) [][]int {
    m, n := len(board), len(board[0])
    for {
        f := false
        for i := 0; i < m; i++ {
            for j := 0; j < n - 2; j++ {
                if board[i][j] != 0 && abs(board[i][j]) == abs(board[i][j+1]) && abs(board[i][j]) == abs(board[i][j+2]) {
                    x := abs(board[i][j])
                    board[i][j], board[i][j+1], board[i][j+2] = -x, -x, -x
                    f = true
                }
            }
        }

        for j := 0; j < n; j++ {
            for i := 0; i < m - 2; i++ {
                if board[i][j] != 0 && abs(board[i][j]) == abs(board[i+1][j]) && abs(board[i+2][j]) == abs(board[i][j]) {
                    x := abs(board[i][j])
                    board[i][j], board[i+1][j], board[i+2][j] = -x, -x, -x
                    f = true
                }
            }
        }

        if !f {
            break
        }

        for j := 0; j < n; j++ {
            p, q := m - 1, m - 1
            for p >= 0 {
                if board[p][j] > 0 {
                    board[q][j] = board[p][j]
                    q--
                } 
                p--
            }
            for q >= 0 {
                board[q][j] = 0 
                q--
            }
        }
    }

    return board
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}