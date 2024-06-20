const MOD int = int(1e9 + 7)
func countNicePairs(nums []int) (ans int) {
    m := map[int]int{}
    
    rev := func(x int) int {
        res := 0
        for x > 0 {
            res = res * 10 + x % 10
            x /= 10
        }
        return res
    }
    
    for _, x := range nums {
        t := x - rev(x)
        ans = (ans + m[t]) % MOD
        m[t]++
    }

    return 
}
