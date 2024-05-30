func smallestRange(nums [][]int) []int {
    a := [][]int{}
    n := len(nums)
    h := hp{}
    heap.Init(&h)
    for i := range nums {
        for _, x := range nums[i] {
            heap.Push(&h, []int{x, i})
        }
    }
    for h.Len() > 0 {
        a = append(a, heap.Pop(&h).([]int))
    }

    cnt := map[int]int{}
    l, r, mn := 0, 0, a[len(a)-1][0] - a[0][0]
    ans := []int{a[0][0], a[len(a)-1][0]}
    for r < len(a) {
        cnt[a[r][1]]++
        for l < r && len(cnt) == n {
            if mn > a[r][0] - a[l][0] {
                mn = a[r][0] - a[l][0]
                ans = []int{a[l][0], a[r][0]}
            }
            cnt[a[l][1]]--
            if cnt[a[l][1]] == 0 {
                delete(cnt, a[l][1])
            }
            l++
        }
        r++
    }

    return ans
}

type hp [][]int
func (h hp) Len() int {return len(h)}
func (h hp) Less(i, j int) bool {return h[i][0] < h[j][0]}
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