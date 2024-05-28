func minRefuelStops(target int, startFuel int, stations [][]int) (ans int) {
    h := hp{}
    heap.Init(&h)
    f := startFuel
    s := 0
    stations = append(stations, []int{target, 0})
    for _, v := range stations {
        for h.Len() > 0 && v[0] - s > f {
            ans++
            f += heap.Pop(&h).(int)
        }
        if v[0] - s > f {
            return -1            
        }
        f -= v[0] - s
        s = v[0]
        heap.Push(&h, v[1])
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