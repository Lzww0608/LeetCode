type ExamRoom struct {
    n int
    a []int
}


func Constructor(n int) ExamRoom {
    return ExamRoom{n: n}
}


func (this *ExamRoom) Seat() int {
    if len(this.a) == 0 {
        this.a = append(this.a, 0)
        return 0
    }
    maxDist := this.a[0]
    seatToSit := 0
    for i := 0; i + 1 < len(this.a); i++ {
        d := (this.a[i+1] - this.a[i]) / 2
        if d > maxDist {
            maxDist = d
            seatToSit = this.a[i] + d
        }
    }
    if this.n - 1 - this.a[len(this.a)-1] > maxDist {
        seatToSit = this.n - 1
    }
    this.a = append(this.a, seatToSit)
    sort.Ints(this.a)
    return seatToSit
}


func (this *ExamRoom) Leave(p int)  {
    for i, x := range this.a {
        if x == p {
            this.a = append(this.a[:i], this.a[i+1:]...)
            break
        }
    }
    return
}


/**
 * Your ExamRoom object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Seat();
 * obj.Leave(p);
 */