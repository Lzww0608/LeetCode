func maximumSumOfHeights(maxHeights []int) int64 {
    var ans int64 = 0
    for i, peak := range maxHeights {
        h := int64(peak)
        cur := h
        for j := i-1; j >= 0; j-- {
            h += int64(min(cur, int64(maxHeights[j])))
            cur = int64(min(cur, int64(maxHeights[j])))
        } 
        cur = int64(peak)
        for j := i+1; j < len(maxHeights); j++ {
            h += int64(min(cur, int64(maxHeights[j])))
            cur = int64(min(cur, int64(maxHeights[j])))
        }
        ans = max(ans, h)
    }
    return ans
}