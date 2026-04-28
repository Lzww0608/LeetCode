func sortVowels(s string) string {
    n := len(s)
    ans := make([]byte, n)
    cnt := make([]int, 5)
    first := make([]int, 5)
    alp := []byte{'a', 'e', 'i', 'o', 'u'}

    for i := range s {
        idx := getIndex(s[i])
        if idx >= 0 {
            cnt[idx]++
            if first[idx] == 0 {
                first[idx] = i + 1
            }
        }
    }

    p := make([]int, 5)
    for i := range p {
        p[i] = i
    }
    sort.Slice(p, func(i, j int) bool {
        if cnt[p[i]] == cnt[p[j]] {
            return first[p[i]] < first[p[j]]
        }
        return cnt[p[i]] > cnt[p[j]]
    })

    j := 0
    for i := range n {
        idx := getIndex(s[i])
        if idx == -1 {
            ans[i] = s[i]
        } else {
            for cnt[p[j]] == 0 {
                j++
            }
            cnt[p[j]]--
            ans[i] = alp[p[j]]
        }
    }

    return string(ans)
}

func getIndex(c byte) int {
    switch c {
        case 'a':
            return 0;
        case 'e':
            return 1;
        case 'i':
            return 2;
        case 'o':
            return 3;
        case 'u':
            return 4;
        default:
            return -1;
    }
}