func rearrangeString(s string, x byte, y byte) string {
    cnt := [26]int{}
    for i := range s {
        x := int(s[i] - 'a')
        cnt[x]++
    }

    ans := make([]byte, 0, len(s))
    for range cnt[int(y - 'a')] {
        ans = append(ans, y)
    }
    for range cnt[int(x - 'a')] {
        ans = append(ans, x)
    }

    for i, t := range cnt {
        if i == int(y - 'a') || i == int(x - 'a') {
            continue
        }

        for range t {
            ans = append(ans, byte('a' + i))
        }
    }

    return string(ans)
}