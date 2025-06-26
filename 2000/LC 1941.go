const N int = 26
func areOccurrencesEqual(s string) bool {
    cnt := [N]int{}
    for _, c := range s {
        cnt[c - 'a']++
    }

    for _, x := range cnt {
        if x != 0 && x != cnt[s[0] - 'a'] {
            return false
        }
    }

    return true
}