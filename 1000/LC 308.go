type NumMatrix struct {
    m, n int
    pre [][]int
    matrix [][]int
}


func Constructor(matrix [][]int) NumMatrix {
    m, n := len(matrix), len(matrix[0])
    pre := make([][]int, m)
    for i := 0; i < m; i++ {
        pre[i] = make([]int, n + 1)
        for j := 0; j < n; j++ {
            pre[i][j+1] = pre[i][j] + matrix[i][j]
        }
    }

    return NumMatrix{m, n, pre, matrix}
}


func (c *NumMatrix) Update(row int, col int, val int)  {
    dif := val - c.matrix[row][col]
    c.matrix[row][col] = val
    for j := col + 1; j <= c.n; j++ {
        c.pre[row][j] += dif
    }
}


func (c *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) (ans int) {
    for i := row1; i <= row2; i++ {
        ans += c.pre[i][col2+1] - c.pre[i][col1]
    }

    return
}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * obj.Update(row,col,val);
 * param_2 := obj.SumRegion(row1,col1,row2,col2);
 */