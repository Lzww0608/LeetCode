func maxHammingDistances(nums []int, m int) []int {
    n := len(nums)
    ans := make([]int, n)
    dis := make([]int, 1 << m)
    for i := range dis {
        dis[i] = math.MaxInt32
    }
    h := &hp{}
    heap.Init(h)
    for _, x := range nums {
        dis[x] = 0
        heap.Push(h, pair{x, 0})
    }
    

    for h.Len() > 0 {
        cur := heap.Pop(h).(pair)
        for i := 0; i < m; i++ {
            dest := cur.to ^ (1 << i)
            if cur.weight + 1 < dis[dest] {
                dis[dest] = cur.weight + 1
                heap.Push(h, pair{dest, dis[dest]})
            }
        }
    }

    mask := 1 << m - 1
    for i := 0; i < n; i++ {
        ans[i] = m - dis[mask ^ nums[i]]
    }
    return ans
}


type pair struct {
    to, weight int
}
type hp []pair
func (h hp) Len() int {return len(h)}
func (h hp) Less(i, j int) bool {return h[i].weight < h[j].weight}
func (h hp) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *hp) Push(x any) {
    *h = append(*h, x.(pair))
} 
func (h *hp) Pop() (x any) {
    old := *h 
    n := len(old)
    x = old[n-1]
    *h = old[:n-1]
    return 
}