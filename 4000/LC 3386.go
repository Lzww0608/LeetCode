func buttonWithLongestTime(events [][]int) int {
    ans := events[0][0]
    mx := events[0][1]
    
    for i := 1; i < len(events); i++ {
        v := events[i]
        t := v[1] - events[i - 1][1] 
        if t > mx || t == mx && ans > v[0] {
            ans = v[0]
            mx = t
        }
    }

    return ans
}