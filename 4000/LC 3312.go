
import (
	"slices"
	"sort"
)
func gcdValues(nums []int, queries []int64) []int {
    mx := slices.Max(nums)
    cnt_x := make([]int, mx + 1)
    cnt_gcd := make([]int, mx + 1)

    for _, x := range nums {
        cnt_x[x]++
    }

    for i := mx; i > 0; i-- {
        cnt := 0
        for j := i; j <= mx; j += i {
            cnt += cnt_x[j]
            cnt_gcd[i] -= cnt_gcd[j]
        }
        cnt_gcd[i] += cnt * (cnt - 1) / 2
    }   

    for i := 1; i <= mx; i++ {
        cnt_gcd[i] += cnt_gcd[i - 1]
    }

    ans := make([]int, len(queries))
    for i, query := range queries {
        ans[i] = sort.SearchInts(cnt_gcd, int(query) + 1)
    }

    return ans
}