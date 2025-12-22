func storeWater(bucket []int, vat []int) int {
    n := len(bucket)
    ans := math.MaxInt32 
    mx := slices.Max(vat)
    if mx == 0 {
        return 0
    }

    for i := 1; i <= mx; i++ {
        cur := 0
        for j := range n {
            cur += max(0, (vat[j] + i - 1) / i - bucket[j])
        }
        ans = min(ans, cur + i)
    }

    return ans
}