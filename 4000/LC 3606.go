func validateCoupons(code []string, businessLine []string, isActive []bool) []string {
    a := []int{}
    n := len(code)
    m := map[string]int {
        "electronics": 0,
        "grocery": 1,
        "pharmacy": 2,
        "restaurant": 3,
    }
    for i := range n {
        if !isActive[i] {
            continue
        }

        f := len(code[i]) != 0 
        for _, c := range code[i] {
            if !unicode.IsLetter(c) && !unicode.IsDigit(c) && c != '_' {
                f = false
                break
            }
        }

        if _, ok := m[businessLine[i]]; !f || !ok {
            continue
        }
        
        a = append(a, i)
    }

    sort.Slice(a, func(i, j int) bool {
        if m[businessLine[a[i]]] != m[businessLine[a[j]]] {
            return m[businessLine[a[i]]] < m[businessLine[a[j]]]
        }

        return code[a[i]] < code[a[j]]
    })

    ans := make([]string, len(a))
    for i, j := range a {
        ans[i] = code[j] 
    }

    return ans
}