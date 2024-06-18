func mincostToHireWorkers(quality []int, wage []int, k int) float64 {
    n := len(quality)
    var ans float64 = 0x7f7f7f7f
    worker := make([]float64, n)
    pos := make([]int, n)
    for i := range quality {
        pos[i] = i
        worker[i] = float64(wage[i]) / float64(quality[i])
    }
    sort.Slice(pos, func(i, j int) bool {
        return worker[pos[i]] < worker[pos[j]]
    })

    h := hp{}
    heap.Init(&h)
    sum := 0
    for _, i := range pos {
        q, w := quality[i], worker[i]
        sum += q
        heap.Push(&h, q)
        if h.Len() > k {
            t := heap.Pop(&h).(int)
            sum -= t
        }
        if h.Len() == k {
            ans = min(ans, float64(sum) * w)
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
func (h *hp) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}