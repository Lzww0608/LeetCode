func minSum(a []int, b []int) int64 {
    sum1, sum2 := 0, 0
    cnt1, cnt2 := 0, 0
    for _, x := range a {
        sum1 += x 
        if x == 0 {
            cnt1++
        }
    }
    for _, x := range b {
        sum2 += x 
        if x == 0 {
            cnt2++
        }
    }

    sum1 += cnt1 
    sum2 += cnt2 
    if sum1 < sum2 && cnt1 == 0 || sum2 < sum1 && cnt2 == 0 {
        return -1
    }

    return int64(max(sum1, sum2))
}