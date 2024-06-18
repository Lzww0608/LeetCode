func orderlyQueue(s string, k int) string {
    if k >= 2 {
        a := []byte(s)
        sort.Slice(a, func(i, j int) bool {
            return a[i] < a[j]
        })
        return string(a)
    }

    ans, mn := s, 'z'
    for i, c := range s {
        if c < mn {
            mn = c
            ans = s[i:] + s[:i] 
        } else if c == mn {
            tmp := s[i:] + s[:i]
            if tmp < ans {
                ans = tmp
            }
        }
    } 

    return ans
}