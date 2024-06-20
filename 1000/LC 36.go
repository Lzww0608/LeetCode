func isValidSudoku(board [][]byte) bool {
    row, col, box := [9][9]int{}, [9][9]int{}, [3][3][9]int{}
    for i, y := range board {
        for j, x := range y {
            if x == '.' {
                continue
            }
            num := int(x - '1')
            row[i][num]++
            col[j][num]++
            box[i/3][j/3][num]++
            if row[i][num] > 1 || col[j][num] > 1 || box[i/3][j/3][num] > 1 {
                return false
            }
        } 
    }
    return true
}
