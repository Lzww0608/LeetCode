func earliestFinishTime(landStartTime []int, landDuration []int, waterStartTime []int, waterDuration []int) int {
    a := math.MaxInt32
    for i := range landDuration {
        a = min(a, landDuration[i] + landStartTime[i])
    }

    ans := math.MaxInt32 
    b := math.MaxInt32 
    for i := range waterDuration {
        ans = min(ans, max(a, waterStartTime[i]) + waterDuration[i])
        b = min(b, waterDuration[i] + waterStartTime[i])
    }

    for i := range landDuration {
        ans = min(ans, max(b, landStartTime[i]) + landDuration[i])
    }

    return ans
}