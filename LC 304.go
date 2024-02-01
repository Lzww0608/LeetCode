type NumMatrix struct {
    pre [][]int
}


func Constructor(matrix [][]int) NumMatrix {
    n, m := len(matrix), len(matrix[0])
    pre := make([][]int, n + 1)
    pre[0] = make([]int, m + 1)
    for i := 0; i < n; i++ {
        pre[i+1] = make([]int, m + 1)
        for j := 0; j < m; j++ {
            pre[i+1][j+1] = pre[i][j+1] + pre[i+1][j] - pre[i][j] + matrix[i][j]
        }
    }
    return NumMatrix{pre}
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
    return this.pre[row2+1][col2+1] + this.pre[row1][col1] - this.pre[row1][col2+1] - this.pre[row2+1][col1]
}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */