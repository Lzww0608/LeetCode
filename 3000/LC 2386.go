func kSum(nums []int, k int) (ans int64) {
    var sum int64 = 0
    for i, x := range nums {
        if x >= 0 {
            sum += int64(x)
        } else {
            nums[i] = -x
        }
    }

    sort.Ints(nums)
    h := hp{[]int{nums[0], 0}}
    ans = sum 
    heap.Init(&h)
    for i := 1; i < k; i++ {
        tmp := heap.Pop(&h).([]int)
        j := tmp[1]
        ans = sum - int64(tmp[0])
        if j + 1 == len(nums) {
            continue
        }
        heap.Push(&h, []int{tmp[0] + nums[j+1], j + 1})
        heap.Push(&h, []int{tmp[0] + nums[j+1] - nums[j], j + 1})
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