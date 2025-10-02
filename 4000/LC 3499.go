func maxActiveSectionsAfterTrade(s string) (ans int) {
    l, r := -1, -1
    a := [][2]int{}
    for i := range s {
        if s[i] == '0' {
            if l == -1 {
                l, r = i, i
            } else {
                r = i
            }
        } else {
            if l != -1 {
                a = append(a, [2]int{l, r})
                l = -1
            }
            ans++
        }
    }

    if l != -1 {
        a = append(a, [2]int{l, r})
    }

    mx := 0
    for i := 0; i + 1 < len(a); i++ {
        d := a[i][1] - a[i][0] + 1 + a[i + 1][1] - a[i + 1][0] + 1
        mx = max(mx, d)
    }

    return ans + mx
}