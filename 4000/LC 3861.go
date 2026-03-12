func minimumIndex(capacity []int, itemSize int) int {
    ans := -1
    mn := math.MaxInt32 
    for i, x := range capacity {
        if x >= itemSize && x < mn {
            mn = x 
            ans = i
        }
    }

    return ans
}