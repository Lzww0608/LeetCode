func maxTransactions(transactions []int) (ans int) {
    sum := 0
    h := &hp{}
    for _, x := range transactions {
        if sum + x >= 0 {
            if x < 0 {
                heap.Push(h, x)
            }
            sum += x
            ans++
        } else {
            if h.Len() > 0 && (*h)[0] < x {
                sum -= heap.Pop(h).(int)
                sum += x
                heap.Push(h, x)
            }
        }
    }

    return 
}

type hp []int 
func (h hp) Len() int {return len(h)}
func (h hp) Less(i, j int) bool {return h[i] < h[j]}
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