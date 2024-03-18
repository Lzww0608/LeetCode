type Solution struct {
    cnt, m, n int
    mp map[int]int
}


func Constructor(m int, n int) Solution {
    return Solution{m * n, m, n, map[int]int{}}
}


func (this *Solution) Flip() []int {
    x := rand.Intn(this.cnt)
    this.cnt--
    idx := x
    if v, ok := this.mp[x]; ok {
        idx = v
    }
    if _, ok := this.mp[this.cnt]; ok {
        this.mp[x] = this.mp[this.cnt]
    } else {
        this.mp[x] = this.cnt
    }

    return []int{idx / this.n, idx % this.n}
}


func (this *Solution) Reset()  {
    this.cnt = this.m * this.n
    this.mp = map[int]int{}
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(m, n);
 * param_1 := obj.Flip();
 * obj.Reset();
 */