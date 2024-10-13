
import "math"
import "container/heap"
func minimumDeviation(nums []int) int {
    h := &hp{}
    heap.Init(h)
    mn := math.MaxInt32
    for i, x := range nums {
        if x & 1 == 1 {
            nums[i] <<= 1
        }
        mn = min(mn, nums[i])
        heap.Push(h, nums[i])
    }
    ans := (*h)[0] - mn
    for h.Len() > 0 && (*h)[0] & 1 == 0 && ans > 0 {
        t := heap.Pop(h).(int)
        t >>= 1
        heap.Push(h, t)
        mn = min(t, mn)
        if h.Len() > 0 {
            ans = min(ans, (*h)[0] - mn)
        }
    }

    return ans
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