
import "container/heap"
func convertArray(nums []int) (ans int) {
    ans = solve(nums)
    for i := range nums {
        nums[i] = - nums[i]
    }

    return min(ans, solve(nums))
}

func solve(nums []int) (ans int) {
    h := &hp{}

    for _, x := range nums {
        if h.Len() > 0 && x < (*h)[0] {
            y := heap.Pop(h).(int)
            ans += y - x 
            heap.Push(h, x)
        }
        heap.Push(h, x)
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