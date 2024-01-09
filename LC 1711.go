var mod int = 1e9 + 7
func countPairs(deliciousness []int) int {
    good := []int{1}
    for i := 1; i <= 21; i++ {
        good = append(good, (1 << i))
    }
    ans := 0
    m := map[int]int{}
    for _, x := range deliciousness {
        for _, k := range good {
            if v, ok := m[k-x]; ok {
                ans = (ans + v) % mod
            }
        }
        m[x]++
    }
    return ans
}