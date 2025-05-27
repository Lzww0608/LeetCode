
import "math"
func maximumBeauty(flowers []int) int {
    n := len(flowers)
    pre := make([]int, n + 1)
    mn := make(map[int]int)
    mx := make(map[int]int)
    for i, x := range flowers {
        if _, ok := mn[x]; !ok {
            mn[x] = i
        } 
        mx[x] = i
        if x > 0 {
            pre[i + 1] = pre[i] + x
        } else {
            pre[i + 1] = pre[i]
        }
    }

    ans := math.MinInt32
    for k := range mn {
        if mn[k] == mx[k] {
            continue
        }
        cur := pre[mx[k] + 1] - pre[mn[k]]
        if k < 0 {
            cur += k * 2
        }
        ans = max(ans, cur)
    }

    return ans
}