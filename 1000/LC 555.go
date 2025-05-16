func splitLoopedString(strs []string) (ans string) {
    for i, s := range strs {
        t := make([]byte, len(s))
        for i, j := 0, len(s) - 1; j >= 0; i, j = i + 1, j - 1 {
            t[i] = s[j]
        }
        if string(t) > s {
            strs[i] = string(t)
        }
    }

    n := len(strs)
    for i := 0; i < n; i++ {
        var cur string 
        for j := i + 1; j < n; j++ {
            cur += strs[j]
        }
        for j := 0; j < i; j++ {
            cur += strs[j]
        }

        for j := 0; j < len(strs[i]); j++ {
            a, b := []byte{}, []byte{}
            for k := len(strs[i]) - 1; k >= j; k-- {
                a = append(a, strs[i][k])
            }
            for k := j - 1; k >= 0; k-- {
                b = append(b, strs[i][k])
            } 
            tmp := string(b) + cur + string(a)
            ans = max(ans, tmp)
        }

        for j := 0; j < len(strs[i]); j++ {
            a, b := []byte{}, []byte{}
            for k := j; k < len(strs[i]); k++ {
                a = append(a, strs[i][k])
            }
            for k := 0; k < j; k++ {
                b = append(b, strs[i][k])
            } 
            tmp := string(a) + cur + string(b)
            ans = max(ans, tmp)
        }
    }

    return
}