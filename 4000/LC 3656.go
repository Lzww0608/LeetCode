
import (
	"sort"
)
func simpleGraphExists(degrees []int) bool {
    n := len(degrees)
    sort.Ints(degrees)

    sum := make([]int, n + 1)
    pre := 0
    for i, x := range degrees {
        sum[i + 1] = sum[i] + x
    }

    if sum[n] & 1 == 1 {
        return false
    }

    search := func(k, r int) int {
        l := -1
        for l < r {
            mid := l + ((r - l + 1) >> 1)
            if degrees[mid] <= k {
                l = mid
            } else {
                r = mid - 1
            }
        }

        return l
    } 

    for k := 1; k < n; k++ {
        pre += degrees[n - k]
        i := search(k, n - k - 1)
        if pre > k * (k - 1) + sum[i + 1] + k * (n - k - i - 1) {
            return false
        }
    }

    return true
}