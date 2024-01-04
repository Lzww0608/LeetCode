func missingRolls(rolls []int, mean int, n int) []int {
    m := len(rolls)
    sum := mean * (m + n)
    for _, x := range rolls {
        sum -= x
    }
    if sum < n || sum > n * 6 {
        return []int{}
    }
    ans := []int{}
    for sum > 0 {
        t := sum / n 
        n--
        sum -= t
        ans = append(ans, t)
    }
    return ans
}