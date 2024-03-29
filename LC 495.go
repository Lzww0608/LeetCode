func findPoisonedDuration(timeSeries []int, duration int) (ans int) {
    for i := 1; i < len(timeSeries); i++ {
        ans += min(duration, timeSeries[i] - timeSeries[i-1])
    }
    ans += duration
    return 
}