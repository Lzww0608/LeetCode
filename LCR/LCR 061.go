func kSmallestPairs(nums1 []int, nums2 []int, k int) (ans [][]int) {
    h := hp{}
    heap.Init(&h)
    for i := range nums1 {
        heap.Push(&h, []int{nums1[i] + nums2[0], i, 0})
    }
    

    for h.Len() > 0 && len(ans) < k {
        tmp := heap.Pop(&h).([]int)
        ans = append(ans, []int{nums1[tmp[1]], nums2[tmp[2]]})
        if tmp[2] + 1 >= len(nums2) {
            continue
        }
        heap.Push(&h, []int{nums1[tmp[1]] + nums2[tmp[2] + 1], tmp[1], tmp[2] + 1})
    }
    
    return 
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
