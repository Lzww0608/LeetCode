const N int = 1e5 + 1
func maxEvents(events [][]int) (ans int) {
    m := make([][]int, N)
    mx := 0
    for i, e := range events {
        m[e[0]] = append(m[e[0]], i)
        mx = max(mx, e[1])
    } 

    h := hp{}
    heap.Init(&h)
    for i := 1; i <= mx; i++ {
        for _, j := range m[i] {
            heap.Push(&h, events[j][1])
        }
        for h.Len() > 0 && h[0] < i {
            heap.Pop(&h)
        }
        if h.Len() > 0 {
            heap.Pop(&h)
            ans++
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
func (h *hp) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}