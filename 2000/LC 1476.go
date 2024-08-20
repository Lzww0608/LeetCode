type SubrectangleQueries struct {
    a, h [][]int
}


func Constructor(rectangle [][]int) SubrectangleQueries {
    return SubrectangleQueries{rectangle, [][]int{}}
}


func (c *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int)  {
    c.h = append(c.h, []int{row1, col1, row2, col2, newValue})
}


func (c *SubrectangleQueries) GetValue(row int, col int) int {
    n := len(c.h)
    for i := n - 1; i >= 0; i-- {
        if c.h[i][0] <= row && c.h[i][2] >= row && c.h[i][1] <= col && c.h[i][3] >= col {
            return c.h[i][4]
        }
    }

    return c.a[row][col]
}


/**
 * Your SubrectangleQueries object will be instantiated and called as such:
 * obj := Constructor(rectangle);
 * obj.UpdateSubrectangle(row1,col1,row2,col2,newValue);
 * param_2 := obj.GetValue(row,col);
 */