func maxPoints(a []int, b []int, k int) int64 {
    n := len(a)
    sum := 0

    h := &hp{}
    for i := range n {
        sum += b[i]
        heap.Push(h, pair{a[i] - b[i], i})
    }

    for range k {
        sum += heap.Pop(h).(pair).x
    }

    for h.Len() > 0 && (*h)[0].x > 0 {
        sum += heap.Pop(h).(pair).x
    }

    return int64(sum)
}

type pair struct {
    x, i int 
}
type hp []pair 
func (h hp) Len() int {return len(h)}
func (h hp) Less(i, j int) bool {return h[i].x > h[j].x}
func (h hp) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *hp) Push(x any) {
    *h = append(*h, x.(pair))
}
func (h *hp) Pop() (x any) {
    old := *h 
    n := len(old)
    x = old[n - 1]
    *h = old[:n - 1]
    return
}