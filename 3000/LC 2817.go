
import (
	"math"
)
func minAbsoluteDifference(a []int, x int) int {
    ans := math.MaxInt

    t := redblacktree.NewWithIntComparator()
    t.Put(math.MaxInt, nil)
    t.Put(math.MinInt / 2, nil)
    for i, y := range a[x:] {
        t.Put(a[i], nil)
        c, _ := t.Ceiling(y)
        f, _ := t.Floor(y)
        ans = min(ans, c.Key.(int) - y, y - f.Key.(int))
    }

    return ans
}