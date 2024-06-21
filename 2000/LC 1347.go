func minSteps(s string, t string) (ans int) {
    m := [26]int{}
    for i := range s {
        x := int(s[i] - 'a')
        y := int(t[i] - 'a')
        m[x]++
        m[y]--
    }
    for _, x := range m {
        if x > 0 {
            ans += x
        }
    }

    return 
}