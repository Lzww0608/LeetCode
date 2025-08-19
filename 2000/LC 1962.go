
import "container/heap"
func minStoneSum(piles []int, k int) (ans int) {
    h := hp(piles)
    heap.Init(&h)

    for i := 0; i < k; i++ {
        h[0] = (h[0] + 1) / 2
        heap.Fix(&h, 0)
    }

    for _, x := range h {
        ans += x
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
    x = old[n - 1]
    *h = old[:n - 1]
    return
}