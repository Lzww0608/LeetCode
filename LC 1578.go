func minCost(colors string, neededTime []int) (ans int) {
    n := len(colors)
    for i := 0; i < n - 1; i++ {
        if colors[i] == colors[i+1] {
            j := i + 1
            mx, sum := neededTime[i], neededTime[i]
            for ; j < n && colors[j] == colors[i]; j++ {
                mx = max(mx, neededTime[j])
                sum += neededTime[j]
            }
            ans += sum - mx
            i = j - 1
        }
    }

    return 
}