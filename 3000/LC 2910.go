
import "math"
func minGroupsForValidAssignment(balls []int) int {
    cnt := make(map[int]int)
    for _, ball := range balls {
        cnt[ball]++
    }
    mn := math.MaxInt32 
    for _, x := range cnt {
        mn = min(mn, x)
    }

    a := make([]int, mn + 1)
    a[0] = math.MaxInt32
    for _, x := range cnt {
        for i := 1; i <= mn; i++ {
            if a[i] >= math.MaxInt32 {
                continue
            }
            p, q := x / i, x % i 
            if q > p {
                a[i] += math.MaxInt32
            } else {
                a[i] += (x + i) / (i + 1)
            }
        }
    }

    return slices.Min(a)
}