func maxWidthRamp(a []int) int {
    ans := 0
    s := []int{}
    for i, x := range a {
        if len(s) == 0 || a[s[len(s)-1]] > x {
            s = append(s, i)
        } 
    }

    for i := len(a) - 1; i >= 0; i-- {
        for len(s) > 0 && a[s[len(s)-1]] <= a[i] {
            ans = max(ans, i - s[len(s)-1])
            s = s[:len(s)-1]
        }
    }
    return ans
}