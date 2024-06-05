func gameOfLife(board [][]int)  {
    m, n := len(board), len(board[0])

    dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
    for i := range board {
        for j := range board[i] {
            cnt := 0
            for _, dir := range dirs {
                x, y := i + dir[0], j +dir[1]
                if x >= 0 && x < m && y >= 0 && y < n {
                    if board[x][y] == 1  || board[x][y] == 2 {
                        cnt++
                    }
                }
            }
            if board[i][j] == 0 && cnt == 3 {
                board[i][j] = 3
            } else if board[i][j] == 1 && (cnt < 2 || cnt > 3) {
                board[i][j] = 2
            }
        }
    }

    for i := range board {
        for j, x := range board[i] {
            if x == 3 {
                board[i][j] = 1
            } else if x == 2 {
                board[i][j] = 0
            }
        }
    }

    return
}