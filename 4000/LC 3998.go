func transformStr(s string, strs []string) []bool {
    m := len(strs)
    ans := make([]bool, m)
    zero := strings.Count(s, "0")

next:
    for k, t := range strs {
        a := strings.Count(t, "0") 
        b := strings.Count(t, "?") 
        if a > zero || a + b < zero {
            continue
        }
        x, y := 0, 0
        for j := range t {
            if t[j] == '?' {
                if a < zero {
                    x++
                    a++
                }
            } else if t[j] == '0' {
                x++
            }

            if s[j] == '0' {
                y++
            }
            if x < y {
                continue next
            }
        }

        ans[k] = true
    }

    return ans
}