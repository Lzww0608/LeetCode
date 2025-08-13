const N int = 3
func validTicTacToe(board []string) bool {
    hasWin := func(c byte) bool {
        for i := 0; i < N; i++ {
            cnt := 0
            for j := 0; j < N; j++ {
                if board[i][j] == c {
                    cnt++
                }
            }
            if cnt == N {
                return true 
            }
        }

        for j := 0; j < N; j++ {
            cnt := 0
            for i := 0; i < N; i++ {
                if board[i][j] == c {
                    cnt++
                }
            }
            if cnt == N {
                return true
            }
        }
        if board[0][0] == c && board[1][1] == c && board[2][2] == c {
            return true
        }
        if board[0][2] == c && board[1][1] == c && board[2][0] == c {
            return true
        }

        return false
    }

    cntX, cntO := 0, 0 
    for i := 0; i < N; i++ {
        for j := 0; j < N; j++ {
            if board[i][j] == 'X' {
                cntX++
            } else if board[i][j] == 'O' {
                cntO++
            }
        }
    }

    if cntX < cntO || cntX > cntO + 1 {
        return false 
    }

    winX := hasWin('X')
    winO := hasWin('O')

    if winX && winO {
        return false
    }

    if winX && cntX != cntO + 1 || winO && cntX != cntO {
        return false
    }

    return true
}