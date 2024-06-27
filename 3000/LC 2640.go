func totalCost(costs []int, k int, candidates int) (ans int64) {
    h1 := IntHeap{}
    h2 := IntHeap{}
    n := len(costs)
    i, j := 0, n - 1
    for i = 0; i < n && h1.Len() < candidates; i++ {
        heap.Push(&h1, costs[i])
    }
    for j = n - 1; j >= i && h2.Len() < candidates; j-- {
        heap.Push(&h2, costs[j])
    }

    for t := 0; t < k; t++ {
        if h2.Len() == 0 || h1.Len() > 0 && h1[0] <= h2[0] {
            ans += int64(heap.Pop(&h1).(int))
            if i <= j {
                heap.Push(&h1, costs[i])
                i++
            }
        } else {
            ans += int64(heap.Pop(&h2).(int))
            if i <= j {
                heap.Push(&h2, costs[j])
                j--
            }
        }
    }

    return
}

type IntHeap []int
func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } 
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any) {
    *h = append(*h, x.(int))
}
func (h *IntHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}
