func mostCompetitive(a []int, k int) []int {
    k = len(a) - k
    s := []int{}
    for _, x := range a {
        for k > 0 && len(s) > 0 && s[len(s)-1] > x {
            s = s[:len(s)-1]
            k--
        }
        s = append(s, x)
    }
    for k > 0 {
        k--
        s = s[:len(s)-1]
    } 
    return s
}