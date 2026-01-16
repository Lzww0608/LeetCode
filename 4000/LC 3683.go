func earliestTime(tasks [][]int) int {
    ans := math.MaxInt32 

    for _, v := range tasks {
        ans = min(ans, v[0] + v[1])
    }

    return ans
}