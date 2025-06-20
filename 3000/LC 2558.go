
import "container/heap"
func pickGifts(gifts []int, k int) (ans int64) {
    h := hp(gifts)
    heap.Init(&h)

    for k > 0 && gifts[0] > 1 {
        k--
        gifts[0] = int(math.Sqrt(float64(gifts[0])))
        heap.Fix(&h, 0)
    }

    for _, x := range h {
        ans += int64(x)
    }

    return
}


type hp []int
func (h hp) Len() int {return len(h)}
func (h hp) Less(i, j int) bool {return h[i] > h[j]}
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