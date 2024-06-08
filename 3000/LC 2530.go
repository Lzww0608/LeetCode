func maxKelements(nums []int, k int) (ans int64) {
    h := hp(nums)
    heap.Init(&h)
    for i := 0; i < k; i++ {
        t := heap.Pop(&h).(int)
        ans += int64(t)
        heap.Push(&h, (t + 2) / 3)
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
func (h *hp) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}