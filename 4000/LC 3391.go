type Matrix3D struct {
    m []map[int]bool
    n int
}


func Constructor(n int) Matrix3D {
    m := make([]map[int]bool, n)
    for i := range m {
        m[i] = make(map[int]bool)
    }
    return Matrix3D{m, n}
}


func (c *Matrix3D) SetCell(x int, y int, z int)  {
    c.m[x][y * c.n + z] = true
}


func (c *Matrix3D) UnsetCell(x int, y int, z int)  {
    delete(c.m[x], y * c.n + z)
}


func (c *Matrix3D) LargestMatrix() int {
    ans, mx := 0, 0
    for i := 0; i < c.n; i++ {
        if mx <= len(c.m[i]) {
            mx = len(c.m[i])
            ans = i
        }
    }

    return ans
}


/**
 * Your Matrix3D object will be instantiated and called as such:
 * obj := Constructor(n);
 * obj.SetCell(x,y,z);
 * obj.UnsetCell(x,y,z);
 * param_3 := obj.LargestMatrix();
 */