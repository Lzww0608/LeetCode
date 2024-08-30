type ProductOfNumbers struct {
    a []int
    n int
}


func Constructor() ProductOfNumbers {
    return ProductOfNumbers{[]int{1}, 0}
}


func (c *ProductOfNumbers) Add(num int)  {
    if num == 0 {
        c.n = 0
        c.a = []int{1}
    } else {
        c.a = append(c.a, num * c.a[c.n])
        c.n++
    }
    
}


func (c *ProductOfNumbers) GetProduct(k int) int {
    if c.n < k {
        return 0
    }
    return c.a[c.n] / c.a[c.n-k]
}


/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */