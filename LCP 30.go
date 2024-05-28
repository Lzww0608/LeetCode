func magicTower(nums []int) int {
    h := hp{}
    heap.Init(&h)
    sum, ans := 1, 0
    for _, x := range nums {
        sum += x
    }
    if sum <= 0 {
        return -1
    } 
    sum = 1
    for _, x := range nums {
        if x < 0 {
            heap.Push(&h, x)
        }
        sum += x
        if sum <= 0 {
            sum -= heap.Pop(&h).(int)
            ans++
        }
    }

    return ans
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