type Checkout struct {
    q []int
    d []int
}


func Constructor() Checkout {
    return Checkout{[]int{}, []int{}}
}


func (this *Checkout) Get_max() int {
    if len(this.d) == 0 {
        return -1
    }
    return this.d[0]
}


func (this *Checkout) Add(value int)  {
    for len(this.d) > 0 && this.d[len(this.d)-1] < value {
        this.d = this.d[:len(this.d)-1]
    }
    this.d = append(this.d, value)
    this.q = append(this.q, value)
}


func (this *Checkout) Remove() int {
    if len(this.q) == 0 {
        return -1
    }
    res := this.q[0]
    if res == this.d[0] {
        this.d = this.d[1:]
    }
    this.q = this.q[1:]
    return res
}


/**
 * Your Checkout object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get_max();
 * obj.Add(value);
 * param_3 := obj.Remove();
 */