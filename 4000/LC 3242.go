type pair struct {
    x, y int
}
type NeighborSum struct {
    n int
    m map[int]pair
    mat [][]int
}


func Constructor(grid [][]int) NeighborSum {
    n := len(grid)
    m := map[int]pair{}
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
           m[grid[i][j]] = pair{i, j} 
        }
    }

    return NeighborSum{n, m, grid}
}


func (c *NeighborSum) AdjacentSum(value int) (ans int) {
    i, j := c.m[value].x, c.m[value].y 
    if i - 1 >= 0 {
        ans += c.mat[i-1][j]
    }
    if i + 1 < c.n {
        ans += c.mat[i+1][j]
    }
    if j - 1 >= 0 {
        ans += c.mat[i][j-1]
    }
    if j + 1 < c.n {
        ans += c.mat[i][j+1]
    }
    return 
}


func (c *NeighborSum) DiagonalSum(value int) (ans int) {
    i, j := c.m[value].x, c.m[value].y 
    if i - 1 >= 0 && j - 1 >= 0 {
        ans += c.mat[i-1][j-1]
    }
    if i - 1 >= 0 && j + 1 < c.n  {
        ans += c.mat[i-1][j+1]
    }
    if i + 1 < c.n && j - 1 >= 0 {
        ans += c.mat[i+1][j-1]
    }
    if i + 1 < c.n && j + 1 < c.n  {
        ans += c.mat[i+1][j+1]
    }
    return 
}


/**
 * Your NeighborSum object will be instantiated and called as such:
 * obj := Constructor(grid);
 * param_1 := obj.AdjacentSum(value);
 * param_2 := obj.DiagonalSum(value);
 */