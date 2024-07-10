func findLongestSubarray(s []string) []string {
    a := make([]int, len(s))
    for i, c := range s {
        if unicode.IsDigit(rune(c[0])) {
            a[i] = 1
        } 
    } 
    
    cnt := 0
    l, r := 0, 0
    m := map[int]int{0: -1}
    for i, x := range a {
        if x == 1 {
            cnt++
        } else {
            cnt--
        }

        if j, ok := m[cnt]; !ok {
            m[cnt] = i
        } else if i - j > r - l {
            l, r = j, i
        }
    }

    return s[l+1:r+1]
}