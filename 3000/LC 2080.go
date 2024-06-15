type RangeFreqQuery struct {
    m map[int][]int
}


func Constructor(arr []int) RangeFreqQuery {
    m := map[int][]int{}
    for i, x := range arr {
        m[x] = append(m[x], i)
    }

    return RangeFreqQuery{m}
}


func (this *RangeFreqQuery) Query(left int, right int, value int) int {
    a := this.m[value]
    ans := sort.SearchInts(a, right + 1)
    ans -= sort.SearchInts(a, left)

    return ans
}


/**
 * Your RangeFreqQuery object will be instantiated and called as such:
 * obj := Constructor(arr);
 * param_1 := obj.Query(left,right,value);
 */