func longestBalanced(s string) (ans int) {
    a, b := 0, 0
    for i := range s {
        if s[i] == '1' {
            a++
        } else {
            b++
        }
    }

    pre := 0
    has0, has1 := false, false 
    m := make(map[int]int)
    m[0] = -1
    m0 := make(map[int]int)
    m1 := make(map[int]int)
    for i := range s {
        if s[i] == '1' {
            a--
            has1 = true 
            pre++
        } else {
            b--
            has0 = true 
            pre--
        }

        if j, ok := m[pre]; ok {
            ans = max(ans, i - j)
        }

        if b > 0 {
            if j, ok := m[pre - 2]; ok {
                ans = max(ans, i - j)
            }
        } else {
            if j, ok := m0[pre - 2]; ok {
                ans = max(ans, i - j)
            }
        }

        if a > 0 {
            if j, ok := m[pre + 2]; ok {
                ans = max(ans, i - j)
            }
        } else {
            if j, ok := m1[pre + 2]; ok {
                ans = max(ans, i - j)
            }
        }

        if has1 {
            if _, ok := m1[pre]; !ok {
                m1[pre] = i
            }
        } 
        if has0 {
            if _, ok := m0[pre]; !ok {
                m0[pre] = i
            }
        }
        if _, ok := m[pre]; !ok {
            m[pre] = i
        }
    }

    return
}