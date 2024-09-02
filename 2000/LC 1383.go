const MOD int = 1_000_000_007
func maxPerformance(n int, speed []int, efficiency []int, k int) (ans int) {
    a := make([]tuple, n)
    for i := range a {
        a[i] = tuple{speed[i], efficiency[i]}
    }
    h := &hp{}
    heap.Init(h)
    sort.Slice(a, func(i, j int) bool {
        return a[i].eff > a[j].eff
    })

    sum := 0
    for _, v := range a {
        if h.Len() == k {
            sum -= heap.Pop(h).(int)
        }

        sum += v.sp
        heap.Push(h, v.sp)
        ans = max(ans, sum * v.eff)
    }

    return ans % MOD
}

type tuple struct {
    sp, eff int
}
type hp []int
func (h hp) Len() int {return len(h)}
func (h hp) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h hp) Less(i, j int) bool {return h[i] < h[j]}
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