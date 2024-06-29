
type Solution struct {
    sum []int
    rects [][]int
}


func Constructor(rects [][]int) Solution {
    res := make([]int, len(rects) + 1)

    for i, v := range rects {
        a, b, x, y := v[0], v[1], v[2], v[3]
        res[i+1] = res[i] + (x - a + 1) * (y - b + 1)
    }

    return Solution{res, rects}
}


func (this *Solution) Pick() []int {
    n := rand.Intn(this.sum[len(this.sum)-1])
    pos := sort.SearchInts(this.sum, n + 1) - 1

    a, b,  y := this.rects[pos][0], this.rects[pos][1], this.rects[pos][3]

    return []int{(n - this.sum[pos]) / (y - b + 1) + a, (n - this.sum[pos])  % (y - b + 1) + b}

}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(rects);
 * param_1 := obj.Pick();
 */
