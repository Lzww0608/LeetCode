func halveArray(nums []int) (ans int) {
    var sum float64
    h := FloatHeap{}
    for _, x := range nums {
        sum += float64(x)
        heap.Push(&h, float64(x))
    }
    sum /= 2.0
    for sum > 0 {
        t := heap.Pop(&h).(float64)
        sum -= t / 2.0
        ans++
        heap.Push(&h, t / 2.0)
    }
    return
}

type FloatHeap []float64
func (h FloatHeap) Len() int     { return len(h) }
func (h FloatHeap) Less(i, j int) bool { return h[i] > h[j] } 
func (h FloatHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *FloatHeap) Push(x any) {
    *h = append(*h, x.(float64))
}
func (h *FloatHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}
