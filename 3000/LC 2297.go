
import "math"
func minCost(nums []int, costs []int) int64 {
    n := len(nums)
    f := make([]int, n)
    a := []int{0}
    b := []int{0}

    for i := 1; i < n; i++ {
        x := nums[i]
        cur := math.MaxInt64 / 2
        for len(a) > 0 && nums[a[len(a)-1]] <= x {
            cur = min(cur, f[a[len(a)-1]])
            a = a[:len(a)-1]
        }
        a = append(a, i)
        for len(b) > 0 && nums[b[len(b)-1]] > x {
            cur = min(cur, f[b[len(b)-1]])
            b = b[:len(b)-1]
        }
        b = append(b, i)
        f[i] = cur + costs[i]
    }

    return int64(f[n-1])
}