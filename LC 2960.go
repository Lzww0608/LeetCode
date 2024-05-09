func countTestedDevices(batteryPercentages []int) int {
    ans, cnt := 0, 0
    //n := len(batteryPercentages)
    for _, x := range batteryPercentages {
        if x - cnt > 0 {
            ans++
            cnt++
        }
    }
    return ans
}