func minOperations(nums []int, k int) (ans int) {
    a := make([]int, 0, len(nums))
    for _, x := range nums {
        if x < k {
            a = append(a, x)
        }
    }
    h := hp(a)
    heap.Init(&h)

    for h.Len() > 1 {
        ans++
        a := heap.Pop(&h).(int)
        b := heap.Pop(&h).(int)
        x := a + b + min(a, b)
        if x < k {
            heap.Push(&h, x)
        }
    }

    return ans + h.Len() & 1
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