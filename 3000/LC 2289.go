var inf int = int(1e9)

type pair struct {
    x, y int
}
func totalSteps(nums []int) int {
    n := len(nums)
    st := []pair{pair{0, inf}}
    res := 0

    for i := 0; i < n; i++ {
        cur := 0
        for len(st) > 0 && nums[st[len(st)-1].x] <= nums[i] {
            cur = max(cur, st[len(st)-1].y)
            st = st[:len(st)-1]
        }

        j := 0
        if len(st) == 0 {
            j = inf
        } else {
            j = cur + 1
            res = max(res, j)
        }
        
        st = append(st, pair{i, j})
    }

    return res
}
