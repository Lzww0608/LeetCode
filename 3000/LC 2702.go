
func minOperations(nums []int, x int, y int) int {
    mx := slices.Max(nums)
    l, r := 0, mx / y + 1 
    dif := x - y

    check := func(k int) bool {
        des := k * y 
        for _, x := range nums {
            x -= des
            if x > 0 {
                k -= (x + dif - 1) / dif
            }
        }

        return k >= 0
    }

    for l + 1 < r {
        mid := l + ((r - l) >> 1)
        if check(mid) {
            r = mid 
        } else {
            l = mid 
        }
    }

    return r 
}


type hp []int
func (h hp) Len() int {return len(h)}
func (h hp) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h hp) Less(i, j int) bool {return h[i] > h[j]}
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