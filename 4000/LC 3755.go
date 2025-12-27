func maxBalancedSubarray(nums []int) (ans int) {
    type pair struct {
        x, y int 
    }

    cur, d := 0, 0
    m := make(map[pair]int)
    m[pair{0, 0}] = -1
    for i, x := range nums {
        cur ^= x 
        if x & 1 == 0 {
            d++
        } else {
            d--
        }
        if v, ok := m[pair{cur, d}]; ok {
            ans = max(ans, i - v)
        } else {
            m[pair{cur, d}] = i
        }
    }

    return
}