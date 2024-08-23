func movesToChessboard(board [][]int) int {
    n := len(board)
    for i := range board {
        for j := range board {
            if board[0][0] ^ board[i][0] ^ board[0][j] ^ board[i][j] != 0 {
                return -1
            }
        }
    }

    var colCnt, rowCnt, colSwap, rowSwap int
    for i := range board {
        rowCnt += board[i][0]
        colCnt += board[0][i]

        if board[i][0] == i & 1 {
            rowSwap++
        }
        if board[0][i] == i & 1 {
            colSwap++
        }
    }

    if rowCnt != n / 2 && rowCnt != (n + 1) / 2 {
        return -1
    } else if colCnt != n / 2 && colCnt != (n + 1) / 2 {
        return -1
    }

    if n & 1 == 1 {
        if colSwap & 1 == 1 {
            colSwap = n - colSwap
        }
        if rowSwap & 1 == 1 {
            rowSwap = n - rowSwap
        }
    } else {
        rowSwap = min(rowSwap, n - rowSwap)
        colSwap = min(colSwap, n - colSwap)
    }

    return (rowSwap + colSwap) / 2
}