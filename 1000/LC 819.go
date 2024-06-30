func mostCommonWord(paragraph string, banned []string) string {
    m := map[string]bool{}
    for _, s := range banned {
        m[s] = true
    }
    
    mp := map[string]int{}
    builder := strings.Builder{}
    paragraph += " "
    for i := range paragraph {
        c := toLower(paragraph[i])
        if c >= 'a' && c <= 'z' {
            builder.WriteByte(c)
        } else if builder.Len() > 0 {
            s := builder.String()
            if !m[s] {
                mp[s]++
            }
            builder.Reset()
        }
    }

    cnt, ans := 0, ""
    for k, v := range mp {
        if v > cnt {
            cnt = v 
            ans = k
        }
    }

    return ans
}

func toLower(c byte) byte {
    if c >= 'A' && c <= 'Z' {
        return byte(c + 32)
    }
    return c
}
