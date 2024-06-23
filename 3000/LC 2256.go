func minimumAverageDifference(nums []int) int {
    idx, mn := -1, math.MaxInt32
    n := len(nums)
    pre, suf := 0, 0
    for _, x := range nums {
        suf += x 
    }
    
    for i, x := range nums[:n-1] {
        pre += x
        suf -= x
        t := abs(pre / (i + 1) - suf / (n - i - 1))
        if t < mn {
            mn, idx = t, i
        }
    }

    if (pre + nums[n-1]) / n < mn {
        return n - 1
    }

    return idx
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
