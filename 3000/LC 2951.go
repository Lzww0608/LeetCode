func findPeaks(mountain []int) (ans []int) {
    n := len(mountain)
    for i := 1; i < n - 1; i++ {
        if mountain[i] > mountain[i-1] && mountain[i] > mountain[i+1] {
            ans = append(ans, i)
        }
    }

    return
}
