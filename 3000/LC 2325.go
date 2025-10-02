const N int = 26
func decodeMessage(key string, message string) string {
    sub := [N]int{}
    cnt := 1
    for _, c := range key {
        if c == ' ' {
            continue
        }
        x := int(c - 'a')
        if sub[x] == 0 {
            sub[x] = cnt
            cnt++
        }
    }

    ans := []byte(message)
    for i := range ans {
        if ans[i] == ' ' {
            continue
        }
        x := int(ans[i] - 'a')
        ans[i] = byte('a' + sub[x] - 1)
    }

    return string(ans)
}