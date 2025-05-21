func maxRemoval(nums []int, queries [][]int) int {
    n, m := len(nums), len(queries)
    sort.Slice(queries, func(i, j int) bool {
        return queries[i][0] < queries[j][0]
    })

    h := &hp{}
    cur := 0
    dif := make([]int, n + 1)
    for i, j := 0, 0; i < n; i++ {
        cur += dif[i]
        for j < m && queries[j][0] <= i {
            heap.Push(h, queries[j][1])
            j++
        }

        for cur < nums[i] && h.Len() > 0 && (*h)[0] >= i {
            t := heap.Pop(h).(int)
            dif[t + 1]--
            cur++
        }

        if cur < nums[i] {
            return -1
        }
    }

    return h.Len()
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
    x = old[n-1]
    *h = old[:n-1]
    return
}