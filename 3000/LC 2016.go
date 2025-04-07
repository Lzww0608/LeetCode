func maximumDifference(nums []int) (ans int) {
    mn := math.MaxInt32 
    for _, x := range nums {
        ans = max(ans, x - mn)
        mn = min(mn, x)
    }
    
    if ans == 0 {
        return -1
    }

    return 
}