func lastStoneWeightII(stones []int) int {
    f := make(map[int]bool)
    f[0] = true
    for _, x := range stones {
        cur := make(map[int]bool)
        for k := range f {
            cur[k + x] = true 
            cur[k - x] = true
        }
        f = cur
    }

    ans := math.MaxInt32 
    for k := range f {
        ans = min(ans, max(k, -k))
    }

    return ans
}