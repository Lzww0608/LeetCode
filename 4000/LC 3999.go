const N int = 26
func minimumGroups(words []string) int {
    m := make(map[string]struct{}, len(words))    

    for _, word := range words {
        n := len(word)
        even := make([]byte, 0, (n + 1) / 2)
        odd := make([]byte, 0, n / 2)
        for k := range word {
            if k & 1 == 0 {
                even = append(even, word[k])
            } else {
                odd = append(odd, word[k])
            }
        }

        s := booth(string(even))
        t := booth(string(odd))
        cur := make([]byte, n)
        i, j := 0, 0
        for k := range cur {
            if k & 1 == 0 {
                cur[k] = s[i]
                i++
            } else {
                cur[k] = t[j]
                j++
            }
        }

        m[string(cur)] = struct{}{}
    }

    return len(m)
}

func booth(s string) string {
    n := len(s)
    i, j, k := 0, 1, 0
    for i < n && j < n && k < n {
        a := s[(i + k) % n]
        b := s[(j + k) % n]
        
        if a == b {
            k++
            continue
        }

        if a > b {
            i += k + 1
            if i <= j {
                i = j + 1
            }
        } else {
            j += k + 1
            if j <= i {
                j = i + 1
            }
        }
        k = 0
    }

    t := min(i, j)
    return s[t:] + s[:t]
}