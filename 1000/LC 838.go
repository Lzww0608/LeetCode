func pushDominoes(dominoes string) string {
    s := []rune(dominoes)
    n := len(s)
    if n == 1 {
        return string(s)
    }
    
    l := 'L'
    for i := 0; i < n; i++ {
        j := i
        for j < n && s[j] == '.' {
            j++
        }

        r := 'R'
        if j < n {
            r = s[j]
        }

        if l == r {
            for i < j {
                s[i] = l
                i++
            }
        } else if l == 'R' {
            for t := j - 1; t > i; t-- {
                s[i] = 'R'
                s[t] = 'L'
                i++
            }
        }
        i, l = j, r
    }

    return string(s)
}