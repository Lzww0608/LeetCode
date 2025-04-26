func putMarbles(weights []int, k int) int64 {
    h := &hp{}
    hh := &hp{}
    mx, mn := 0, 0
    n := len(weights)

    for i := 0; i < n - 1; i++ {
        x := weights[i] + weights[i+1]
        heap.Push(h, x)
        heap.Push(hh, -x)
        mx += x 
        mn += x 
        if (*h).Len() > k - 1 {
            mx -= heap.Pop(h).(int)
            mn -= -heap.Pop(hh).(int)
        }
    }

    return int64(mx - mn)
}


type hp []int
func (h hp) Len() int {return len(h)}
func (h hp) Less(i, j int) bool {return h[i] < h[j]}
func (h hp) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *hp) Push(x any) {
    *h = append(*h, x.(int))
}
func (h *hp) Pop() (x any) {
    old := *h 
    n := len(old)
    x = old[n-1]
    *h = old[:n-1]
    return
}