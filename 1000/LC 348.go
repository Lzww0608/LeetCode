type TicTacToe struct {
    row, col []int
    dg, udg, n int
}


func Constructor(n int) TicTacToe {
    row := make([]int, n)
    col := make([]int, n)
    return TicTacToe{row, col, 0, 0, n}
}


func (c *TicTacToe) Move(i int, j int, player int) int {
    x := 1
    if player == 2 {
        x = -1
    }

    c.row[i] += x
    c.col[j] += x
    if i == j {
        c.dg += x
    }
    if i + j == c.n - 1 {
        c.udg += x
    }

    if c.row[i] == c.n || c.col[j] == c.n || c.dg == c.n || c.udg == c.n || c.row[i] == -c.n || c.col[j] == -c.n || c.dg == -c.n || c.udg == -c.n {
        return player
    }

    return 0
}


/**
 * Your TicTacToe object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Move(row,col,player);
 */