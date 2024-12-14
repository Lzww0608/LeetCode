
import "math"
func kEmptySlots(bulbs []int, k int) int {
    n := len(bulbs)
    if k >= n - 1 {
        return -1
    }
    days := make([]int, n)
    for i, x := range bulbs {
        days[x-1] = i + 1
    }

    ans := math.MaxInt32
    l, r := 0, k + 1
    for i := 0; r < n; i++ {
        if days[i] < days[l] || days[i] <= days[r] {
            if i == r {
                mx := max(days[l], days[r])
                ans = min(ans, mx)
            }
            l, r = i, i + k + 1
        }
    }

    if ans == math.MaxInt32 {
        return -1
    }
    return ans
}