func spellchecker(wordlist []string, queries []string) []string {
    n := len(queries)
    ans := make([]string, n)
    exists := map[string]bool{}
    m := map[string]string{}
    mp := map[string]string{}

    for _, s := range wordlist {
        exists[s] = true
    }

    for _, s := range wordlist {
        builder := strings.Builder{}
        for i := range s {
            c := s[i]
            if c >= 'A' && c <= 'Z' {
                c = byte(c + 32)
            }
            builder.WriteByte(c)
        }
        tmp := builder.String()
        if _, ok := mp[tmp]; !ok {
            mp[tmp] = s
        }

        builder.Reset()
        for i := range tmp {
            if tmp[i] == 'e' || tmp[i] == 'i' || tmp[i] == 'o' || tmp[i] == 'u' {
                builder.WriteByte('a')
            } else {
                builder.WriteByte(tmp[i])
            }
        }
        tmp = builder.String()

        if _, ok := m[tmp]; !ok {
            m[tmp] = s
        }
    }

    for i, s := range queries {
        if exists[queries[i]] {
            ans[i] = queries[i]
        } else {
            builder := strings.Builder{}
            for i := range s {
                c := s[i]
                if c >= 'A' && c <= 'Z' {
                    c = byte(c + 32)
                }
                builder.WriteByte(c)
            }
            tmp := builder.String()
            if _, ok := mp[tmp]; ok {
                ans[i] = mp[tmp]
                continue
            }

            builder.Reset()
            for i := range tmp {
                if tmp[i] == 'e' || tmp[i] == 'i' || tmp[i] == 'o' || tmp[i] == 'u' {
                    builder.WriteByte('a')
                } else {
                    builder.WriteByte(tmp[i])
                }
            }
            tmp = builder.String()

            if _, ok := m[tmp]; ok {
                ans[i] = m[tmp]
            }
        }
    }
    return ans
}
