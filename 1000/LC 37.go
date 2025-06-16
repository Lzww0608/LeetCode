
const N int = 9
func solveSudoku(board [][]byte)  {
    row := make([]int, N)
    col := make([]int, N)
    cell := make([]int, N)
    for i := range board {
        for j, c := range board[i] {
            if c != '.' {
                x := int(c - '1')
                row[i] |= 1 << x 
                col[j] |= 1 << x 
                pos := i / 3 * 3 + j / 3 
                cell[pos] |= 1 << x
            }
        }
    }
    dfs(board, row, col, cell)
}

func dfs(a [][]byte, row, col, cell []int) bool {
    for i := 0; i < N; i++ {
        for j := 0; j < N; j++ {
            if a[i][j] != '.' {
                continue
            }
            for k := 0; k < N; k++ {
                if row[i] & (1 << k) != 0 || col[j] & (1 << k) != 0 {
                    continue
                }
                pos := i / 3 * 3 + j / 3
                if cell[pos] & (1 << k) != 0 {
                    continue
                }
                a[i][j] = byte('0' + k + 1)
                row[i] |= 1 << k 
                col[j] |= 1 << k 
                cell[pos] |= 1 << k
                if dfs(a, row, col, cell) {
                    return true
                }
                a[i][j] = '.'
                row[i] &^= 1 << k
                col[j] &^= 1 << k 
                cell[pos] &^= 1 << k
            }
            return false
        }
    }

    return true
}