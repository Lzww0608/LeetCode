func countBattleships(board [][]byte) (ans int) {
    for i := range board {
        for j := range board[i] {
            if i > 0 && board[i][j] == board[i-1][j] {
                continue
            } else if j > 0 && board[i][j] == board[i][j-1] {
                continue
            } else if board[i][j] == 'X' {
                ans++
            }
        }
    }

    return
}