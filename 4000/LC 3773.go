func maxSameLengthRuns(s string) (ans int) {
    m := make(map[int]int)

    n := len(s)
    for i := 0; i < n; i++ {
        j := i
        for j < n && s[j] == s[i] {
            j++
        }
        d := j - i
        m[d]++
        ans = max(ans, m[d])
        i = j - 1
    }
    
    return
}