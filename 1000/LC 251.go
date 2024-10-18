type Vector2D struct {
    a []int
    p int
}


func Constructor(vec [][]int) Vector2D {
    a := make([]int, 0, 10000)
    for _, v := range vec {
        for _, x := range v {
            a = append(a, x)
        }
    }

    return Vector2D{a, 0}
}


func (c *Vector2D) Next() int {
    c.p++
    return c.a[c.p-1]
}


func (c *Vector2D) HasNext() bool {
    return c.p < len(c.a)
}


/**
 * Your Vector2D object will be instantiated and called as such:
 * obj := Constructor(vec);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */