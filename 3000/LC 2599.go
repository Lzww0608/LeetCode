
import "container/heap"
func makePrefSumNonNegative(nums []int) (ans int) {
    h := &hp{}
    sum := 0
    for _, x := range nums {
        sum += x
        heap.Push(h, x)
        if sum < 0 {
            ans++
            sum -= heap.Pop(h).(int)
        }
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