func findMaxSum(a []int, b []int, k int) []int64 {
    n := len(a)
    ans := make([]int64, n)
    id := make([]int, n)
    for i := range id {
        id[i] = i
    }
    sort.Slice(id, func(i, j int) bool {
        return a[id[i]] < a[id[j]]
    })

    h := &hp{}
    cur := 0
    for j := 0; j < n; {
        cnt := 0
        for j + cnt < n && a[id[j]] == a[id[j+cnt]] {
            cnt++
        }
        for _, i := range id[j:j+cnt] {
            ans[i] = int64(cur)
        }
        for _, i := range id[j:j+cnt] {
            x := b[i]
            cur += x 
            heap.Push(h, x)
            if (*h).Len() > k {
                cur -= heap.Pop(h).(int)
            }
        }
        j += cnt
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
func (h *hp) Pop() (x any) {
    old := *h 
    n := len(old)
    x = old[n-1]
    *h = old[:n-1]
    return
}