func numRookCaptures(board [][]byte) (ans int) {
    const N = 8
    x, y := 0, 0
    dirs := [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
    for i := range board {
        for j := range board[i] {
            if board[i][j] == 'R' {
                x, y = i, j
                for _, dir := range dirs {
                    i, j := x, y
                    for i >= 0 && i < N && j >= 0 && j < N && board[i][j] != 'B' {
                        if board[i][j] == 'p' {
                            ans++
                            break
                        }
                        i, j = i + dir[0], j + dir[1]
                    }
                }
                return 
            }
        }
    }

    return
}
