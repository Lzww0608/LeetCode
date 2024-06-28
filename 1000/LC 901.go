type pair struct {
    x, y int
}

type StockSpanner struct {
    a []pair
    n int
}


func Constructor() StockSpanner {
    return StockSpanner{n: -1}
}


func (this *StockSpanner) Next(price int) int {
    this.n++
    for len(this.a) > 0 && price >= this.a[len(this.a)-1].x {
        this.a = this.a[:len(this.a)-1]
    }
    this.a = append(this.a, pair{price, this.n})
    if len(this.a) == 1 {
        return this.n + 1
    } 
    return this.n - this.a[len(this.a)-2].y
}


/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
