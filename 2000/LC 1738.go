func kthLargestValue(matrix [][]int, k int) int {
    n := len(matrix[0])
    pre := make([]int, n + 1)
    g := make([]int, n + 1)
    h := hp{}
    heap.Init(&h)
    for i := range matrix {
        t := 0
        copy(g, pre)
        for j, x := range matrix[i] {
            t ^= g[j + 1] ^ g[j] ^ x
            pre[j+1] = t 
            heap.Push(&h, t)
            if h.Len() > k {
                heap.Pop(&h)
            }
        }
    }

    return h[0]
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
