func earliestFullBloom(plantTime []int, growTime []int) (ans int) {
    n := len(plantTime)
    pair := make([][]int, n)
    for i := range plantTime {
        pair[i] = []int{plantTime[i], growTime[i]}
    }

    h := hp(pair)
    heap.Init(&h)
    time := 0
    for h.Len() > 0 {
        v := heap.Pop(&h).([]int)
        time += v[0]
        ans = max(ans, time + v[1])
    }

    return ans
}


type hp [][]int
func (h hp) Len() int {return len(h)}
func (h hp) Less(i, j int) bool {return h[i][1] > h[j][1] || (h[i][1] == h[j][1] && h[i][0] > h[j][0])}
func (h hp) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *hp) Push(x any) {
    *h = append(*h, x.([]int))
} 
func (h *hp) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}