func lengthLongestPath(input string) (ans int) {
    n := len(input)
    st := []int{}
    
    for i := 0; i < n; i++ {
        dep := 1
        for i < n && input[i] == '\t' {
            i++
            dep++
        }
        d := 0
        isFile := false
        for i < n && input[i] != '\n' {
            d++
            if input[i] == '.' {
                isFile = true
            }
            i++
        }

        for len(st) >= dep {
            st = st[:len(st)-1]
        }
        if len(st) > 0 {
            d += st[len(st)-1] + 1
        }

        if isFile {
            ans = max(ans, d)
        } else {
            st = append(st, d)
        }
    }

    return
}