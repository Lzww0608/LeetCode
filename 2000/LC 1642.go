func furthestBuilding(heights []int, bricks int, ladders int) int {
    n := len(heights)
    h := hp{}
    heap.Init(&h)

    for i := 1; i < n; i++ {
        x := heights[i] - heights[i-1]
        if x < 0 {
            continue
        }

        heap.Push(&h, x)
        if h.Len() > ladders {
            tmp := heap.Pop(&h).(int)
            if tmp > bricks {
                return i - 1
            } 
            bricks -= tmp
        }
        
    }
    return n - 1
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