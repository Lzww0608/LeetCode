
import "math"
func minimumSum(a []int, b []int) int {
    pos := make(map[int]int)
    for i, x := range a {
        if _, ok := pos[x]; !ok {
            pos[x] = i
        }
    }

    ans := math.MaxInt32
    for i, x := range b {
        if j, ok := pos[x]; ok {
            ans = min(ans, i + j)
        } 
    }

    if ans == math.MaxInt32 {
        return -1
    }
    return ans
}