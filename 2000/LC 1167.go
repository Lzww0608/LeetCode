
import "container/heap"
func connectSticks(sticks []int) (ans int) {
    h := hp(sticks)
    heap.Init(&h)
    for h.Len() > 1 {
        a := heap.Pop(&h).(int)
        b := heap.Pop(&h).(int)
        ans += a + b
        heap.Push(&h, a + b)
    }

    return
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

