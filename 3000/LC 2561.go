
import "math"
func minCost(a []int, b []int) (ans int64) {
    cnt := make(map[int]int)
    for i := 0; i < len(a); i++ {
        cnt[a[i]]++
        cnt[b[i]]--
    }

    s := []int{}
    mn := math.MaxInt32
    for i, x := range cnt {
        x = max(x, -x)
        if x % 2 != 0 {
            return -1
        }
        mn = min(mn, i)
        for t := x / 2; t > 0; t-- {
            s = append(s, i)
        }
    }

    sort.Ints(s)
    for i := 0; i < len(s) / 2; i++ {
        ans += int64(min(s[i], s[len(s)-i-1], mn * 2))
    }

    return 
}