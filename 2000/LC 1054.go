type pair struct {
    x, y int
}
type IntHeap []pair
func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i].x > h[j].x } //maxheap
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any) {
    *h = append(*h, x.(pair))
}
func (h *IntHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}


func rearrangeBarcodes(barcodes []int) []int {
    n := len(barcodes)
    ans := make([]int, n)
    m := map[int]int{}
    for _, x := range barcodes {
        m[x]++
    }
    
    h := &IntHeap{}
    heap.Init(h)
    for k, v := range m {
        heap.Push(h, pair{v, k})
    }

    k := 0
    var pre pair
    for h.Len() > 0 {
        cur := heap.Pop(h).(pair)
        ans[k] = cur.y
        cur.x--
        k++
        if pre.x > 0 {
            heap.Push(h, pre)
        }
        pre = cur
    }

    return ans
}
