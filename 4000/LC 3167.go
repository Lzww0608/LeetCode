const N int = 26
func betterCompression(s string) string {
    cnt := [N]int{}
    n := len(s)
    for i := 0; i < n; {
        c := int(s[i] - 'a')
        i++
        x := 0
        for i < n && s[i] >= '0' && s[i] <= '9' {
            x = x * 10 + int(s[i] - '0')
            i++
        }

        cnt[c] += x
    }

    ans := strings.Builder{}
    for i, x := range cnt {
        if x == 0 {
            continue
        }
        ans.WriteByte(byte('a' + i))
        ans.WriteString(strconv.Itoa(x))
    }

    return ans.String()
}