const N int = 26
func checkStrings(s string, t string) bool {
    cnt := [N]int{}
    n := len(s)
    for i := 0; i < n; i += 2 {
        x := int(s[i] - 'a')
        y := int(t[i] - 'a')
        cnt[x]++
        cnt[y]--
    }
    for _, x := range cnt {
        if x != 0 {
            return false
        }
    }

    for i := 1; i < n; i += 2 {
        x := int(s[i] - 'a')
        y := int(t[i] - 'a')
        cnt[x]++
        cnt[y]--
    }
    for _, x := range cnt {
        if x != 0 {
            return false
        }
    }

    return true
}