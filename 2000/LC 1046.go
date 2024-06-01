type IntHeap []int
func lastStoneWeight(stones []int) int {
    h := &IntHeap{}
    //heap.Init(h);
    for _, x := range stones {
        heap.Push(h, x)
    }

    for h.Len() > 1 {
        x := heap.Pop(h).(int)
        y := heap.Pop(h).(int)
        heap.Push(h, abs(x - y))
    }

    if h.Len() == 0 {
        return 0
    }
    return heap.Pop(h).(int)
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] } // Min-Heap
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
