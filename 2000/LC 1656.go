type OrderedStream struct {
    m map[int]string
    p int
}


func Constructor(n int) OrderedStream {
    m := make(map[int]string)
    return OrderedStream{m, 1}
}


func (c *OrderedStream) Insert(idKey int, value string) (ans []string) {
    c.m[idKey] = value
    for true {
        if  v, ok := c.m[c.p]; ok {
            ans = append(ans, v)
            c.p++
        } else {
            break
        }
    }
    

    return 
}


/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(idKey,value);
 */