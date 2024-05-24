type SummaryRanges struct {
    mn, mx int
    m map[int]bool
}


func Constructor() SummaryRanges {
    return SummaryRanges{math.MaxInt32, math.MinInt32, map[int]bool{}}
}


func (this *SummaryRanges) AddNum(value int)  {
    this.m[value] = true
    this.mx = max(this.mx, value)
    this.mn = min(this.mn, value)
}


func (this *SummaryRanges) GetIntervals() (ans [][]int) {
    l, d := this.mn, 0
    for i := this.mn + 1; i <= this.mx + 1; i++ {
        if !this.m[i] {
            if l != -1 {
                ans = append(ans, []int{l, l + d})
            }
            l = -1
        } else {
            if l == -1 {
                l, d = i, 0
            } else {
                d++
            }
        }
    }
    return ans
}


/**
 * Your SummaryRanges object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(value);
 * param_2 := obj.GetIntervals();
 */