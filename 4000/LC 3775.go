func reverseWords(str string) string {
    s := strings.Split(str, " ")
    ans := strings.Builder{}

    calc := func(t string) int {
        cnt := 0
        for i := range t {
            if strings.Contains("aeiou", string(t[i])) {
                cnt++
            }
        }

        return cnt
    }

    cnt := calc(s[0])
    ans.WriteString(s[0])
    if len(s) > 1 {
        ans.WriteByte(' ')
    }
    for i := 1; i < len(s); i++ {
        cur := calc(s[i])
        if cur == cnt {
            tmp := []byte(s[i])
            slices.Reverse(tmp)
            ans.WriteString(string(tmp))
        } else {
            ans.WriteString(s[i])
        }
        if i + 1 < len(s) {
            ans.WriteByte(' ')
        }
    }

    return ans.String()
}