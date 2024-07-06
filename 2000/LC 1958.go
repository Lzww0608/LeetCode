func checkMove(board [][]byte, rMove int, cMove int, color byte) bool {
    m, n := len(board), len(board[0])
    oppositeColor := byte('W')
    if color == 'W' {
        oppositeColor = 'B'
    }

    check := func(x_incre, y_incre int) bool {
        i, j := rMove + x_incre, cMove + y_incre
        foundOppositeColor := false
        for i >= 0 && i < m && j >= 0 && j < n {
            if board[i][j] == '.' {
                return false
            }
            if board[i][j] == color {
                return foundOppositeColor
            }
            if board[i][j] == oppositeColor {
                foundOppositeColor = true
            }
            i += x_incre
            j += y_incre
        }
        return false
    }

    directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
    for _, dir := range directions {
        if check(dir[0], dir[1]) {
            return true
        }
    }
    return false
}
