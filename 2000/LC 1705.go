func eatenApples(apples []int, days []int) (ans int) {
    h := hp{}
    heap.Init(&h)
    for i := 0; i < len(apples) || h.Len() > 0; i++ {
        if i < len(apples) {
            heap.Push(&h, pair{apples[i], i + days[i]})
        }
        for h.Len() > 0 && h[0].y <= i {
            heap.Pop(&h)
        }
        if h.Len() > 0 {
            ans++
            t := heap.Pop(&h).(pair)
            t.x--
            if t.x > 0 {
                heap.Push(&h, t)
            }
        }
    }

    return
}

type pair struct {
    x, y int
}
type hp []pair
func (h hp) Len() int {return len(h)}
func (h hp) Less(i, j int) bool {return h[i].y < h[j].y}
func (h hp) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *hp) Push(x any) {
    *h = append(*h, x.(pair))
}

func (h *hp) Pop() any {
    old := *h 
    n := len(old)
    x := old[n-1]
    old = old[:n-1]
    *h = old
    return x
}