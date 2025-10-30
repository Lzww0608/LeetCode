func earliestFinishTime(landStartTime []int, landDuration []int, waterStartTime []int, waterDuration []int) int {
    mn := math.MaxInt32

    for i := range landDuration {
        mn = min(mn, landDuration[i] + landStartTime[i])
    }

    mnn := math.MaxInt32
    ans := math.MaxInt32
    for i := range waterDuration {
        ans = min(ans, max(waterStartTime[i], mn) + waterDuration[i])

        mnn = min(mnn, waterDuration[i] + waterStartTime[i])
    }
    //fmt.Println(mn, mnn)
    for i := range landDuration { 
        ans = min(ans, max(landStartTime[i], mnn) + landDuration[i])
    }

    return ans
}