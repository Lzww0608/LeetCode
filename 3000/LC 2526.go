type DataStream struct {
    val, cnt, k int 
}


func Constructor(value int, k int) DataStream {
    return DataStream{value, 0, k}
}


func (c *DataStream) Consec(num int) bool {
    if num == c.val {
        c.cnt++
    } else {
        c.cnt = 0
    }

    return c.cnt >= c.k
}


/**
 * Your DataStream object will be instantiated and called as such:
 * obj := Constructor(value, k);
 * param_1 := obj.Consec(num);
 */