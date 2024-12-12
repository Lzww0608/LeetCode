
import "container/heap"
func maxSpending(values [][]int) int64 {
    h := &hp{}
    n := len(values[0])
    for i := range values {
        heap.Push(h, []int{values[i][n-1], i})
        values[i] = values[i][:n-1]
    }    

    ans := 0
    for d := 1; h.Len() > 0; d++ {
        cur := heap.Pop(h).([]int)
        x, i := cur[0], cur[1]
        ans += x * d 
        if len(values[i]) > 0 {
            m := len(values[i])
            heap.Push(h, []int{values[i][m-1], i})
            values[i] = values[i][:m-1]
        }
    }

    return int64(ans)
}


type hp [][]int
func (h hp) Len() int {return len(h)}
func (h hp) Less(i, j int) bool {return h[i][0] < h[j][0]}
func (h hp) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *hp) Push(x any) {
    *h = append(*h, x.([]int))
}
func (h *hp) Pop() (x any) {
    old := *h
    n := len(old)
    x = old[n-1]
    *h = old[:n-1]
    return
}