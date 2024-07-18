func tictactoe(board []string) string {
    n := len(board)
    is_full := true
    var xDia, xuDia, oDia, ouDia bool

    for i, row := range board {
        var xRow, xCol, oRow, oCOl bool
        for j := range row {
            if is_full && board[i][j] == ' ' {
                is_full = false
            }
            if !xRow {
                xRow = xRow || ('X' != board[i][j])
            }
            if !oRow {
                oRow = oRow || ('O' != board[i][j])
            }
            if !xCol {
                xCol = xCol || ('X' != board[j][i])
            }
            if !oCOl {
                oCOl = oCOl || ('O' != board[j][i])
            }
        }
        if !xRow || !xCol {
            return "X"
        }
        if !oRow || !oCOl {
            return "O"
        }

        if !xDia {
            xDia = xDia || ('X' != board[i][i])
        }
        if !xuDia {
            xuDia = xuDia || ('X' != board[i][n-i-1])
        }
        if !oDia {
            oDia = oDia || ('O' != board[i][i])
        }
        if !ouDia {
            ouDia = ouDia || ('O' != board[i][n-i-1])
        }
    }
    if !xDia || !xuDia {
        return "X"
    }
    if !oDia || !ouDia {
        return "O"
    }
    if is_full {
        return "Draw"
    }
    return "Pending"
}