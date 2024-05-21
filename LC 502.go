func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
    n := len(profits)
    IPO := make([][2]int, n)
    for i := range IPO {
        IPO[i][0] = capital[i]
        IPO[i][1] = profits[i] 
    }
    sort.Slice(IPO, func(i, j int) bool {
        return IPO[i][0] < IPO[j][0]
    })
    h := IntHeap{}
    heap.Init(&h)
    j := 0
    for k > 0 {
        for j < n && IPO[j][0] <= w {
            heap.Push(&h, IPO[j][1])
            j++
        }
        if h.Len() == 0 {
            break
        }
        w += heap.Pop(&h).(int)
        k--
    }

    return w
}

type IntHeap []int
func (h IntHeap) Len() int {return len(h)}
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j]}
func (h IntHeap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *IntHeap) Push(x any) {
    *h = append(*h, x.(int))
}
func (h *IntHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}