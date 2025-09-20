func resultsArray(queries [][]int, k int) []int {
    ans := make([]int, len(queries))
    h := &hp{}
    for i, q := range queries {
        x, y := q[0], q[1]
        d := abs(x) + abs(y)
        heap.Push(h, d)

        if h.Len() > k {
            heap.Pop(h)
        }

        if i >= k - 1 {
            ans[i] = (*h)[0]
        } else {
            ans[i] = -1
        }
    }

    return ans
}

func abs(x int) int {
    if x < 0 {
        return -x
    }

    return x
}

type hp []int 
func (h hp) Len() int {return len(h)}
func (h hp) Less(i, j int) bool {return h[i] > h[j]}
func (h hp) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *hp) Push(x any) {
    *h = append(*h, x.(int))
}
func (h *hp) Pop() (x any) {
    old := *h 
    n := len(old)
    x = old[n - 1]
    *h = old[:n - 1]
    return 
}