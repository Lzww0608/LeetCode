func lexGreaterPermutation(s string, target string) string {
    cnt := [26]int{}
    for i := 0; i < len(s); i++ {
        cnt[s[i]-'a']++
    }

    n := len(target)
    
    matchLen := 0
    for matchLen < n {
        idx := int(target[matchLen] - 'a')
        if cnt[idx] > 0 {
            cnt[idx]--
            matchLen++
        } else {
            break
        }
    }

    for i := matchLen; i >= 0; i-- {
        if i == n {
            continue
        }

        tVal := int(target[i] - 'a')

        if i < matchLen {
            cnt[tVal]++
        }

        for charCode := tVal + 1; charCode < 26; charCode++ {
            if cnt[charCode] > 0 {
                cnt[charCode]--

                res := make([]byte, 0, n)
                
                res = append(res, target[:i]...)
                
                res = append(res, byte('a'+charCode))

                for k := 0; k < 26; k++ {
                    for cnt[k] > 0 {
                        res = append(res, byte('a'+k))
                        cnt[k]--
                    }
                }
                return string(res)
            }
        }
    }

    return ""
}