const N int = 26
func greatestLetter(s string) string {
    cnt := [N]int{}
    for i := range s {
        if s[i] >= 'a' && s[i] <= 'z' {
            x := int(s[i] - 'a')
            cnt[x] |= 1
        } else {
            x := int(s[i] - 'A')
            cnt[x] |= 2
        }
    }

    for i := N - 1; i >= 0; i-- {
        if cnt[i] == 3 {
            return string(byte('A' + i))
        }
    }

    return ""
}